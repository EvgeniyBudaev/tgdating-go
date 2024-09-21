package app

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func InitPublicRoutes(router fiber.Router, profileController *controller.ProfileController) {
	router.Post("/profiles", profileController.AddProfile())
	router.Put("/profiles", profileController.UpdateProfile())
	router.Delete("/profiles", profileController.DeleteProfile())
	router.Get("/profiles/:sessionId", profileController.GetProfileBySessionId())
	router.Post("/profiles/blocks", profileController.AddBlock())
	router.Post("/profiles/likes", profileController.AddLike())
	router.Post("/profiles/complaints", profileController.AddComplaint())
}

func InitProtectedRoutes(router fiber.Router, profileController *controller.ProfileController) {}
