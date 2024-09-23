package controller

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/controller/http/api/v1"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	errorFilePath = "internal/controller/profileController.go"
)

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
		req := &request.ProfileAddRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		profileResponse, err := pc.service.AddProfile(ctx, ctf, req)
		if err != nil {
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
		req := &request.ProfileUpdateRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("UpdateProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		profileResponse, err := pc.service.UpdateProfile(ctx, ctf, req)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) DeleteProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("DELETE /api/v1/profiles")
		ctx, cancel := context.WithTimeout(ctf.Context(), TimeoutDuration)
		defer cancel()
		req := &request.ProfileDeleteRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("DeleteProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		profileResponse, err := pc.service.DeleteProfile(ctx, req)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) GetProfileBySessionId() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /api/v1/profiles/session/:sessionId")
		ctx, cancel := context.WithTimeout(ctf.Context(), TimeoutDuration)
		defer cancel()
		req := &request.ProfileGetBySessionIdRequestDto{}
		if err := ctf.QueryParser(req); err != nil {
			errorMessage := pc.getErrorMessage("GetProfileBySessionId", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		sessionId := ctf.Params("sessionId")
		profileResponse, err := pc.service.GetProfileBySessionId(ctx, sessionId, req)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) GetProfileList() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /api/v1/profiles/list")
		ctx, cancel := context.WithTimeout(ctf.Context(), TimeoutDuration)
		defer cancel()
		req := &request.ProfileGetListRequestDto{}
		if err := ctf.QueryParser(req); err != nil {
			errorMessage := pc.getErrorMessage("GetProfileList", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		fmt.Println("GetProfileList req: ", req)
		profileListResponse, err := pc.service.GetProfileList(ctx, req)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileListResponse)
	}
}

func (pc *ProfileController) AddBlock() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /api/v1/profiles/blocks")
		ctx, cancel := context.WithTimeout(ctf.Context(), TimeoutDuration)
		defer cancel()
		req := &request.BlockRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddBlock", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		profileResponse, err := pc.service.AddBlock(ctx, req)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) AddLike() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /api/v1/profiles/likes")
		ctx, cancel := context.WithTimeout(ctf.Context(), TimeoutDuration)
		defer cancel()
		req := &request.LikeAddRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddLike", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		profileResponse, err := pc.service.AddLike(ctx, req)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) AddComplaint() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /api/v1/profiles/complaints")
		ctx, cancel := context.WithTimeout(ctf.Context(), TimeoutDuration)
		defer cancel()
		req := &request.ComplaintAddRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddComplaint", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		profileResponse, err := pc.service.AddComplaint(ctx, req)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}
