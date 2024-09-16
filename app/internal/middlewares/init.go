package middlewares

import (
	"context"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/controller"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/shared/enums"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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
	app.Use(requestid.New())
	app.Use(func(c *fiber.Ctx) error {
		// get the request id that was added by requestid middleware
		var requestId = c.Locals("requestid")
		// create a new context and add the requestid to it
		var ctx = context.WithValue(context.Background(), enums.ContextKeyRequestId, requestId)
		c.SetUserContext(ctx)
		return c.Next()
	})
	// routes that don't require a JWT token
	initPublicRoutes(router, profileController)
	//tokenRetrospector := usecases.NewIdentity(cfg, l)
	//app.Use(NewJwtMiddleware(cfg, tokenRetrospector, l))
	// routes that require authentication/authorization
	initProtectedRoutes(router, profileController)
}
