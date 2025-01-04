package app

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/telegram/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/telegram/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

const (
	errorFilePathApp = "internal/telegram/app/app.go"
	bodyLimit        = 61 * 1024 * 1024 // 61 MB
)

// App - application structure
type App struct {
	config      *config.Config
	fiber       *fiber.App
	gRPCServer  *grpc.Server
	kafkaReader *kafka.Reader
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

	// Kafka
	var brokers = []string{cfg.Kafka1, cfg.Kafka2, cfg.Kafka3}
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  "consumer-group-id",
		Topic:    "like_topic",
		MaxBytes: bodyLimit,
	})

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
		config:      cfg,
		fiber:       f,
		kafkaReader: r,
		Logger:      loggerLevel,
	}
}

// Run launches the application
func (app *App) Run(ctx context.Context) {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		if err := app.StartBot(ctx); err != nil {
			errorMessage := getErrorMessage("Run", "StartBot", errorFilePathApp)
			app.Logger.Fatal(errorMessage, zap.Error(err))
			return err
		}
		return nil
	})
	// Close kafka reader when context done
	g.Go(func() error {
		select {
		case <-ctx.Done():
			if err := app.kafkaReader.Close(); err != nil {
				errorMessage := getErrorMessage("Run", "kafkaReader.Close",
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
