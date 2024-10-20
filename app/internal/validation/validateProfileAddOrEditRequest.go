package validation

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
)

func ValidateProfileAddOrEditRequestDto(req *request.ProfileAddRequestDto,
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

	if req.AgeFrom < 18 {
		fieldErrorsLanguages["ageFrom"] = append(fieldErrorsLanguages["ageFrom"],
			errorMessages.GetMoreOrEqualMinNumber(locale, 18))
	}

	if req.AgeTo > 100 {
		fieldErrorsLanguages["ageTo"] = append(fieldErrorsLanguages["ageTo"],
			errorMessages.GetLessOrEqualMaxNumber(locale, 100))
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

	if req.Page < 0 {
		fieldErrorsLanguages["page"] = append(fieldErrorsLanguages["page"],
			errorMessages.GetNonNegativeNumber(locale))
	}

	if req.Size < 0 {
		fieldErrorsLanguages["size"] = append(fieldErrorsLanguages["size"],
			errorMessages.GetNonNegativeNumber(locale))
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
