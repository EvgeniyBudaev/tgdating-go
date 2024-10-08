package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type TelegramMapper struct {
}

func (pm *TelegramMapper) MapToResponse(
	pe *entity.TelegramEntity) *response.TelegramResponseDto {
	return &response.TelegramResponseDto{
		SessionId:       pe.SessionId,
		UserId:          pe.UserId,
		Username:        pe.UserName,
		FirstName:       pe.FirstName,
		LastName:        pe.LastName,
		LanguageCode:    pe.LanguageCode,
		AllowsWriteToPm: pe.AllowsWriteToPm,
		QueryId:         pe.QueryId,
		ChatId:          pe.ChatId,
	}
}

func (pm *TelegramMapper) MapToAddRequest(
	pr *request.ProfileAddRequestDto) *request.TelegramAddRequestRepositoryDto {
	return &request.TelegramAddRequestRepositoryDto{
		SessionId:       pr.SessionId,
		UserId:          pr.TelegramUserId,
		UserName:        pr.TelegramUsername,
		FirstName:       pr.TelegramFirstName,
		LastName:        pr.TelegramLastName,
		LanguageCode:    pr.TelegramLanguageCode,
		AllowsWriteToPm: pr.TelegramAllowsWriteToPm,
		QueryId:         pr.TelegramQueryId,
		ChatId:          pr.TelegramChatId,
		IsDeleted:       false,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (pm *TelegramMapper) MapToUpdateRequest(
	pr *request.ProfileUpdateRequestDto) *request.TelegramUpdateRequestRepositoryDto {
	return &request.TelegramUpdateRequestRepositoryDto{
		SessionId:       pr.SessionId,
		UserId:          pr.TelegramUserId,
		UserName:        pr.TelegramUsername,
		FirstName:       pr.TelegramFirstName,
		LastName:        pr.TelegramLastName,
		LanguageCode:    pr.TelegramLanguageCode,
		AllowsWriteToPm: pr.TelegramAllowsWriteToPm,
		QueryId:         pr.TelegramQueryId,
		ChatId:          pr.TelegramChatId,
		UpdatedAt:       time.Now(),
	}
}

func (pm *TelegramMapper) MapToDeleteRequest(sessionId string) *request.TelegramDeleteRequestRepositoryDto {
	return &request.TelegramDeleteRequestRepositoryDto{
		SessionId: sessionId,
		IsDeleted: true,
		UpdatedAt: time.Now(),
	}
}
