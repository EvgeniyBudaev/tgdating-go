package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
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

func (pm *StatusMapper) MapToResponse(s *entity.StatusEntity, isPremium bool) *response.StatusResponseDto {
	return &response.StatusResponseDto{
		IsBlocked:        s.IsBlocked,
		IsFrozen:         s.IsFrozen,
		IsHiddenAge:      s.IsHiddenAge,
		IsHiddenDistance: s.IsHiddenDistance,
		IsInvisible:      s.IsInvisible,
		IsLeftHand:       s.IsLeftHand,
		IsPremium:        isPremium,
	}
}
