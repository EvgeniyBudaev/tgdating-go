package app

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func InitPublicRoutes(router fiber.Router, profileController *controller.ProfileController) {
	router.Post("/profiles", profileController.AddProfile())
	router.Put("/profiles", profileController.UpdateProfile())
	router.Delete("/profiles", profileController.DeleteProfile())
	router.Post("/profiles/blocks", profileController.AddBlock())
}

func InitProtectedRoutes(router fiber.Router, profileController *controller.ProfileController) {}
