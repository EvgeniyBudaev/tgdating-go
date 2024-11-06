package controller

import (
	"context"
	"fmt"
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	v1 "github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/controller/http/api/v1"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	defaultLocale   = "ru"
	errorFilePath   = "internal/gateway/controller/profileController.go"
	timeoutDuration = 30 * time.Second
)

type ProfileController struct {
	logger logger.Logger
	proto  pb.ProfileClient
}

func NewProfileController(l logger.Logger, pc pb.ProfileClient) *ProfileController {
	return &ProfileController{
		logger: l,
		proto:  pc,
	}
}

func (pc *ProfileController) AddProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /gateway/api/v1/profiles")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileAddRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		fmt.Println("gateway req.SessionId: ", req.SessionId)
		resp, err := pc.proto.Add(ctx, &pb.ProfileAddRequest{
			SessionId: req.SessionId,
		})
		fmt.Println("gateway resp: ", resp)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, resp)
	}
}

func (pc *ProfileController) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}
