package app

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/controller"
	"github.com/gofiber/fiber/v2"
)

var prefix = "/gateway/api/v1"

func InitPublicRoutes(app *fiber.App, profileController *controller.ProfileController) {
	router := app.Group(prefix)
	router.Get("/profiles/session/:sessionId", profileController.GetProfileBySessionId())
	router.Get("/profiles/detail/:viewedSessionId", profileController.GetProfileDetail())
}

func InitProtectedRoutes(app *fiber.App, profileController *controller.ProfileController) {
	router := app.Group(prefix)
	router.Post("/profiles", profileController.AddProfile())
	router.Put("/profiles", profileController.UpdateProfile())
}
