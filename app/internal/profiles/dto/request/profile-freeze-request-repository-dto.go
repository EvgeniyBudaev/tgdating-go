package request

import "time"

type ProfileFreezeRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	IsFrozen       bool      `json:"isFrozen"`
	UpdatedAt      time.Time `json:"updatedAt"`
	LastOnline     time.Time `json:"lastOnline"`
}
