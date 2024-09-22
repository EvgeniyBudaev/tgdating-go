package app

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
	"sync"
)

const (
	errorFilePathApp = "internal/app/app.go"
)

// App - application structure
type App struct {
	Logger logger.Logger
	config *config.Config
	db     *Database
	fiber  *fiber.App
}

// New - create new application
func New() *App {
	// Default logger
	defaultLogger, err := logger.New(logger.GetDefaultLevel())
	if err != nil {
		errorMessage := getErrorMessage("New", "logger.New")
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	// Config
	cfg, err := config.Load(defaultLogger)
	if err != nil {
		errorMessage := getErrorMessage("New", "config.Load")
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	// Logger level
	loggerLevel, err := logger.New(cfg.LoggerLevel)
	if err != nil {
		errorMessage := getErrorMessage("New", "logger.New")
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	// Database connection
	postgresConnection, err := newPostgresConnection(cfg)
	if err != nil {
		errorMessage := getErrorMessage("New", "newPostgresConnection")
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	database := NewDatabase(loggerLevel, postgresConnection)
	err = postgresConnection.Ping()
	if err != nil {
		errorMessage := getErrorMessage("New", "Ping")
		defaultLogger.Fatal(errorMessage, zap.Error(err))
	}
	// Fiber
	f := fiber.New(fiber.Config{
		ReadBufferSize: 4 << 12,
	})
	// CORS
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, X-Requested-With, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	return &App{
		config: cfg,
		db:     database,
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
			errorMessage := getErrorMessage("Run", "StartHTTPServer")
			app.Logger.Fatal(errorMessage, zap.Error(err))
		}
		wg.Done()
	}()
	wg.Wait()
}

func getErrorMessage(repositoryMethodName, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathApp)
}
