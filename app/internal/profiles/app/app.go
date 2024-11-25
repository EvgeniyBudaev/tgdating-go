package app

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"os"
)

const (
	errorFilePathApp = "internal/profiles/app/app.go"
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
	migrationsPath := fmt.Sprintf("file://%s/migrations", dir)
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		cfg.DBName,
		driver)
	if err != nil {
		errorMessage := getErrorMessage("New", "NewWithDatabaseInstance",
			errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	m.Up()

	// gRPC-сервер
	s := grpc.NewServer()

	// Kafka
	w := &kafka.Writer{
		Addr:         kafka.TCP("127.0.0.1:9095", "27.0.0.1:9096", "127.0.0.1:9097"),
		Topic:        "like_topic",
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    1048576,
		BatchTimeout: 1000,
		Compression:  kafka.Gzip,
		RequiredAcks: kafka.RequireOne,
	}

	// Fiber
	f := fiber.New(fiber.Config{
		ReadBufferSize: 256 << 8,
		BodyLimit:      50 * 1024 * 1024, // 50 MB
	})

	// CORS
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, X-Requested-With, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	return &App{
		config:      cfg,
		db:          database,
		fiber:       f,
		gRPCServer:  s,
		kafkaWriter: w,
		Logger:      loggerLevel,
	}
}

// Run launches the application
func (app *App) Run(ctx context.Context) {
	g, ctx := errgroup.WithContext(ctx)

	// Start server
	g.Go(func() error {
		if err := app.StartServer(ctx); err != nil {
			errorMessage := getErrorMessage("Run", "StartServer",
				errorFilePathApp)
			app.Logger.Fatal(errorMessage, zap.Error(err))
			return err
		}
		return nil
	})

	// Close kafka writer when context done
	g.Go(func() error {
		select {
		case <-ctx.Done():
			if err := app.kafkaWriter.Close(); err != nil {
				errorMessage := getErrorMessage("Run", "kafkaWriter.Close",
					errorFilePathApp)
				app.Logger.Fatal(errorMessage, zap.Error(err))
			}
		}
		return nil
	})

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
