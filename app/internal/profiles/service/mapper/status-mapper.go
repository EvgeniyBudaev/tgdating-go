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
		TelegramUserId: pr.TelegramUserId,
		IsFrozen:       false,
		IsBlocked:      false,
		IsPremium:      false,
		IsShowDistance: true,
		IsInvisible:    false,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
}
func (pm *StatusMapper) MapToResponse(pe *entity.StatusEntity) *response.StatusResponseDto {
	return &response.StatusResponseDto{
		IsFrozen:       pe.IsFrozen,
		IsBlocked:      pe.IsBlocked,
		IsPremium:      pe.IsPremium,
		IsShowDistance: pe.IsShowDistance,
		IsInvisible:    pe.IsInvisible,
	}
}
