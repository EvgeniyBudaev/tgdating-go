package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	v1 "github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/controller/http/api/v1"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/controller/mapper"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/logger"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enum"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	defaultLocale   = "ru"
	errorFilePath   = "internal/gateway/controller/profile-controller.go"
	timeoutDuration = 30 * time.Second
)

type ProfileController struct {
	logger      logger.Logger
	kafkaWriter *kafka.Writer
	proto       pb.ProfileClient
}

func NewProfileController(l logger.Logger, kw *kafka.Writer, pc pb.ProfileClient) *ProfileController {
	return &ProfileController{
		logger:      l,
		kafkaWriter: kw,
		proto:       pc,
	}
}

func (pc *ProfileController) AddProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /api/v1/profiles")
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
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("AddProfile", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		validateErr := validation.ValidateProfileAddRequestDto(ctf, req, locale)
		if validateErr != nil {
			errorMessage := pc.getErrorMessage("AddProfile",
				"ValidateProfileAddRequestDto")
			pc.logger.Debug(errorMessage)
			return v1.ResponseFieldsError(ctf, validateErr)
		}
		md := metadata.New(map[string]string{"Accept-Language": locale})
		ctx = metadata.NewOutgoingContext(ctx, md)
		fileList, err := pc.getFiles(ctf)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddProfile", "getFiles")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToAddRequest(req, fileList)
		resp, err := pc.proto.AddProfile(ctx, profileRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddProfile", "proto.AddProfile")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, resp)
	}
}

func (pc *ProfileController) UpdateProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("PUT /api/v1/profiles")
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
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("UpdateProfile", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		validateErr := validation.ValidateProfileEditRequestDto(ctf, req, locale)
		if validateErr != nil {
			errorMessage := pc.getErrorMessage("UpdateProfile",
				"ValidateProfileEditRequestDto")
			pc.logger.Debug(errorMessage)
			return v1.ResponseFieldsError(ctf, validateErr)
		}
		fileList, err := pc.getFiles(ctf)
		if err != nil {
			errorMessage := pc.getErrorMessage("UpdateProfile", "getFiles")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToUpdateRequest(req, fileList)
		profileUpdated, err := pc.proto.UpdateProfile(ctx, profileRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("UpdateProfile", "proto.UpdateProfile")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		profileResponse := profileMapper.MapToByTelegramUserIdResponse(profileUpdated)
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) FreezeProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /api/v1/profiles/freeze")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileFreezeRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("FreezeProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("FreezeProfile", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToFreezeRequest(req)
		profileResponse, err := pc.proto.FreezeProfile(ctx, profileRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("FreezeProfile", "proto.FreezeProfile")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) RestoreProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /api/v1/profiles/restore")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileRestoreRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("RestoreProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("RestoreProfile", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToRestoreRequest(req)
		profileResponse, err := pc.proto.RestoreProfile(ctx, profileRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("RestoreProfile",
				"proto.RestoreProfile")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) DeleteProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("DELETE /api/v1/profiles")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileDeleteRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("DeleteProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("DeleteProfile", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToDeleteRequest(req)
		profileResponse, err := pc.proto.DeleteProfile(ctx, profileRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("DeleteProfile",
				"proto.DeleteProfile")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) GetProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /api/v1/profiles/telegram/:telegramUserId")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileGetByTelegramUserIdRequestDto{}
		if err := ctf.QueryParser(req); err != nil {
			errorMessage := pc.getErrorMessage("GetProfile", "QueryParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		telegramUserId := ctf.Params("telegramUserId")
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToGetRequest(req, telegramUserId)
		profileByTelegramUserId, err := pc.proto.GetProfile(ctx, profileRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("GetProfile", "proto.GetProfile")
			pc.logger.Debug(errorMessage, zap.Error(err))
			if e, ok := status.FromError(err); ok {
				if e.Code() == codes.NotFound {
					return v1.ResponseError(ctf, err, http.StatusNotFound)
				}
				return v1.ResponseError(ctf, err, http.StatusInternalServerError)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		profileResponse := profileMapper.MapToByTelegramUserIdResponse(profileByTelegramUserId)
		return v1.ResponseOk(ctf, profileResponse)
	}
}

func (pc *ProfileController) GetProfileDetail() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /api/v1/profiles/detail/:viewedTelegramUserId")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileGetDetailRequestDto{}
		if err := ctf.QueryParser(req); err != nil {
			errorMessage := pc.getErrorMessage("GetProfileDetail", "QueryParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		viewedTelegramUserId := ctf.Params("viewedTelegramUserId")
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToGetDetailRequest(req, viewedTelegramUserId)
		profileDetail, err := pc.proto.GetProfileDetail(ctx, profileRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("GetProfileDetail",
				"proto.GetProfileDetail")
			pc.logger.Debug(errorMessage, zap.Error(err))
			if e, ok := status.FromError(err); ok {
				if e.Code() == codes.NotFound {
					return v1.ResponseError(ctf, err, http.StatusNotFound)
				}
				return v1.ResponseError(ctf, err, http.StatusInternalServerError)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		profileResponse := profileMapper.MapToDetailResponse(profileDetail)
		return v1.ResponseOk(ctf, profileResponse)
	}
}

func (pc *ProfileController) GetProfileShortInfo() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /api/v1/profiles/short/:telegramUserId")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		telegramUserId := ctf.Params("telegramUserId")
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToGetShortInfoRequest(telegramUserId)
		profileShortInfo, err := pc.proto.GetProfileShortInfo(ctx, profileRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("GetProfileShortInfo",
				"GetProfileShortInfo")
			pc.logger.Debug(errorMessage, zap.Error(err))
			if e, ok := status.FromError(err); ok {
				if e.Code() == codes.NotFound {
					return v1.ResponseError(ctf, err, http.StatusNotFound)
				}
				return v1.ResponseError(ctf, err, http.StatusInternalServerError)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		profileResponse := profileMapper.MapToShortInfoResponse(profileShortInfo)
		return v1.ResponseOk(ctf, profileResponse)
	}
}

func (pc *ProfileController) GetProfileList() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /api/v1/profiles/list")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileGetListRequestDto{}
		if err := ctf.QueryParser(req); err != nil {
			errorMessage := pc.getErrorMessage("GetProfileList", "QueryParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToListRequest(req)
		profileList, err := pc.proto.GetProfileList(ctx, profileRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("GetProfileList",
				"proto.GetProfileList")
			pc.logger.Debug(errorMessage, zap.Error(err))
			if e, ok := status.FromError(err); ok {
				if e.Code() == codes.NotFound {
					return v1.ResponseError(ctf, err, http.StatusNotFound)
				}
				return v1.ResponseError(ctf, err, http.StatusInternalServerError)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		profileListResponse := profileMapper.MapToListResponse(profileList)
		return v1.ResponseOk(ctf, profileListResponse)
	}
}

func (pc *ProfileController) CheckProfileExists() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /api/v1/profiles/:telegramUserId/check")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		telegramUserId := ctf.Params("telegramUserId")
		profileMapper := &mapper.ProfileMapper{}
		checkProfileExistsRequest := profileMapper.MapToCheckProfileExistsRequest(telegramUserId)
		_, err := pc.proto.CheckProfileExists(ctx, checkProfileExistsRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("CheckProfileExists",
				"proto.CheckProfileExists")
			pc.logger.Debug(errorMessage, zap.Error(err))
			if e, ok := status.FromError(err); ok {
				if e.Code() == codes.NotFound {
					//return v1.ResponseError(ctf, err, http.StatusNotFound)
					checkProfileExistsResponse := profileMapper.MapToCheckProfileExistsResponse(false)
					return v1.ResponseOk(ctf, checkProfileExistsResponse)
				}
				return v1.ResponseError(ctf, err, http.StatusInternalServerError)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		checkProfileExistsResponse := profileMapper.MapToCheckProfileExistsResponse(true)
		return v1.ResponseOk(ctf, checkProfileExistsResponse)
	}
}

func (pc *ProfileController) GetImageByTelegramUserId() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /api/v1/profiles/:telegramUserId/images/:fileName")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		telegramUserId := ctf.Params("telegramUserId")
		fileName := ctf.Params("fileName")
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToImageByTelegramUserIdRequest(telegramUserId, fileName)
		file, err := pc.proto.GetImageByTelegramUserId(ctx, profileRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("GetImageByTelegramUserId",
				"proto.GetImageByTelegramUserId")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		fileResponse := profileMapper.MapToImageByTelegramUserIdResponse(file)
		ctf.Set("Content-Type", "image/jpeg")
		return v1.ResponseImage(ctf, fileResponse)
	}
}

func (pc *ProfileController) DeleteImage() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("DELETE /api/v1/profiles/images/:id")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		id := ctf.Params("id")
		idUint64, err := pc.convertToUint64("id", id)
		if err != nil {
			errorMessage := pc.getErrorMessage("DeleteImage", "convertToUint64")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		imageByIdRequest := &pb.GetImageByIdRequest{
			Id: idUint64,
		}
		image, err := pc.proto.GetImageById(ctx, imageByIdRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("DeleteImage", "proto.GetImageById")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		if err := pc.validateAuthUser(ctf, image.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("DeleteImage", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		req := &pb.ImageDeleteRequest{
			Id: idUint64,
		}
		response, err := pc.proto.DeleteImage(ctx, req)
		if err != nil {
			errorMessage := pc.getErrorMessage("DeleteImage", "proto.DeleteImage")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, response)
	}
}

func (pc *ProfileController) GetFilter() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /api/v1/profiles/filters/:telegramUserId")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		telegramUserId := ctf.Params("telegramUserId")
		profileMapper := &mapper.ProfileMapper{}
		filterRequest := profileMapper.MapToFilterGetRequest(telegramUserId)
		filterResponse, err := pc.proto.GetFilter(ctx, filterRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("GetFilter", "proto.GetFilter")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, filterResponse)
	}
}

func (pc *ProfileController) UpdateFilter() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("PUT /api/v1/profiles/filters")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.FilterUpdateRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("UpdateFilter", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("UpdateFilter", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileMapper := &mapper.ProfileMapper{}
		filterRequest := profileMapper.MapToFilterUpdateRequest(req)
		filterResponse, err := pc.proto.UpdateFilter(ctx, filterRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("UpdateFilter", "proto.UpdateFilter")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, filterResponse)
	}
}

func (pc *ProfileController) AddBlock() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /api/v1/profiles/blocks")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.BlockAddRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddBlock", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("AddBlock", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileMapper := &mapper.ProfileMapper{}
		blockRequest := profileMapper.MapToBlockAddRequest(req)
		blockAdded, err := pc.proto.AddBlock(ctx, blockRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddBlock", "proto.AddBlock")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, blockAdded)
	}
}

func (pc *ProfileController) GetBlockedList() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /api/v1/profiles/:telegramUserId/blocks/list")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		telegramUserId := ctf.Params("telegramUserId")
		profileMapper := &mapper.ProfileMapper{}
		blockedListRequest := profileMapper.MapToGetBlockedListRequest(telegramUserId)
		blockedListResponse, err := pc.proto.GetBlockedList(ctx, blockedListRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("GetBlockedList", "proto.GetBlockedList")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, blockedListResponse)
	}
}

func (pc *ProfileController) Unblock() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("PUT /api/v1/profiles/unblock")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.UnblockRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("Unblock", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("Unblock", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileMapper := &mapper.ProfileMapper{}
		unblockRequest := profileMapper.MapToUnblockRequest(req)
		unblockResponse, err := pc.proto.Unblock(ctx, unblockRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("Unblock", "proto.Unblock")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, unblockResponse)
	}
}

func (pc *ProfileController) AddLike() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /api/v1/profiles/likes")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.LikeAddRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddLike", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("AddLike", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		locale := ctf.Get("Accept-Language")
		if locale == "" {
			locale = defaultLocale
		}
		profileMapper := &mapper.ProfileMapper{}
		telegramRequest := profileMapper.MapToTelegramGetRequest(req.TelegramUserId)
		telegramProfile, err := pc.proto.GetTelegram(ctx, telegramRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddLike",
				"proto.GetTelegram")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		likedTelegramRequest := profileMapper.MapToTelegramGetRequest(req.LikedTelegramUserId)
		likedTelegramProfile, err := pc.proto.GetTelegram(ctx, likedTelegramRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddLike",
				"proto.GetTelegram likedTelegramRequest")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		imageRequest := profileMapper.MapToGetImageLastRequest(req.TelegramUserId)
		lastImage, err := pc.proto.GetImageLastByTelegramUserId(ctx, imageRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddLike",
				"proto.GetImageLastByTelegramUserId")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		hc := &entity.HubContent{
			LikedTelegramUserId: req.LikedTelegramUserId,
			Message:             pc.GetMessageLike(likedTelegramProfile.LanguageCode),
			Type:                "like",
			UserImageUrl:        lastImage.Url,
			Username:            telegramProfile.Username,
		}
		hubContentJson, err := json.Marshal(hc)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddLike", "json.Marshal")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		err = pc.kafkaWriter.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(req.LikedTelegramUserId),
				Value: hubContentJson,
			},
		)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddLike", "kafkaWriter.WriteMessages")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		likeRequest := profileMapper.MapToLikeAddRequest(req, locale)
		likeAdded, err := pc.proto.AddLike(ctx, likeRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddLike", "proto.AddLike")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, likeAdded)
	}
}

func (pc *ProfileController) UpdateLike() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("PUT /api/v1/profiles/likes")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.LikeUpdateRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("UpdateLike", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("UpdateLike", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileMapper := &mapper.ProfileMapper{}
		likeRequest := profileMapper.MapToLikeUpdateRequest(req)
		likeUpdated, err := pc.proto.UpdateLike(ctx, likeRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("UpdateLike", "proto.UpdateLike")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, likeUpdated)
	}
}

func (pc *ProfileController) GetLastLike() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /api/v1/profiles/likes/last")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.LikeGetLastRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("GetLastLike", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		fmt.Println("GetLastLike TelegramUserId:", req.TelegramUserId)
		profileMapper := &mapper.ProfileMapper{}
		likeRequest := profileMapper.MapToLikeGetLastRequest(req.TelegramUserId)
		likeEntity, err := pc.proto.GetLastLike(ctx, likeRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("GetLastLike", "proto.GetLastLike")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		likeResponse := profileMapper.MapToLikeGetLastResponse(likeEntity)
		return v1.ResponseCreated(ctf, likeResponse)
	}
}

func (pc *ProfileController) AddComplaint() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /api/v1/profiles/complaints")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ComplaintAddRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddComplaint", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("AddComplaint", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileMapper := &mapper.ProfileMapper{}
		complaintRequest := profileMapper.MapToComplaintAddRequest(req)
		complaintAdded, err := pc.proto.AddComplaint(ctx, complaintRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddComplaint", "proto.AddComplaint")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, complaintAdded)
	}
}

func (pc *ProfileController) AddPayment() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /api/v1/profiles/payments")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.PaymentAddRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddPayment", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("AddPayment", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileMapper := &mapper.ProfileMapper{}
		paymentRequest := profileMapper.MapToPaymentAddRequest(req)
		paymentAdded, err := pc.proto.AddPayment(ctx, paymentRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddPayment", "proto.AddPayment")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, paymentAdded)
	}
}

func (pc *ProfileController) CheckPremium() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /api/v1/profiles/:telegramUserId/premium/check")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		telegramUserId := ctf.Params("telegramUserId")
		profileMapper := &mapper.ProfileMapper{}
		checkIsPremiumRequest := profileMapper.MapToCheckPremiumRequest(telegramUserId)
		checkIsPremium, err := pc.proto.CheckPremium(ctx, checkIsPremiumRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("CheckPremium", "proto.CheckPremium")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		checkIsPremiumResponse := profileMapper.MapToCheckPremiumResponse(checkIsPremium)
		return v1.ResponseOk(ctf, checkIsPremiumResponse)
	}
}

func (pc *ProfileController) UpdateSettings() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("PUT /api/v1/profiles/settings")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileUpdateSettingsRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("UpdateSettings", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("UpdateSettings", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileMapper := &mapper.ProfileMapper{}
		updateSettingsRequest := profileMapper.MapToUpdateSettingsRequest(req)
		updateSettingsResponse, err := pc.proto.UpdateSettings(ctx, updateSettingsRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("UpdateSettings",
				"proto.UpdateSettings")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, updateSettingsResponse)
	}
}

func (pc *ProfileController) UpdateCoordinates() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("PUT /api/v1/profiles/navigators")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.NavigatorUpdateRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("UpdateCoordinates", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.TelegramUserId); err != nil {
			errorMessage := pc.getErrorMessage("UpdateCoordinates", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileMapper := &mapper.ProfileMapper{}
		updateCoordinatesRequest := profileMapper.MapToUpdateCoordinatesRequest(req)
		updateCoordinatesResponse, err := pc.proto.UpdateCoordinates(ctx, updateCoordinatesRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("UpdateCoordinates",
				"proto.UpdateCoordinates")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, updateCoordinatesResponse)
	}
}

func (pc *ProfileController) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}

func (pc *ProfileController) validateAuthUser(ctf *fiber.Ctx, telegramUserId string) error {
	telegramInitData, ok := ctf.UserContext().Value(enum.ContextKeyTelegram).(initdata.InitData)
	if !ok {
		err := errors.New("missing telegram data in context")
		return err
	}
	if telegramUserId != strconv.FormatInt(telegramInitData.User.ID, 10) {
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

func (pc *ProfileController) GetMessageLike(locale string) string {
	switch locale {
	case "ru":
		return "Есть симпатия! Начинай общаться"
	case "en":
		return "There is sympathy! Start communicating"
	case "ar":
		return "هناك تعاطف! ابدأ بالتواصل"
	case "be":
		return "Ёсць сімпатыя! Пачынай мець зносіны"
	case "ca":
		return "Hi ha simpatia! Comença a comunicar-te"
	case "cs":
		return "Jsou tam sympatie! Začněte komunikovat"
	case "de":
		return "Es gibt Mitgefühl! Beginnen Sie mit der kommunikation"
	case "es":
		return "¡Hay simpatía! Empezar a comunicar"
	case "fi":
		return "Sympatiaa on! Aloita kommunikointi"
	case "fr":
		return "Il y a de la sympathie ! Commencez à communiquer"
	case "he":
		return "יש סימפטיה! תתחיל לתקשר"
	case "hr":
		return "Postoji simpatija! Počnite komunicirati"
	case "hu":
		return "Együttérzés van! Kezdj el kommunikálni"
	case "id":
		return "Ada simpati! Mulailah berkomunikasi"
	case "it":
		return "C'è simpatia! Inizia a comunicare"
	case "ja":
		return "共感があるよ！通信を開始する"
	case "kk":
		return "Жанашырлық бар! Қарым-қатынасты бастаңыз"
	case "ko":
		return "동정심이 있습니다! 소통을 시작해 보세요"
	case "nl":
		return "Er is sympathie! Begin met communiceren"
	case "no":
		return "Det er sympati! Begynn å kommunisere"
	case "pt":
		return "Existe simpatia! Comece a se comunicar"
	case "sv":
		return "Det finns sympati! Börja kommunicera"
	case "uk":
		return "Є симпатія! Починай спілкуватися"
	case "zh":
		return "有同情心！开始沟通"
	default:
		return "There is sympathy! Start communicating"
	}
}

func (pc *ProfileController) convertToUint64(name, value string) (uint64, error) {
	if value == "" {
		errorMessage := fmt.Sprintf("%s is empty", name)
		return 0, errors.New(errorMessage)
	}
	value64, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		errorMessage := pc.getErrorMessage("convertToUint64", "ParseUint")
		pc.logger.Debug(errorMessage, zap.Error(err))
		return 0, err
	}
	return value64, nil
}
