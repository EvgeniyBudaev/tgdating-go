package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type ProfileTelegramMapper struct {
}

func (pm *ProfileTelegramMapper) MapToResponse(r *entity.ProfileTelegramEntity) *response.ProfileTelegramResponseDto {
	return &response.ProfileTelegramResponseDto{
		SessionID:       r.SessionID,
		UserID:          r.UserID,
		Username:        r.UserName,
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		LanguageCode:    r.LanguageCode,
		AllowsWriteToPm: r.AllowsWriteToPm,
		QueryID:         r.QueryID,
		ChatID:          r.ChatID,
	}
}

func (pm *ProfileTelegramMapper) MapToAddRequest(r *request.ProfileAddRequestDto) *entity.ProfileTelegramEntity {
	return &entity.ProfileTelegramEntity{
		SessionID:       r.SessionID,
		UserID:          r.TelegramUserID,
		UserName:        r.TelegramUsername,
		FirstName:       r.TelegramFirstName,
		LastName:        r.TelegramLastName,
		LanguageCode:    r.TelegramLanguageCode,
		AllowsWriteToPm: r.TelegramAllowsWriteToPm,
		QueryID:         r.TelegramQueryID,
		ChatID:          r.TelegramChatID,
		IsDeleted:       false,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (pm *ProfileTelegramMapper) MapToUpdateRequest(r *request.ProfileUpdateRequestDto) *entity.ProfileTelegramEntity {
	return &entity.ProfileTelegramEntity{
		SessionID:       r.SessionID,
		UserID:          r.TelegramUserID,
		UserName:        r.TelegramUsername,
		FirstName:       r.TelegramFirstName,
		LastName:        r.TelegramLastName,
		LanguageCode:    r.TelegramLanguageCode,
		AllowsWriteToPm: r.TelegramAllowsWriteToPm,
		QueryID:         r.TelegramQueryID,
		ChatID:          r.TelegramChatID,
		UpdatedAt:       time.Now(),
	}
}
