package request

import "time"

type StatusRestoreRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
