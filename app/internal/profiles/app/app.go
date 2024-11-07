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
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
	"sync"
)

const (
	errorFilePathApp = "internal/profiles/app/app.go"
)

// ProfileServer поддерживает все необходимые методы сервера.
//type ProfileServer struct {
//	// нужно встраивать тип pb.Unimplemented<TypeName>
//	// для совместимости с будущими версиями
//	pb.UnimplementedProfileServer
//}

//func NewProfileServer() *ProfileServer {
//	return &ProfileServer{}
//}

// Add реализует интерфейс добавления пользователя.
//func (s *ProfileServer) Add(ctx context.Context, in *pb.ProfileAddRequest) (*pb.ProfileAddResponse, error) {
//	fmt.Println("RESULT sessionId: ", in.SessionId)
//	return &pb.ProfileAddResponse{SessionId: in.SessionId}, nil
//}

// App - application structure
type App struct {
	Logger     logger.Logger
	config     *config.Config
	db         *Database
	fiber      *fiber.App
	gRPCServer *grpc.Server
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
	driver, err := postgres.WithInstance(database.psql, &postgres.Config{})
	if err != nil {
		errorMessage := getErrorMessage("New", "WithInstance", errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	dir, err := os.Getwd()
	if err != nil {
		errorMessage := getErrorMessage("New", "os.Getwd", errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	migrationsPath := fmt.Sprintf("file://%s/migrations/profiles", dir)
	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	if err != nil {
		errorMessage := getErrorMessage("New", "NewWithDatabaseInstance",
			errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	m.Up()

	// gRPC-сервер
	s := grpc.NewServer()

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
		config:     cfg,
		db:         database,
		Logger:     loggerLevel,
		fiber:      f,
		gRPCServer: s,
	}
}

// Run launches the application
func (app *App) Run(ctx context.Context) {
	// Hub for telegram bot
	hub := entity.NewHub()

	//msgChan := make(chan *entity.HubContent, 1) // msgChan - канал для передачи сообщений
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := app.StartServer(ctx, hub); err != nil {
			errorMessage := getErrorMessage("Run", "StartServer",
				errorFilePathApp)
			app.Logger.Fatal(errorMessage, zap.Error(err))
		}
		wg.Done()
	}()
	//wg.Add(1)
	//go func() {
	//	if err := app.StartBot(ctx, msgChan); err != nil {
	//		errorMessage := getErrorMessage("Run", "StartBot", errorFilePathApp)
	//		app.Logger.Fatal(errorMessage, zap.Error(err))
	//	}
	//	wg.Done()
	//}()
	//wg.Add(1)
	//go func() {
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			wg.Done()
	//			return
	//		case c, ok := <-hub.Broadcast:
	//			if !ok {
	//				return
	//			}
	//			msgChan <- c
	//		}
	//	}
	//}()
	wg.Wait()
}

func getErrorMessage(repositoryMethodName, callMethodName, errorFilePath string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}
