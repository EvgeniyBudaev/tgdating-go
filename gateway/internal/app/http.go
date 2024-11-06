package app

import (
	"context"
	"github.com/EvgeniyBudaev/tgdating-go/gateway/internal/controller"
	"github.com/EvgeniyBudaev/tgdating-go/gateway/internal/middlewares"
	"go.uber.org/zap"
)

const (
	errorFilePathHttp = "internal/app/http.go"
)

func (app *App) StartHTTPServer(ctx context.Context) error {
	app.fiber.Static("/static", "./static")
	done := make(chan struct{})
	profileController := controller.NewProfileController(app.Logger)
	middlewares.InitFiberMiddlewares(
		app.fiber, app.config, app.Logger, profileController, InitPublicRoutes, InitProtectedRoutes)
	go func() {
		port := ":" + app.config.Port
		if err := app.fiber.Listen(port); err != nil {
			errorMessage := getErrorMessage("StartHTTPServer", "Listen",
				errorFilePathHttp)
			app.Logger.Error(errorMessage, zap.Error(err))
		}
		close(done)
	}()
	select {
	case <-ctx.Done():
		if err := app.fiber.Shutdown(); err != nil {
			errorMessage := getErrorMessage("StartHTTPServer", "Shutdown",
				errorFilePathHttp)
			app.Logger.Error(errorMessage, zap.Error(err))
		}
	case <-done:
		app.Logger.Info("server finished successfully")
	}
	return nil
}
