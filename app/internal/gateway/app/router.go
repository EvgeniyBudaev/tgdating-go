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
	router.Get("/profiles/:telegramUserId/check", profileController.CheckProfileExists())
	router.Get("/profiles/filters/:telegramUserId", profileController.GetFilter())
	router.Get("/profiles/:telegramUserId/premium/check", profileController.CheckPremium())
	router.Get("/profiles/:telegramUserId/blocks/list", profileController.GetBlockedList())

	//router.Post("/profiles", profileController.AddProfile())
	//router.Put("/profiles", profileController.UpdateProfile())
	//router.Post("/profiles/freeze", profileController.FreezeProfile())
	//router.Post("/profiles/restore", profileController.RestoreProfile())
	//router.Delete("/profiles", profileController.DeleteProfile())
	//router.Delete("/profiles/images/:id", profileController.DeleteImage())
	//router.Put("/profiles/filters", profileController.UpdateFilter())
	//router.Put("/profiles/navigators", profileController.UpdateCoordinates())
	//router.Post("/profiles/blocks", profileController.AddBlock())
	//router.Post("/profiles/likes", profileController.AddLike())
	//router.Put("/profiles/likes", profileController.UpdateLike())
	//router.Post("/profiles/complaints", profileController.AddComplaint())
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
	router.Put("/profiles/unblock", profileController.Unblock())
	router.Post("/profiles/likes", profileController.AddLike())
	router.Put("/profiles/likes", profileController.UpdateLike())
	router.Post("/profiles/likes/last", profileController.GetLastLike())
	router.Post("/profiles/complaints", profileController.AddComplaint())
	router.Post("/profiles/payments", profileController.AddPayment())
	router.Put("/profiles/settings", profileController.UpdateSettings())
}
