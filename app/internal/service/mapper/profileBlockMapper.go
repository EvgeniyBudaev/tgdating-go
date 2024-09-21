package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"time"
)

type ProfileBlockMapper struct {
}

func (pm *ProfileBlockMapper) MapToAddRequest(
	pr *request.ProfileBlockRequestDto) *request.ProfileBlockAddRequestRepositoryDto {
	return &request.ProfileBlockAddRequestRepositoryDto{
		SessionId:            pr.SessionId,
		BlockedUserSessionId: pr.BlockedUserSessionId,
		IsBlocked:            true,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
}
