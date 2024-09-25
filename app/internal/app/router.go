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
	router.Get("/profiles/detail/:sessionId", profileController.GetProfileDetail())
	router.Get("/profiles/short/:sessionId", profileController.GetProfileShortInfo())
	router.Get("/profiles/list", profileController.GetProfileList())
	router.Post("/profiles/blocks", profileController.AddBlock())
	router.Post("/profiles/likes", profileController.AddLike())
	router.Post("/profiles/complaints", profileController.AddComplaint())
}

func InitProtectedRoutes(router fiber.Router, profileController *controller.ProfileController) {}
