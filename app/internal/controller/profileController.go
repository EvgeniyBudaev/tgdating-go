package controller

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/controller/http/api/v1"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/repository/psql"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/shared/enums"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

const (
	defaultLocale   = "ru"
	errorFilePath   = "internal/controller/profileController.go"
	timeoutDuration = 30 * time.Second
)

type ProfileController struct {
	logger  logger.Logger
	service ProfileService
}

func NewProfileController(l logger.Logger, ps ProfileService) *ProfileController {
	return &ProfileController{
		logger:  l,
		service: ps,
	}
}

func (pc *ProfileController) AddProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /gateway/api/v1/profiles")
		//ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		//defer cancel()
		//locale := ctf.Get("Accept-Language")
		//if locale == "" {
		//	locale = defaultLocale
		//}
		//req := &request.ProfileAddRequestDto{}
		//if err := ctf.BodyParser(req); err != nil {
		//	errorMessage := pc.getErrorMessage("AddProfile", "BodyParser")
		//	pc.logger.Debug(errorMessage, zap.Error(err))
		//	return v1.ResponseError(ctf, err, http.StatusBadRequest)
		//}
		//if err := pc.validateAuthUser(ctf, req.SessionId); err != nil {
		//	return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		//}
		//validateErr := validation.ValidateProfileAddRequestDto(ctf, req, locale)
		//if validateErr != nil {
		//	return v1.ResponseFieldsError(ctf, validateErr)
		//}
		//profileResponse, err := pc.service.AddProfile(ctx, ctf, req)
		//if err != nil {
		//	return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		//}
		//return v1.ResponseCreated(ctf, profileResponse)
		return v1.ResponseCreated(ctf, "OK")
	}
}

func (pc *ProfileController) UpdateProfile() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("PUT /gateway/api/v1/profiles")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		acceptLanguage := ctf.Get("Accept-Language")
		if acceptLanguage == "" {
			acceptLanguage = defaultLocale
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
		validateErr := validation.ValidateProfileEditRequestDto(ctf, req, acceptLanguage)
		fmt.Println("UpdateProfile validateErr: ", validateErr)
		if validateErr != nil {
			return v1.ResponseFieldsError(ctf, validateErr)
		}
		profileResponse, err := pc.service.UpdateProfile(ctx, ctf, req)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		fmt.Println("UpdateProfile profileResponse: ", profileResponse)
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
		profileResponse, err := pc.service.DeleteProfile(ctx, req)
		if err != nil {
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
		profileResponse, err := pc.service.GetProfileBySessionId(ctx, sessionId, req)
		if err != nil {
			if errors.Is(err, psql.ErrNotRowFound) {
				return v1.ResponseError(ctf, err, http.StatusNotFound)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseOk(ctf, profileResponse)
	}
}

func (pc *ProfileController) GetProfileDetail() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("GET /gateway/api/v1/profiles/detail/:sessionId")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.ProfileGetDetailRequestDto{}
		if err := ctf.QueryParser(req); err != nil {
			errorMessage := pc.getErrorMessage("GetProfileDetail", "QueryParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		viewedSessionId := ctf.Params("viewedSessionId")
		profileResponse, err := pc.service.GetProfileDetail(ctx, viewedSessionId, req)
		if err != nil {
			if errors.Is(err, psql.ErrNotRowFound) {
				return v1.ResponseError(ctf, err, http.StatusNotFound)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
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
		profileResponse, err := pc.service.GetProfileShortInfo(ctx, sessionId, req)
		if err != nil {
			if errors.Is(err, psql.ErrNotRowFound) {
				return v1.ResponseError(ctf, err, http.StatusNotFound)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseOk(ctf, profileResponse)
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
		profileListResponse, err := pc.service.GetProfileList(ctx, req)
		if err != nil {
			if errors.Is(err, psql.ErrNotRowFound) {
				return v1.ResponseError(ctf, err, http.StatusNotFound)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
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
		response, err := pc.service.GetImageBySessionId(ctx, sessionId, fileName)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		ctf.Set("Content-Type", "image/jpeg")
		return v1.ResponseImage(ctf, response)
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
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		image, err := pc.service.GetImageById(ctx, idUint64)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		sessionId := image.SessionId
		if err := pc.validateAuthUser(ctf, sessionId); err != nil {
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		response, err := pc.service.DeleteImage(ctx, idUint64)
		if err != nil {
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
		profileListResponse, err := pc.service.GetFilterBySessionId(ctx, sessionId, req)
		if err != nil {
			if errors.Is(err, psql.ErrNotRowFound) {
				return v1.ResponseError(ctf, err, http.StatusNotFound)
			}
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseOk(ctf, profileListResponse)
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
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileListResponse, err := pc.service.UpdateFilter(ctx, req)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileListResponse)
	}
}

func (pc *ProfileController) AddBlock() fiber.Handler {
	return func(ctf *fiber.Ctx) error {
		pc.logger.Info("POST /gateway/api/v1/profiles/blocks")
		ctx, cancel := context.WithTimeout(ctf.Context(), timeoutDuration)
		defer cancel()
		req := &request.BlockRequestDto{}
		if err := ctf.BodyParser(req); err != nil {
			errorMessage := pc.getErrorMessage("AddBlock", "BodyParser")
			pc.logger.Debug(errorMessage, zap.Error(err))
			return v1.ResponseError(ctf, err, http.StatusBadRequest)
		}
		if err := pc.validateAuthUser(ctf, req.SessionId); err != nil {
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
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
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		locale := ctf.Get("Accept-Language")
		if locale == "" {
			locale = defaultLocale
		}
		profileResponse, err := pc.service.AddLike(ctx, req, locale)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
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
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileResponse, err := pc.service.UpdateLike(ctx, req)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
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
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileResponse, err := pc.service.AddComplaint(ctx, req)
		if err != nil {
			return v1.ResponseError(ctf, err, http.StatusInternalServerError)
		}
		return v1.ResponseCreated(ctf, profileResponse)
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
			return v1.ResponseError(ctf, err, http.StatusUnauthorized)
		}
		profileResponse, err := pc.service.UpdateCoordinates(ctx, req)
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
