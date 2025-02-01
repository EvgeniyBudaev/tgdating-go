package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"time"
)

type SettingsMapper struct {
}

func (pm *SettingsMapper) MapToAddRequest(pr *request.ProfileAddRequestDto) *request.SettingsAddRequestRepositoryDto {
	return &request.SettingsAddRequestRepositoryDto{
		TelegramUserId: pr.TelegramUserId,
		Measurement:    pr.Measurement,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
}

func (pm *SettingsMapper) MapToUpdateRequest(
	pr *request.ProfileUpdateRequestDto) *request.SettingsUpdateRequestRepositoryDto {
	return &request.SettingsUpdateRequestRepositoryDto{
		TelegramUserId: pr.TelegramUserId,
		Measurement:    pr.Measurement,
		UpdatedAt:      time.Now().UTC(),
	}
}
