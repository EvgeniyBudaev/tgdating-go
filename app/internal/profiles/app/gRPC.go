package app

import (
	"context"
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/controller"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/repository/psql"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/service"
	"go.uber.org/zap"
	"net"
)

const (
	errorFilePathHttp = "internal/profiles/app/gRPC.go"
)

func (app *App) StartServer(ctx context.Context) error {
	app.fiber.Static("/static", "./static")
	s3Client := config.NewS3(app.config)
	ufw := service.NewUnitOfWorkFactory(app.Logger, app.db.psql)
	navigatorRepository := psql.NewNavigatorRepository(app.Logger, app.db.psql)
	filterRepository := psql.NewFilterRepository(app.Logger, app.db.psql)
	telegramRepository := psql.NewTelegramRepository(app.Logger, app.db.psql)
	imageRepository := psql.NewImageRepository(app.Logger, app.db.psql)
	imageStatusRepository := psql.NewImageStatusRepository(app.Logger, app.db.psql)
	likeRepository := psql.NewLikeRepository(app.Logger, app.db.psql)
	blockRepository := psql.NewBlockRepository(app.Logger, app.db.psql)
	complaintRepository := psql.NewComplaintRepository(app.Logger, app.db.psql)
	statusRepository := psql.NewStatusRepository(app.Logger, app.db.psql)
	profileRepository := psql.NewProfileRepository(app.Logger, app.db.psql)
	profileService := service.NewProfileService(app.Logger, app.db.psql, app.config, app.kafkaWriter, s3Client, ufw,
		profileRepository, navigatorRepository, filterRepository, telegramRepository, imageRepository,
		imageStatusRepository, likeRepository, blockRepository, complaintRepository, statusRepository)
	profileController := controller.NewProfileController(app.Logger, profileService)
	pb.RegisterProfileServer(app.gRPCServer, profileController)
	go func() {
		port := ":" + app.config.ProfilesPort
		app.Logger.Info("Starting Profile service on port: ", zap.String("port", port))
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
	}()
	select {
	case <-ctx.Done():
		if err := app.fiber.Shutdown(); err != nil {
			errorMessage := getErrorMessage("StartServer", "Shutdown",
				errorFilePathHttp)
			app.Logger.Error(errorMessage, zap.Error(err))
		}
	}
	return nil
}
