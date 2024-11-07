package app

import (
	"context"
	"fmt"
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/controller"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/repository/psql"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/service"
	"go.uber.org/zap"
	"net"
)

const (
	errorFilePathHttp = "internal/app/gRPC.go"
)

func (app *App) StartServer(ctx context.Context, hub *entity.Hub) error {
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
	//middlewares.InitFiberMiddlewares(
	//	app.fiber, app.config, app.Logger, profileController, InitPublicRoutes, InitProtectedRoutes)
	profileController := controller.NewProfileController(app.Logger, profileService)
	pb.RegisterProfileServer(app.gRPCServer, profileController)
	fmt.Println("Сервер gRPC слушатель")
	go func() {
		port := ":" + app.config.ProfilesPort
		listen, err := net.Listen("tcp", port)
		if err != nil {
			errorMessage := getErrorMessage("StartServer", "net.Listen",
				errorFilePathApp)
			app.Logger.Error(errorMessage, zap.Error(err))
		}
		if err := app.gRPCServer.Serve(listen); err != nil {
			errorMessage := getErrorMessage("StartServer", "s.Serve",
				errorFilePathApp)
			app.Logger.Error(errorMessage, zap.Error(err))
		}
		close(done)
	}()
	select {
	case <-ctx.Done():
		if err := app.fiber.Shutdown(); err != nil {
			errorMessage := getErrorMessage("StartServer", "Shutdown",
				errorFilePathHttp)
			app.Logger.Error(errorMessage, zap.Error(err))
		}
	case <-done:
		app.Logger.Info("server finished successfully")
	}
	return nil
}
