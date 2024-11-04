package app

import (
	"context"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/controller"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/middlewares"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/repository/psql"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/service"
	"go.uber.org/zap"
)

const (
	errorFilePathHttp = "internal/app/http.go"
)

func (app *App) StartHTTPServer(ctx context.Context, hub *entity.Hub) error {
	app.fiber.Static("/static", "./static")
	done := make(chan struct{})
	s3Client := config.NewS3(app.config)
	navigatorRepository := psql.NewNavigatorRepository(app.Logger, app.db.psql)
	filterRepository := psql.NewFilterRepository(app.Logger, app.db.psql)
	telegramRepository := psql.NewTelegramRepository(app.Logger, app.db.psql)
	imageRepository := psql.NewImageRepository(app.Logger, app.db.psql)
	likeRepository := psql.NewLikeRepository(app.Logger, app.db.psql)
	blockRepository := psql.NewBlockRepository(app.Logger, app.db.psql)
	complaintRepository := psql.NewComplaintRepository(app.Logger, app.db.psql)
	profileRepository := psql.NewProfileRepository(app.Logger, app.db.psql)
	profileService := service.NewProfileService(app.Logger, app.config, hub, s3Client, profileRepository, navigatorRepository, filterRepository,
		telegramRepository, imageRepository, likeRepository, blockRepository, complaintRepository)
	profileController := controller.NewProfileController(app.Logger, profileService)
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
