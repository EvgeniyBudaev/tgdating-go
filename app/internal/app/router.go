package app

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func InitPublicRoutes(router fiber.Router, profileController *controller.ProfileController) {
	router.Post("/profiles", profileController.AddProfile())
}

func InitProtectedRoutes(router fiber.Router, profileController *controller.ProfileController) {}
