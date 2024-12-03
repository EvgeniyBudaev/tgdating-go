package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"time"
)

type ProfileUpdateLastOnlineMapper struct{}

func (pm *ProfileUpdateLastOnlineMapper) MapToAddRequest(
	telegramUserId string) *request.ProfileUpdateLastOnlineRequestRepositoryDto {
	return &request.ProfileUpdateLastOnlineRequestRepositoryDto{
		TelegramUserId: telegramUserId,
		LastOnline:     time.Now().UTC(),
	}
}
