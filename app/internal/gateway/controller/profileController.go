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
	"github.com/gofiber/fiber/v2"
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
		profileResponse := profileMapper.MapToBySessionIdResponse(profileUpdated)
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) DeleteProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("DELETE /gateway/api/v1/profiles")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileDeleteRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("DeleteProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToDeleteRequest(req)
		profileResponse, err := pc.proto.DeleteProfile(ctx, profileRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("DeleteProfile", "proto.DeleteProfile")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
	}
}

func (pc *ProfileController) RestoreProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /gateway/api/v1/profiles/restore")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileRestoreRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("RestoreProfile", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
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
			errorMessage := pc.getErrorMessage("GetProfileBySessionId",
				"proto.GetProfileBySessionId")
			pc.logger.Debug(errorMessage, zap.Error(err))
			if e, ok := status.FromError(err); ok {
				if e.Code() == codes.NotFound {
					return v1.ResponseError(ctf, err, http.StatusNotFound)
				}
				return v1.ResponseError(ctf, err, http.StatusInternalServerError)
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
		pc.logger.Info("GET /gateway/api/v1/profiles/short/:sessionId")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileGetShortInfoRequestDto{}
		if err := ctf.QueryParser(req); err != nil {
			errorMessage := pc.getErrorMessage("GetProfileShortInfo", "QueryParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		sessionId := ctf.Params("sessionId")
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToGetShortInfoRequest(req, sessionId)
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
		return v1.ResponseOk(ctf, profileShortInfo)
	}
}

func (pc *ProfileController) GetProfileList() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /gateway/api/v1/profiles/list")
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

func (pc *ProfileController) GetImageBySessionId() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /gateway/api/v1/profiles/:sessionId/images/:fileName")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		sessionId := ctf.Params("sessionId")
		fileName := ctf.Params("fileName")
		profileMapper := &mapper.ProfileMapper{}
		profileRequest := profileMapper.MapToImageBySessionIdRequest(sessionId, fileName)
		file, err := pc.proto.GetImageBySessionId(ctx, profileRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("GetImageBySessionId",
				"proto.GetImageBySessionId")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		fileResponse := profileMapper.MapToImageBySessionIdResponse(file)
		ctf.Set("Content-Type", "image/jpeg")
		return v1.ResponseImage(ctf, fileResponse)
	}
}

func (pc *ProfileController) DeleteImage() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("DELETE /gateway/api/v1/profiles/images/:id")
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
		sessionId := image.SessionId
		if err := pc.validateAuthUser(ctf, sessionId); err != nil {
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

func (pc *ProfileController) GetFilterBySessionId() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /gateway/api/v1/profiles/filter/:sessionId")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		sessionId := ctf.Params("sessionId")
		req := &request.FilterGetRequestDto{}
		if err := ctf.QueryParser(req); err != nil {
			errorMessage := pc.getErrorMessage("GetFilterBySessionId", "QueryParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		profileMapper := &mapper.ProfileMapper{}
		filterRequest := profileMapper.MapToFilterRequest(req, sessionId)
		filterResponse, err := pc.proto.GetFilterBySessionId(ctx, filterRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("GetFilterBySessionId",
				"proto.GetFilterBySessionId")
			pc.logger.Debug(errorMessage, zap.Error(err))
			if e, ok := status.FromError(err); ok {
				if e.Code() == codes.NotFound {
					return v1.ResponseError(ctf, err, http.StatusNotFound)
				}
				return v1.ResponseError(ctf, err, http.StatusInternalServerError)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseOk(ctf, filterResponse)
	}
}

func (pc *ProfileController) UpdateFilter() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("PUT /gateway/api/v1/profiles/filters")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.FilterUpdateRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("UpdateFilter", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.SessionId); err != nil {
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
		pc.logger.Info("POST /gateway/api/v1/profiles/blocks")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.BlockAddRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddBlock", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.SessionId); err != nil {
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
		blockResponse := profileMapper.MapToBlockAddResponse(blockAdded)
		return v1.ResponseCreated(ctf, blockResponse)
	}
}

func (pc *ProfileController) AddLike() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /gateway/api/v1/profiles/likes")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.LikeAddRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddLike", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.SessionId); err != nil {
			errorMessage := pc.getErrorMessage("AddLike", "validateAuthUser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		locale := ctf.Get("Accept-Language")
		if locale == "" {
			locale = defaultLocale
		}
		profileMapper := &mapper.ProfileMapper{}
		likeRequest := profileMapper.MapToLikeAddRequest(req, locale)
		likeAdded, err := pc.proto.AddLike(ctx, likeRequest)
		if err != nil {
			errorMessage := pc.getErrorMessage("AddLike", "proto.AddLike")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		likeResponse := profileMapper.MapToLikeAddResponse(likeAdded)
		return v1.ResponseCreated(ctf, likeResponse)
	}
}

func (pc *ProfileController) UpdateLike() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("PUT /gateway/api/v1/profiles/likes")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.LikeUpdateRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("UpdateLike", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.SessionId); err != nil {
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
		likeResponse := profileMapper.MapToLikeUpdateResponse(likeUpdated)
		return v1.ResponseCreated(ctf, likeResponse)
	}
}

func (pc *ProfileController) AddComplaint() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /gateway/api/v1/profiles/complaints")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ComplaintAddRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddComplaint", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.SessionId); err != nil {
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
		complaintResponse := profileMapper.MapToComplaintAddResponse(complaintAdded)
		return v1.ResponseCreated(ctf, complaintResponse)
	}
}

func (pc *ProfileController) UpdateCoordinates() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("PUT /gateway/api/v1/profiles/navigators")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.NavigatorUpdateRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("UpdateCoordinates", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.SessionId); err != nil {
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
