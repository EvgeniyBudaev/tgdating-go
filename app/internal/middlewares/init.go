package middlewares

import (
	"context"
	"errors"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/controller"
	v1 "github.com/EvgeniyBudaev/tgdating-go/app/internal/controller/http/api/v1"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/shared/enums"
	"github.com/gofiber/fiber/v2"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"net/http"
	"strings"
	"time"
)

func InitFiberMiddlewares(
	app *fiber.App,
	config *config.Config,
	logger logger.Logger,
	profileController *controller.ProfileController,
	initPublicRoutes func(app *fiber.App, profileController *controller.ProfileController),
	initProtectedRoutes func(app *fiber.App, profileController *controller.ProfileController),
) {
	// routes that don't require a JWT token
	initPublicRoutes(app, profileController)

	app.Use(NewJwtMiddleware(config))
	// routes that require authentication/authorization
	initProtectedRoutes(app, profileController)
}

func NewJwtMiddleware(config *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return successHandler(c, config)
	}
}

func successHandler(c *fiber.Ctx, config *config.Config) error {
	token := c.Get("Authorization")
	authParts := strings.Split(token, " ")
	if len(authParts) != 2 {
		err := errors.New("unauthorized")
		return v1.ResponseError(c, err, http.StatusUnauthorized)
	}
	authType := authParts[0]
	authData := authParts[1]
	switch authType {
	case "tma":
		// Validate init data. We consider init data sign valid for 1 hour from their creation moment
		if err := initdata.Validate(authData, config.TelegramBotToken, time.Hour); err != nil {
			err := errors.New("invalid token")
			return v1.ResponseError(c, err, http.StatusUnauthorized)
		}
		// Parse init data
		telegramInitData, err := initdata.Parse(authData)
		if err != nil {
			err := errors.New("invalid parse token")
			return v1.ResponseError(c, err, http.StatusUnauthorized)
		}
		// Save to context
		var ctx = c.UserContext()
		var contextWithClaims = context.WithValue(ctx, enums.ContextKeyTelegram, telegramInitData)
		c.SetUserContext(contextWithClaims)
		return c.Next()
	}
	err := errors.New("unauthorized")
	return v1.ResponseError(c, err, http.StatusUnauthorized)
}
