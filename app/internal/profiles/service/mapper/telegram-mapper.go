package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"time"
)

type TelegramMapper struct {
}

func (pm *TelegramMapper) MapToResponse(
	pe *entity.TelegramEntity) *response.TelegramResponseDto {
	return &response.TelegramResponseDto{
		UserId:          pe.UserId,
		Username:        pe.UserName,
		FirstName:       pe.FirstName,
		LastName:        pe.LastName,
		LanguageCode:    pe.LanguageCode,
		AllowsWriteToPm: pe.AllowsWriteToPm,
		QueryId:         pe.QueryId,
	}
}

func (pm *TelegramMapper) MapToAddRequest(
	pr *request.ProfileAddRequestDto) *request.TelegramAddRequestRepositoryDto {
	return &request.TelegramAddRequestRepositoryDto{
		UserId:          pr.TelegramUserId,
		UserName:        pr.TelegramUsername,
		FirstName:       pr.TelegramFirstName,
		LastName:        pr.TelegramLastName,
		LanguageCode:    pr.TelegramLanguageCode,
		AllowsWriteToPm: pr.TelegramAllowsWriteToPm,
		QueryId:         pr.TelegramQueryId,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (pm *TelegramMapper) MapToUpdateRequest(
	pr *request.ProfileUpdateRequestDto) *request.TelegramUpdateRequestRepositoryDto {
	return &request.TelegramUpdateRequestRepositoryDto{
		UserId:          pr.TelegramUserId,
		UserName:        pr.TelegramUsername,
		FirstName:       pr.TelegramFirstName,
		LastName:        pr.TelegramLastName,
		LanguageCode:    pr.TelegramLanguageCode,
		AllowsWriteToPm: pr.TelegramAllowsWriteToPm,
		QueryId:         pr.TelegramQueryId,
		UpdatedAt:       time.Now(),
	}
}
