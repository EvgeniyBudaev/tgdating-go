package request

import "time"

type NavigatorUpdateRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
