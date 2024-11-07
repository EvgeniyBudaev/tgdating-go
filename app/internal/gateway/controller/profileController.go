package controller

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	v1 "github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/controller/http/api/v1"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/logger"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enums"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/validation"
	"github.com/gofiber/fiber/v2"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"net/http"
	"strconv"
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
		locale := ctf.Get("Accept-Language")
		if locale == "" {
			locale = defaultLocale
		}
		req := &request.ProfileAddRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.SessionId); err != nil {
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		validateErr := validation.ValidateProfileAddRequestDto(ctf, req, locale)
		if validateErr != nil {
			return v1.ResponseFieldsError(ctf, validateErr)
		}
		md := metadata.New(map[string]string{"Accept-Language": locale})
		ctx = metadata.NewOutgoingContext(ctx, md)
		newTimestampBirthday := timestamppb.New(req.Birthday)
		form, err := ctf.MultipartForm()
		if err != nil {
			if err != nil {
				return v1.ResponseError(ctf, err, http.StatusInternalServerError)
			}
		}
		files := form.File["image"]
		fileList := make([]*pb.FileMetadata, 0)
		if len(files) > 0 {
			for _, file := range files {
				f, err := file.Open()
				if err != nil {
					return v1.ResponseError(ctf, err, http.StatusBadRequest)
				}
				data, err := io.ReadAll(f)
				if err != nil {
					return v1.ResponseError(ctf, err, http.StatusBadRequest)
				}
				fileList = append(fileList, &pb.FileMetadata{
					Filename: file.Filename,
					Size:     file.Size,
					Content:  data,
				})
			}
		}
		resp, err := pc.proto.Add(ctx, &pb.ProfileAddRequest{
			SessionId:               req.SessionId,
			DisplayName:             req.DisplayName,
			Birthday:                newTimestampBirthday,
			Gender:                  req.Gender,
			SearchGender:            req.SearchGender,
			Location:                req.Location,
			Description:             req.Description,
			Height:                  req.Height,
			Weight:                  req.Weight,
			TelegramUserId:          req.TelegramUserId,
			TelegramUsername:        req.TelegramUsername,
			TelegramFirstName:       req.TelegramFirstName,
			TelegramLastName:        req.TelegramLastName,
			TelegramLanguageCode:    req.TelegramLanguageCode,
			TelegramAllowsWriteToPm: req.TelegramAllowsWriteToPm,
			TelegramQueryId:         req.TelegramQueryId,
			Latitude:                req.Latitude,
			Longitude:               req.Longitude,
			AgeFrom:                 uint32(req.AgeFrom),
			AgeTo:                   uint32(req.AgeTo),
			Distance:                req.Distance,
			Page:                    req.Page,
			Size:                    req.Size,
			Files:                   fileList,
		})
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

func (pc *ProfileController) validateAuthUser(ctf *fiber.Ctx, sessionId string) error {
	telegramInitData, ok := ctf.UserContext().Value(enums.ContextKeyTelegram).(initdata.InitData)
	if !ok {
		err := errors.New("missing telegram data in context")
		return err
	}
	if sessionId != strconv.FormatInt(telegramInitData.User.ID, 10) {
		err := errors.New("unauthorized")
		return err
	}
	return nil
}
