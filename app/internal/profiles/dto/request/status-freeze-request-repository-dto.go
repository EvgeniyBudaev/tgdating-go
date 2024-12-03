package request

import "time"

type StatusFreezeRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
