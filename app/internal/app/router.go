package app

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func InitPublicRoutes(router fiber.Router, profileController *controller.ProfileController) {
	router.Post("/profiles", profileController.AddProfile())
	router.Put("/profiles", profileController.UpdateProfile())
	router.Delete("/profiles", profileController.DeleteProfile())
	router.Get("/profiles/session/:sessionId", profileController.GetProfileBySessionId())
	router.Get("/profiles/detail/:viewedSessionId", profileController.GetProfileDetail())
	router.Get("/profiles/short/:sessionId", profileController.GetProfileShortInfo())
	router.Get("/profiles/list", profileController.GetProfileList())
	router.Get("/profiles/:sessionId/images/:fileName", profileController.GetImageBySessionId())
	router.Delete("/profiles/images/:id", profileController.DeleteImage())
	router.Get("/profiles/filter/:sessionId", profileController.GetFilterBySessionId())
	router.Put("/profiles/filters", profileController.UpdateFilter())
	router.Post("/profiles/blocks", profileController.AddBlock())
	router.Post("/profiles/likes", profileController.AddLike())
	router.Put("/profiles/likes", profileController.UpdateLike())
	router.Post("/profiles/complaints", profileController.AddComplaint())
	router.Put("/profiles/navigators", profileController.UpdateCoordinates())
}

func InitProtectedRoutes(router fiber.Router, profileController *controller.ProfileController) {}
