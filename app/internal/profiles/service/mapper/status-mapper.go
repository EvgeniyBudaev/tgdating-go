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
		IsBlocked:      false,
		IsFrozen:       false,
		IsInvisible:    false,
		IsOnline:       false,
		IsPremium:      false,
		IsShowDistance: true,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
}
func (pm *StatusMapper) MapToResponse(p *entity.StatusEntity) *response.StatusResponseDto {
	return &response.StatusResponseDto{
		IsBlocked:      p.IsBlocked,
		IsFrozen:       p.IsFrozen,
		IsInvisible:    p.IsInvisible,
		IsOnline:       p.IsOnline,
		IsPremium:      p.IsPremium,
		IsShowDistance: p.IsShowDistance,
	}
}
