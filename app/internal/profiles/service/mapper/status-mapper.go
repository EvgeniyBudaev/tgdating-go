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
		IsOnline:         false,
		IsPremium:        false,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	}
}
func (pm *StatusMapper) MapToResponse(p *entity.StatusEntity) *response.StatusResponseDto {
	return &response.StatusResponseDto{
		IsBlocked:        p.IsBlocked,
		IsFrozen:         p.IsFrozen,
		IsHiddenAge:      p.IsHiddenAge,
		IsHiddenDistance: p.IsHiddenDistance,
		IsInvisible:      p.IsInvisible,
		IsLeftHand:       p.IsLeftHand,
		IsOnline:         p.IsOnline,
		IsPremium:        p.IsPremium,
	}
}
