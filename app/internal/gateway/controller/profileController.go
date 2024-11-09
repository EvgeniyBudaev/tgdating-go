package controller

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	v1 "github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/controller/http/api/v1"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/controller/mapper"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/logger"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enums"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/validation"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/repository/psql"
	"github.com/gofiber/fiber/v2"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
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
		fileList, err := pc.getFiles(ctf)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToAddRequest(req, fileList)
		resp, err := pc.proto.AddProfile(ctx, profileRequest)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, resp)
	}
}

func (pc *ProfileController) UpdateProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("PUT /gateway/api/v1/profiles")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		locale := ctf.Get("Accept-Language")
		if locale == "" {
			locale = defaultLocale
		}
		req := &request.ProfileUpdateRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("UpdateProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.SessionId); err != nil {
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		validateErr := validation.ValidateProfileEditRequestDto(ctf, req, locale)
		if validateErr != nil {
			return v1.ResponseFieldsError(ctf, validateErr)
		}
		fileList, err := pc.getFiles(ctf)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToUpdateRequest(req, fileList)
		profileUpdated, err := pc.proto.UpdateProfile(ctx, profileRequest)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		profileResponse := profileMapper.MapToBySessionIdResponse(profileUpdated)
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) GetProfileBySessionId() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /gateway/api/v1/profiles/session/:sessionId")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileGetBySessionIdRequestDto{}
		if err := ctf.QueryParser(req); err != nil {
			errorMessage := pc.getErrorMessage("GetProfileBySessionId", "QueryParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		sessionId := ctf.Params("sessionId")
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToGetBySessionIdRequest(req, sessionId)
		profileBySessionId, err := pc.proto.GetProfileBySessionId(ctx, profileRequest)
		if err != nil {
			if errors.Is(err, psql.ErrNotRowFound) {
				return v1.ResponseError(ctf, err, http.StatusNotFound)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		profileResponse := profileMapper.MapToBySessionIdResponse(profileBySessionId)
		return v1.ResponseOk(ctf, profileResponse)
	}
}

func (pc *ProfileController) GetProfileDetail() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /gateway/api/v1/profiles/detail/:viewedSessionId")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileGetDetailRequestDto{}
		if err := ctf.QueryParser(req); err != nil {
			errorMessage := pc.getErrorMessage("GetProfileDetail", "QueryParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		viewedSessionId := ctf.Params("viewedSessionId")
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToGetDetailRequest(req, viewedSessionId)
		profileDetail, err := pc.proto.GetProfileDetail(ctx, profileRequest)
		if err != nil {
			if errors.Is(err, psql.ErrNotRowFound) {
				return v1.ResponseError(ctf, err, http.StatusNotFound)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		fmt.Println("profileDetail.SessionId: ", profileDetail.SessionId)
		profileResponse := profileMapper.MapToDetailResponse(profileDetail)
		return v1.ResponseOk(ctf, profileResponse)
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

func (pc *ProfileController) getFiles(ctf *fiber.Ctx) ([]*pb.FileMetadata, error) {
	form, err := ctf.MultipartForm()
	if err != nil {
		errorMessage := pc.getErrorMessage("getFiles", "ctf.MultipartForm")
		pc.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	files := form.File["image"]
	fileList := make([]*pb.FileMetadata, 0)
	if len(files) > 0 {
		for _, file := range files {
			f, err := file.Open()
			if err != nil {
				errorMessage := pc.getErrorMessage("getFiles", "file.Open")
				pc.logger.Debug(errorMessage, zap.Error(err))
				return nil, err
			}
			data, err := io.ReadAll(f)
			if err != nil {
				errorMessage := pc.getErrorMessage("getFiles", "io.ReadAll")
				pc.logger.Debug(errorMessage, zap.Error(err))
				return nil, err
			}
			fileList = append(fileList, &pb.FileMetadata{
				Filename: file.Filename,
				Size:     file.Size,
				Content:  data,
			})
		}
	}
	return fileList, nil
}
