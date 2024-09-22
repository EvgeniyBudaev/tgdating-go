package app

import (
	"context"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/controller"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/middlewares"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/repository/psql"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/service"
	"go.uber.org/zap"
)

var prefix = "/api/v1"

func (app *App) StartHTTPServer(ctx context.Context) error {
	app.fiber.Static("/static", "./static")
	done := make(chan struct{})
	navigatorRepository := psql.NewNavigatorRepository(app.Logger, app.db.psql)
	filterRepository := psql.NewFilterRepository(app.Logger, app.db.psql)
	telegramRepository := psql.NewTelegramRepository(app.Logger, app.db.psql)
	imageRepository := psql.NewImageRepository(app.Logger, app.db.psql)
	likeRepository := psql.NewLikeRepository(app.Logger, app.db.psql)
	blockRepository := psql.NewBlockRepository(app.Logger, app.db.psql)
	complaintRepository := psql.NewComplaintRepository(app.Logger, app.db.psql)
	profileRepository := psql.NewProfileRepository(app.Logger, app.db.psql)
	profileService := service.NewProfileService(app.Logger, profileRepository, navigatorRepository, filterRepository,
		telegramRepository, imageRepository, likeRepository, blockRepository, complaintRepository)
	profileController := controller.NewProfileController(app.Logger, profileService)
	router := app.fiber.Group(prefix)
	middlewares.InitFiberMiddlewares(
		app.fiber, app.config, app.Logger, router, profileController, InitPublicRoutes, InitProtectedRoutes)
	go func() {
		port := ":" + app.config.Port
		if err := app.fiber.Listen(port); err != nil {
			app.Logger.Fatal("error func StartHTTPServer, method Listen by path internal/app/http.go",
				zap.Error(err))
		}
		close(done)
	}()
	select {
	case <-ctx.Done():
		if err := app.fiber.Shutdown(); err != nil {
			app.Logger.Error("error func StartHTTPServer, method Shutdown by path internal/app/http.go,"+
				" error shutting down the server", zap.Error(err))
		}
	case <-done:
		app.Logger.Info("server finished successfully")
	}
	return nil
}
