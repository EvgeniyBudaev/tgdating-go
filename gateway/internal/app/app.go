package app

import (
	"context"
	"fmt"
	pb "github.com/EvgeniyBudaev/tgdating-go/app/protobuf"
	"github.com/EvgeniyBudaev/tgdating-go/gateway/internal/config"
	"github.com/EvgeniyBudaev/tgdating-go/gateway/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
)

const (
	errorFilePathApp = "internal/app/app.go"
)

// App - application structure
type App struct {
	Logger logger.Logger
	config *config.Config
	fiber  *fiber.App
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

	// gRPC
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		errorMessage := getErrorMessage("New", "grpc.Dial", errorFilePathApp)
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	defer conn.Close()
	c := pb.NewUsersClient(conn)

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
		config: cfg,
		Logger: loggerLevel,
		fiber:  f,
	}
}

// Run launches the application
func (app *App) Run(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := app.StartHTTPServer(ctx); err != nil {
			errorMessage := getErrorMessage("Run", "StartHTTPServer",
				errorFilePathApp)
			app.Logger.Fatal(errorMessage, zap.Error(err))
		}
		wg.Done()
	}()
	wg.Wait()
}

func getErrorMessage(repositoryMethodName, callMethodName, errorFilePath string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}
