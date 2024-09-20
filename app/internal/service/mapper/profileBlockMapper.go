package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"time"
)

type ProfileBlockMapper struct {
}

func (pm *ProfileBlockMapper) MapToAddRequest(
	profileBlockRequestDto *request.ProfileBlockRequestDto) *request.ProfileBlockAddRequestRepositoryDto {
	return &request.ProfileBlockAddRequestRepositoryDto{
		SessionID:            profileBlockRequestDto.SessionID,
		BlockedUserSessionID: profileBlockRequestDto.BlockedUserSessionID,
		IsBlocked:            true,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
}
