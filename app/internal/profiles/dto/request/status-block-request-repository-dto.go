package request

import "time"

type StatusBlockRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
