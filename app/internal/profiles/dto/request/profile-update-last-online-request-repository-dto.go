package request

import (
	"time"
)

type ProfileUpdateLastOnlineRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	LastOnline     time.Time `json:"lastOnline"`
}
