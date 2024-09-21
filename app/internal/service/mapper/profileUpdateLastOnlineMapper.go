package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"time"
)

type ProfileUpdateLastOnlineMapper struct{}

func (pm *ProfileUpdateLastOnlineMapper) MapToAddRequest(sessionId string) *request.ProfileUpdateLastOnlineRequestRepositoryDto {
	return &request.ProfileUpdateLastOnlineRequestRepositoryDto{
		SessionId:  sessionId,
		LastOnline: time.Now().UTC(),
	}
}
