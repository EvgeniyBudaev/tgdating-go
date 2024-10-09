package middlewares

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/controller"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"github.com/gofiber/fiber/v2"
)

func InitFiberMiddlewares(
	app *fiber.App,
	config *config.Config,
	logger logger.Logger,
	router fiber.Router,
	profileController *controller.ProfileController,
	initPublicRoutes func(router fiber.Router, profileController *controller.ProfileController),
	initProtectedRoutes func(router fiber.Router, profileController *controller.ProfileController),
) {
	// routes that don't require a JWT token
	initPublicRoutes(router, profileController)
	initProtectedRoutes(router, profileController)
}
