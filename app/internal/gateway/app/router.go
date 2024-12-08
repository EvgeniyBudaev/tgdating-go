package app

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/controller"
	"github.com/gofiber/fiber/v2"
)

var prefix = "/api/v1"

func InitPublicRoutes(app *fiber.App, profileController *controller.ProfileController) {
	router := app.Group(prefix)
	router.Get("/profiles/telegram/:telegramUserId", profileController.GetProfile())
	router.Get("/profiles/detail/:viewedTelegramUserId", profileController.GetProfileDetail())
	router.Get("/profiles/short/:telegramUserId", profileController.GetProfileShortInfo())
	router.Get("/profiles/list", profileController.GetProfileList())
}

func InitProtectedRoutes(app *fiber.App, profileController *controller.ProfileController) {
	router := app.Group(prefix)
	router.Post("/profiles", profileController.AddProfile())
	router.Put("/profiles", profileController.UpdateProfile())
	router.Post("/profiles/freeze", profileController.FreezeProfile())
	router.Post("/profiles/restore", profileController.RestoreProfile())
	router.Delete("/profiles", profileController.DeleteProfile())
	router.Delete("/profiles/images/:id", profileController.DeleteImage())
	router.Put("/profiles/filters", profileController.UpdateFilter())
	router.Put("/profiles/navigators", profileController.UpdateCoordinates())
	router.Post("/profiles/blocks", profileController.AddBlock())
	router.Post("/profiles/likes", profileController.AddLike())
	router.Put("/profiles/likes", profileController.UpdateLike())
	router.Post("/profiles/complaints", profileController.AddComplaint())
}
