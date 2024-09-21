package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type ProfileTelegramMapper struct {
}

func (pm *ProfileTelegramMapper) MapToResponse(
	pe *entity.ProfileTelegramEntity) *response.ProfileTelegramResponseDto {
	return &response.ProfileTelegramResponseDto{
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

func (pm *ProfileTelegramMapper) MapToAddRequest(
	pr *request.ProfileAddRequestDto) *request.ProfileTelegramAddRequestRepositoryDto {
	return &request.ProfileTelegramAddRequestRepositoryDto{
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

func (pm *ProfileTelegramMapper) MapToUpdateRequest(
	pr *request.ProfileUpdateRequestDto) *request.ProfileTelegramUpdateRequestRepositoryDto {
	return &request.ProfileTelegramUpdateRequestRepositoryDto{
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

func (pm *ProfileTelegramMapper) MapToDeleteRequest(sessionId string) *request.ProfileTelegramDeleteRequestRepositoryDto {
	return &request.ProfileTelegramDeleteRequestRepositoryDto{
		SessionId: sessionId,
		IsDeleted: true,
		UpdatedAt: time.Now(),
	}
}
