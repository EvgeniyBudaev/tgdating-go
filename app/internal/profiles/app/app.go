package app

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/segmentio/kafka-go"

	//"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"os"
	"strings"
)

const (
	errorFilePathApp = "internal/profiles/app/app.go"
	bodyLimit        = 61 * 1024 * 1024 // 61 MB
)

// App - application structure
type App struct {
	config      *config.Config
	db          *Database
	fiber       *fiber.App
	gRPCServer  *grpc.Server
	kafkaWriter *kafka.Writer
	Logger      logger.Logger
}

// New - create new application
func New() *App {
	// Default logger
	defaultLogger, err := logger.New(logger.GetDefaultLevel())
	if err != nil {
		errorMessage := getErrorMessage("New", "logger.New", errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}

	// Config
	cfg, err := config.Load(defaultLogger)
	if err != nil {
		errorMessage := getErrorMessage("New", "config.Load", errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}

	// Logger level
	loggerLevel, err := logger.New(cfg.LoggerLevel)
	if err != nil {
		errorMessage := getErrorMessage("New", "logger.New", errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}

	// Database connection
	postgresConnection, err := newPostgresConnection(cfg)
	if err != nil {
		errorMessage := getErrorMessage("New", "newPostgresConnection",
			errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	database := NewDatabase(loggerLevel, postgresConnection)
	err = postgresConnection.Ping()
	if err != nil {
		errorMessage := getErrorMessage("New", "Ping", errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}

	// Auto migrate
	// Create schema
	_, err = database.psql.Exec("CREATE SCHEMA IF NOT EXISTS dating;")
	if err != nil {
		errorMessage := getErrorMessage("New", "CREATE SCHEMA", errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	driver, err := postgres.WithInstance(database.psql, &postgres.Config{
		DatabaseName: cfg.DBName,
		SchemaName:   cfg.DBSchema,
	})
	if err != nil {
		errorMessage := getErrorMessage("New", "WithInstance", errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	dir, err := os.Getwd()
	if err != nil {
		errorMessage := getErrorMessage("New", "os.Getwd", errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	migrationsPath := fmt.Sprintf("file://%s/migrations/", dir)
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		cfg.DBName,
		driver)
	if err != nil {
		errorMessage := getErrorMessage("New", "NewWithDatabaseInstance",
			errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	err = m.Up()
	if err != nil && !strings.Contains(err.Error(), "no change") {
		errorMessage := getErrorMessage("New", "m.Up", errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}

	// gRPC-сервер
	s := grpc.NewServer()

	// Kafka
	//w := &kafka.Writer{
	//	//Addr:         kafka.TCP("172.18.0.1:10095", "172.18.0.1:10096", "172.18.0.1:10097"), // docker inspect network web-network
	//	Addr:         kafka.TCP("127.0.0.1:10095", "127.0.0.1:10096", "127.0.0.1:10097"),
	//	Topic:        "like_topic",
	//	Balancer:     &kafka.LeastBytes{},
	//	BatchSize:    1048576,
	//	BatchTimeout: 1000,
	//	Compression:  kafka.Gzip,
	//	RequiredAcks: kafka.RequireOne,
	//}

	// Fiber
	f := fiber.New(fiber.Config{
		ReadBufferSize: 256 << 8,
		BodyLimit:      bodyLimit,
	})

	// CORS
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, X-Requested-With, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	return &App{
		config:     cfg,
		db:         database,
		fiber:      f,
		gRPCServer: s,
		//kafkaWriter: w,
		Logger: loggerLevel,
	}
}

// Run launches the application
func (app *App) Run(ctx context.Context) {
	g, ctx := errgroup.WithContext(ctx)
	// Hub for telegram bot
	hub := entity.NewHub()
	msgChan := make(chan *entity.HubContent, 1) // msgChan - канал для передачи сообщений

	// Start server
	g.Go(func() error {
		if err := app.StartServer(ctx, hub); err != nil {
			errorMessage := getErrorMessage("Run", "StartServer",
				errorFilePathApp)
			app.Logger.Fatal(errorMessage, zap.Error(err))
			return err
		}
		return nil
	})

	// Start telegram bot
	g.Go(func() error {
		if err := app.StartBot(ctx, msgChan); err != nil {
			errorMessage := getErrorMessage("Run", "StartServer",
				errorFilePathApp)
			app.Logger.Fatal(errorMessage, zap.Error(err))
			return err
		}
		return nil
	})

	// Start Hub
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		case c, ok := <-hub.Broadcast:
			if !ok {
				return nil
			}
			msgChan <- c
		}
		return nil
	})

	// Close kafka writer when context done
	//g.Go(func() error {
	//	select {
	//	case <-ctx.Done():
	//		if err := app.kafkaWriter.Close(); err != nil {
	//			errorMessage := getErrorMessage("Run", "kafkaWriter.Close",
	//				errorFilePathApp)
	//			app.Logger.Fatal(errorMessage, zap.Error(err))
	//		}
	//	}
	//	return nil
	//})

	if err := g.Wait(); err != nil {
		errorMessage := getErrorMessage("Run", "g.Wait",
			errorFilePathApp)
		app.Logger.Fatal(errorMessage, zap.Error(err))
	}
}

func getErrorMessage(repositoryMethodName, callMethodName, errorFilePath string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}
