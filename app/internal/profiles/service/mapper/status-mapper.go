package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"time"
)

type StatusMapper struct {
}

func (pm *StatusMapper) MapToAddRequest(pr *request.ProfileAddRequestDto) *request.StatusAddRequestRepositoryDto {
	return &request.StatusAddRequestRepositoryDto{
		TelegramUserId:   pr.TelegramUserId,
		IsBlocked:        false,
		IsFrozen:         false,
		IsHiddenAge:      false,
		IsHiddenDistance: false,
		IsInvisible:      false,
		IsLeftHand:       pr.IsLeftHand,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	}
}

func (pm *StatusMapper) MapToUpdateSettingsRequest(
	pr *request.ProfileUpdateSettingsRequestDto) *request.StatusUpdateSettingsRequestRepositoryDto {
	return &request.StatusUpdateSettingsRequestRepositoryDto{
		TelegramUserId: pr.TelegramUserId,
		IsHiddenAge:    pr.IsHiddenAge,
	}
}
