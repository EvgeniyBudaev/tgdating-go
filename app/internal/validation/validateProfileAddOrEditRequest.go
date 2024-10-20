package validation

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"github.com/gofiber/fiber/v2"
	"time"
)

var (
	fileMaxSizeMb    = 10
	fileMaxSizeBytes = int64(fileMaxSizeMb * 1024 * 1024)
	fileMaxAmount    = 3
)

func ValidateProfileAddOrEditRequestDto(ctf *fiber.Ctx, req *request.ProfileAddRequestDto,
	locale string) *entity.ValidationErrorEntity {
	errorMessages := entity.NewErrorMessagesEntity()
	message := errorMessages.GetBadRequest(locale)
	fieldErrorsLanguages := map[string][]string{}

	if req.SessionId == "" {
		fieldErrorsLanguages["sessionId"] = append(fieldErrorsLanguages["sessionId"],
			errorMessages.GetNotEmpty(locale))
	}

	if req.DisplayName == "" {
		fieldErrorsLanguages["displayName"] = append(fieldErrorsLanguages["displayName"],
			errorMessages.GetNotEmpty(locale))
	}

	if len(req.DisplayName) > 64 {
		fieldErrorsLanguages["displayName"] = append(fieldErrorsLanguages["displayName"],
			errorMessages.GetMaxSymbols(locale, 64))
	}

	birthday := req.Birthday.Format("2006-01-02")
	now := time.Now()
	eighteenYearsAgo := now.AddDate(-18, 0, 1)
	hundredYearsAgo := now.AddDate(-100, 0, 0)

	if birthday == "" {
		fieldErrorsLanguages["birthday"] = append(fieldErrorsLanguages["birthday"],
			errorMessages.GetNotEmpty(locale))
	}

	if req.Birthday.After(eighteenYearsAgo) {
		fieldErrorsLanguages["birthday"] = append(fieldErrorsLanguages["birthday"],
			errorMessages.GetMoreOrEqualMinNumber(locale, 18))
	}

	if req.Birthday.Before(hundredYearsAgo) {
		fieldErrorsLanguages["birthday"] = append(fieldErrorsLanguages["birthday"],
			errorMessages.GetLessOrEqualMaxNumber(locale, 100))
	}

	if req.Gender == "" {
		fieldErrorsLanguages["gender"] = append(fieldErrorsLanguages["gender"],
			errorMessages.GetNotEmpty(locale))
	}

	if req.Gender != "" && req.Gender != "man" && req.Gender != "woman" {
		fieldErrorsLanguages["gender"] = append(fieldErrorsLanguages["gender"],
			errorMessages.GetBadRequest(locale))
	}

	if req.SearchGender != "" && req.SearchGender != "man" && req.SearchGender != "woman" && req.SearchGender != "all" {
		fieldErrorsLanguages["searchGender"] = append(fieldErrorsLanguages["searchGender"],
			errorMessages.GetBadRequest(locale))
	}

	if len(req.Description) > 1000 {
		fieldErrorsLanguages["description"] = append(fieldErrorsLanguages["description"],
			errorMessages.GetMaxSymbols(locale, 1000))
	}

	if req.Height < 0 {
		fieldErrorsLanguages["height"] = append(fieldErrorsLanguages["height"],
			errorMessages.GetNonNegativeNumber(locale))
	}

	if req.Height > 250 {
		fieldErrorsLanguages["height"] = append(fieldErrorsLanguages["height"],
			errorMessages.GetLessOrEqualMaxNumber(locale, 250))
	}

	if req.Weight < 0 {
		fieldErrorsLanguages["weight"] = append(fieldErrorsLanguages["weight"],
			errorMessages.GetNonNegativeNumber(locale))
	}

	if req.Weight > 650 {
		fieldErrorsLanguages["weight"] = append(fieldErrorsLanguages["weight"],
			errorMessages.GetLessOrEqualMaxNumber(locale, 650))
	}

	if req.TelegramUserId == 0 {
		fieldErrorsLanguages["telegramUserId"] = append(fieldErrorsLanguages["telegramUserId"],
			errorMessages.GetNotEmpty(locale))
	}

	if req.TelegramUsername == "" {
		fieldErrorsLanguages["telegramUsername"] = append(fieldErrorsLanguages["telegramUsername"],
			errorMessages.GetNotEmpty(locale))
	}

	if req.AgeFrom < 18 {
		fieldErrorsLanguages["ageFrom"] = append(fieldErrorsLanguages["ageFrom"],
			errorMessages.GetMoreOrEqualMinNumber(locale, 18))
	}

	if req.AgeTo > 100 {
		fieldErrorsLanguages["ageTo"] = append(fieldErrorsLanguages["ageTo"],
			errorMessages.GetLessOrEqualMaxNumber(locale, 100))
	}

	if req.Distance < 0 {
		fieldErrorsLanguages["distance"] = append(fieldErrorsLanguages["distance"],
			errorMessages.GetNonNegativeNumber(locale))
	}

	if req.Page < 0 {
		fieldErrorsLanguages["page"] = append(fieldErrorsLanguages["page"],
			errorMessages.GetNonNegativeNumber(locale))
	}

	if req.Size < 0 {
		fieldErrorsLanguages["size"] = append(fieldErrorsLanguages["size"],
			errorMessages.GetNonNegativeNumber(locale))
	}

	form, err := ctf.MultipartForm()

	if err != nil {
		fieldErrorsLanguages["image"] = append(fieldErrorsLanguages["image"],
			errorMessages.GetNotEmpty(locale))
	}

	files := form.File["image"]

	if len(files) <= 0 {
		fieldErrorsLanguages["image"] = append(fieldErrorsLanguages["image"],
			errorMessages.GetNotEmpty(locale))
	}

	if len(files) > fileMaxAmount {
		fieldErrorsLanguages["image"] = append(fieldErrorsLanguages["image"],
			errorMessages.GetFileMaxAmount(locale, fileMaxAmount))
	}

	for _, file := range files {
		fileSize := file.Size
		if fileSize > fileMaxSizeBytes {
			fieldErrorsLanguages["image"] = append(fieldErrorsLanguages["image"],
				errorMessages.GetFileMaxSize(locale, fileMaxSizeMb))
			break
		}
	}

	fieldErrors := map[string][]*entity.FieldErrorEntity{}

	for key, errors := range fieldErrorsLanguages {
		for _, err := range errors {
			fieldErrors[key] = append(fieldErrors[key], &entity.FieldErrorEntity{
				Message: err,
			})
		}
	}

	vee := entity.NewValidationErrorEntity(message, fieldErrors)

	if len(fieldErrors) > 0 {
		return vee
	}

	return nil
}
