package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/controller"
	v1 "github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/controller/http/api/v1"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/logger"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enum"
	"github.com/Luzifer/go-openssl/v4"
	"github.com/gofiber/fiber/v2"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"go.uber.org/zap"
	"net/http"
	//"time"
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

	app.Use(NewJwtMiddleware(config, logger))
	// routes that require authentication/authorization
	initProtectedRoutes(app, profileController)
}

func NewJwtMiddleware(config *config.Config, logger logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return successHandler(c, config, logger)
	}
}

func successHandler(c *fiber.Ctx, config *config.Config, logger logger.Logger) error {
	encryptedToken := c.Get("Authorization")
	secretKey := config.CryptoSecretKey
	fmt.Println("encryptedToken: ", encryptedToken)
	authData, err := decrypt(encryptedToken, secretKey)
	if err != nil {
		errorMessage := "invalid decrypt token"
		err := errors.New(errorMessage)
		logger.Debug(errorMessage, zap.Error(err))
		return v1.ResponseError(c, err, http.StatusUnauthorized)
	}
	// Validate init data. We consider init data sign valid for 1 hour from their creation moment
	//if err := initdata.Validate(authData, config.TelegramBotToken, time.Hour); err != nil {
	//	errorMessage := "invalid token"
	//	err := errors.New(errorMessage)
	//	logger.Debug(errorMessage, zap.Error(err))
	//	return v1.ResponseError(c, err, http.StatusUnauthorized)
	//}
	// Parse init data
	telegramInitData, err := initdata.Parse(authData)
	if err != nil {
		errorMessage := "invalid parse token"
		err := errors.New(errorMessage)
		logger.Debug(errorMessage, zap.Error(err))
		return v1.ResponseError(c, err, http.StatusUnauthorized)
	}
	// Save to context
	var ctx = c.UserContext()
	var contextWithClaims = context.WithValue(ctx, enum.ContextKeyTelegram, telegramInitData)
	c.SetUserContext(contextWithClaims)
	return c.Next()
}

func decrypt(encryptedString, secretKey string) (string, error) {
	o := openssl.New()
	key := openssl.BytesToKeyMD5
	dec, err := o.DecryptBytes(secretKey, []byte(encryptedString), key)
	if err != nil {
		return "", err
	}
	return string(dec), nil
}
