package controller

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/controller/http/api/v1"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	errorFilePath = "internal/controller/profileController.go"
)

type ProfileService interface {
	AddProfile(ctx context.Context, ctf *fiber.Ctx,
		pr *request.ProfileAddRequestDto) (*response.ProfileAddResponseDto, error)
	UpdateProfile(ctx context.Context, ctf *fiber.Ctx,
		pr *request.ProfileUpdateRequestDto) (*response.ProfileUpdateResponseDto, error)
}

type ProfileController struct {
	logger  logger.Logger
	service ProfileService
}

const TimeoutDuration = 30 * time.Second

func NewProfileController(l logger.Logger, ps ProfileService) *ProfileController {
	return &ProfileController{
		logger:  l,
		service: ps,
	}
}

func (pc *ProfileController) AddProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /api/v1/profiles")
		ctx, cancel := context.WithTimeout(ctf.Context(), TimeoutDuration)
		defer cancel()
		req := request.ProfileAddRequestDto{}
		if err := ctf.BodyParser(&req); err != nil {
			errorMessage := pc.getErrorMessage("AddProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		fmt.Println("ProfileAddRequestDto: ", req)
		profileResponse, err := pc.service.AddProfile(ctx, ctf, &req)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddProfile", "AddProfile")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) UpdateProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("PUT /api/v1/profiles")
		ctx, cancel := context.WithTimeout(ctf.Context(), TimeoutDuration)
		defer cancel()
		req := request.ProfileUpdateRequestDto{}
		if err := ctf.BodyParser(&req); err != nil {
			errorMessage := pc.getErrorMessage("UpdateProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		fmt.Println("ProfileUpdateRequestDto: ", req)
		profileResponse, err := pc.service.UpdateProfile(ctx, ctf, &req)
		if err != nil {
			errorMessage := pc.getErrorMessage("UpdateProfile", "UpdateProfile")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}
