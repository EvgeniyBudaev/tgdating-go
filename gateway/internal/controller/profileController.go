package controller

import (
	v1 "github.com/EvgeniyBudaev/tgdating-go/gateway/internal/controller/http/api/v1"
	"github.com/EvgeniyBudaev/tgdating-go/gateway/internal/logger"
	"github.com/gofiber/fiber/v2"
)

type ProfileController struct {
	logger logger.Logger
}

func NewProfileController(l logger.Logger) *ProfileController {
	return &ProfileController{
		logger: l,
	}
}

func (pc *ProfileController) AddProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /gateway/api/v1/profiles")
		return v1.ResponseCreated(ctf, "OK")
	}
}
