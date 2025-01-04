package app

import (
	"context"
	"fmt"
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
)

const (
	errorFilePathApp = "internal/app/app.go"
)

// App - application structure
type App struct {
	Logger      logger.Logger
	config      *config.Config
	fiber       *fiber.App
	kafkaWriter *kafka.Writer
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

	// Fiber
	f := fiber.New(fiber.Config{
		ReadBufferSize: 256 << 8,
		BodyLimit:      61 * 1024 * 1024, // 61 MB
	})

	// Rate limiter to prevent DDOS attacks
	// https://docs.gofiber.io/api/middleware/limiter/
	f.Use(limiter.New(limiter.Config{
		Max:        300,
		Expiration: 30 * time.Second,
	}))

	// Kafka
	w := &kafka.Writer{
		Addr:         kafka.TCP(cfg.Kafka1, cfg.Kafka2, cfg.Kafka3),
		Topic:        "like_topic",
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    1048576,
		BatchTimeout: 1000,
		Compression:  kafka.Gzip,
		RequiredAcks: kafka.RequireOne,
	}

	// CORS
	f.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AllowOrigins,
		AllowHeaders: "Content-Type, X-Requested-With, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	return &App{
		config:      cfg,
		Logger:      loggerLevel,
		fiber:       f,
		kafkaWriter: w,
	}
}

// Run launches the application
func (app *App) Run(ctx context.Context) {
	addr := app.config.ProfilesHost
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		errorMessage := getErrorMessage("New", "grpc.NewClient", errorFilePathApp)
		app.Logger.Fatal(errorMessage, zap.Error(err))
	}
	defer conn.Close()
	c := pb.NewProfileClient(conn)
	app.Logger.Info("Listening gRPC server on host: ", zap.String("host", addr))

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := app.StartHTTPServer(ctx, c); err != nil {
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
