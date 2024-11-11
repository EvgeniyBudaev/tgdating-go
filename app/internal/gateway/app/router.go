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
	router.Get("/profiles/short/:sessionId", profileController.GetProfileShortInfo())
	router.Get("/profiles/list", profileController.GetProfileList())
	router.Get("/profiles/filter/:sessionId", profileController.GetFilterBySessionId())
	//router.Get("/profiles/:sessionId/images/:fileName", profileController.GetImageBySessionId()) // not used
}

func InitProtectedRoutes(app *fiber.App, profileController *controller.ProfileController) {
	router := app.Group(prefix)
	router.Post("/profiles", profileController.AddProfile())
	router.Put("/profiles", profileController.UpdateProfile())
	router.Put("/profiles/filters", profileController.UpdateFilter())
	router.Post("/profiles/blocks", profileController.AddBlock())
	router.Post("/profiles/likes", profileController.AddLike())
	router.Put("/profiles/likes", profileController.UpdateLike())
	router.Post("/profiles/complaints", profileController.AddComplaint())
}
