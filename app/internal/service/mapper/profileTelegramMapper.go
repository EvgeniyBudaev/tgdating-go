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
	profileTelegramEntity *entity.ProfileTelegramEntity) *response.ProfileTelegramResponseDto {
	return &response.ProfileTelegramResponseDto{
		SessionID:       profileTelegramEntity.SessionID,
		UserID:          profileTelegramEntity.UserID,
		Username:        profileTelegramEntity.UserName,
		FirstName:       profileTelegramEntity.FirstName,
		LastName:        profileTelegramEntity.LastName,
		LanguageCode:    profileTelegramEntity.LanguageCode,
		AllowsWriteToPm: profileTelegramEntity.AllowsWriteToPm,
		QueryID:         profileTelegramEntity.QueryID,
		ChatID:          profileTelegramEntity.ChatID,
	}
}

func (pm *ProfileTelegramMapper) MapToAddRequest(
	profileAddRequestDto *request.ProfileAddRequestDto) *entity.ProfileTelegramEntity {
	return &entity.ProfileTelegramEntity{
		SessionID:       profileAddRequestDto.SessionID,
		UserID:          profileAddRequestDto.TelegramUserID,
		UserName:        profileAddRequestDto.TelegramUsername,
		FirstName:       profileAddRequestDto.TelegramFirstName,
		LastName:        profileAddRequestDto.TelegramLastName,
		LanguageCode:    profileAddRequestDto.TelegramLanguageCode,
		AllowsWriteToPm: profileAddRequestDto.TelegramAllowsWriteToPm,
		QueryID:         profileAddRequestDto.TelegramQueryID,
		ChatID:          profileAddRequestDto.TelegramChatID,
		IsDeleted:       false,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (pm *ProfileTelegramMapper) MapToUpdateRequest(
	profileUpdateRequestDto *request.ProfileUpdateRequestDto) *entity.ProfileTelegramEntity {
	return &entity.ProfileTelegramEntity{
		SessionID:       profileUpdateRequestDto.SessionID,
		UserID:          profileUpdateRequestDto.TelegramUserID,
		UserName:        profileUpdateRequestDto.TelegramUsername,
		FirstName:       profileUpdateRequestDto.TelegramFirstName,
		LastName:        profileUpdateRequestDto.TelegramLastName,
		LanguageCode:    profileUpdateRequestDto.TelegramLanguageCode,
		AllowsWriteToPm: profileUpdateRequestDto.TelegramAllowsWriteToPm,
		QueryID:         profileUpdateRequestDto.TelegramQueryID,
		ChatID:          profileUpdateRequestDto.TelegramChatID,
		UpdatedAt:       time.Now(),
	}
}
