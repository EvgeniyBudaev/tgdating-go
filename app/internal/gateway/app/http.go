package app

import (
	"context"
	proto "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/controller"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/middlewares"
	"go.uber.org/zap"
)

const (
	errorFilePathHttp = "internal/app/gRPC.go"
)

func (app *App) StartHTTPServer(ctx context.Context, proto proto.ProfileClient) error {
	app.fiber.Static("/static", "./static")
	profileController := controller.NewProfileController(app.Logger, proto)
	middlewares.InitFiberMiddlewares(
		app.fiber, app.config, app.Logger, profileController, InitPublicRoutes, InitProtectedRoutes)
	go func() {
		port := ":" + app.config.GatewayPort
		if err := app.fiber.Listen(port); err != nil {
			errorMessage := getErrorMessage("StartHTTPServer", "Listen",
				errorFilePathHttp)
			app.Logger.Error(errorMessage, zap.Error(err))
		}
	}()
	select {
	case <-ctx.Done():
		if err := app.fiber.Shutdown(); err != nil {
			errorMessage := getErrorMessage("StartHTTPServer", "Shutdown",
				errorFilePathHttp)
			app.Logger.Error(errorMessage, zap.Error(err))
		}
	}
	return nil
}
