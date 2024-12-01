package validation

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/entity"
	"github.com/gofiber/fiber/v2"
	"time"
)

func ValidateProfileEditRequestDto(ctf *fiber.Ctx, req *request.ProfileUpdateRequestDto,
	locale string) *entity.ValidationErrorEntity {
	errorMessages := entity.NewErrorMessagesEntity()
	message := errorMessages.GetBadRequest(locale)
	fieldErrorsLanguages := map[string][]string{}

	if req.TelegramUserId == "" {
		fieldErrorsLanguages["telegramUserId"] = append(fieldErrorsLanguages["telegramUserId"],
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

	if req.Gender != "" && !req.Gender.IsValid() {
		fieldErrorsLanguages["gender"] = append(fieldErrorsLanguages["gender"],
			errorMessages.GetBadRequest(locale))
	}

	if req.SearchGender == "" {
		fieldErrorsLanguages["searchGender"] = append(fieldErrorsLanguages["searchGender"],
			errorMessages.GetNotEmpty(locale))
	}

	if req.SearchGender != "" && !req.SearchGender.IsValid() {
		fieldErrorsLanguages["searchGender"] = append(fieldErrorsLanguages["searchGender"],
			errorMessages.GetBadRequest(locale))
	}

	if req.LookingFor == "" {
		fieldErrorsLanguages["lookingFor"] = append(fieldErrorsLanguages["lookingFor"],
			errorMessages.GetNotEmpty(locale))
	}

	if req.LookingFor != "" && !req.LookingFor.IsValid() {
		fieldErrorsLanguages["lookingFor"] = append(fieldErrorsLanguages["lookingFor"],
			errorMessages.GetBadRequest(locale))
	}

	if req.Description != "" && len(req.Description) > maxCharacters {
		fieldErrorsLanguages["description"] = append(fieldErrorsLanguages["description"],
			errorMessages.GetMaxSymbols(locale, maxCharacters))
	}

	if req.Height != 0 && req.Height < 0 {
		fieldErrorsLanguages["height"] = append(fieldErrorsLanguages["height"],
			errorMessages.GetNonNegativeNumber(locale))
	}

	if req.Height != 0 && req.Height > 0 && int(req.Height) < minHeight {
		fieldErrorsLanguages["height"] = append(fieldErrorsLanguages["height"],
			errorMessages.GetMoreOrEqualMinNumber(locale, minHeight))
	}

	if req.Height != 0 && int(req.Height) > maxHeight {
		fieldErrorsLanguages["height"] = append(fieldErrorsLanguages["height"],
			errorMessages.GetLessOrEqualMaxNumber(locale, maxHeight))
	}

	if req.Weight != 0 && req.Weight < 0 {
		fieldErrorsLanguages["weight"] = append(fieldErrorsLanguages["weight"],
			errorMessages.GetNonNegativeNumber(locale))
	}

	if req.Weight != 0 && req.Weight > 0 && int(req.Weight) < minWeight {
		fieldErrorsLanguages["weight"] = append(fieldErrorsLanguages["weight"],
			errorMessages.GetMoreOrEqualMinNumber(locale, minWeight))
	}

	if req.Weight != 0 && int(req.Weight) > maxWeight {
		fieldErrorsLanguages["weight"] = append(fieldErrorsLanguages["weight"],
			errorMessages.GetLessOrEqualMaxNumber(locale, maxWeight))
	}

	if req.TelegramUsername == "" {
		fieldErrorsLanguages["telegramUsername"] = append(fieldErrorsLanguages["telegramUsername"],
			errorMessages.GetNotEmpty(locale))
	}

	if req.AgeFrom < minAge {
		fieldErrorsLanguages["ageFrom"] = append(fieldErrorsLanguages["ageFrom"],
			errorMessages.GetMoreOrEqualMinUint64Number(locale, minAge))
	}

	if req.AgeTo > maxAge {
		fieldErrorsLanguages["ageTo"] = append(fieldErrorsLanguages["ageTo"],
			errorMessages.GetLessOrEqualMaxUint64Number(locale, maxAge))
	}

	if req.Distance < 0 {
		fieldErrorsLanguages["distance"] = append(fieldErrorsLanguages["distance"],
			errorMessages.GetNonNegativeNumber(locale))
	}

	if int(req.Distance) > maxDistance {
		fieldErrorsLanguages["distance"] = append(fieldErrorsLanguages["distance"],
			errorMessages.GetLessOrEqualMaxNumber(locale, maxDistance))
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

	if !req.IsImages && len(files) <= 0 {
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
