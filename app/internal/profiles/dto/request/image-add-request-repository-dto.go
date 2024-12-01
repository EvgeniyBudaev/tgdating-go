package request

import (
	"time"
)

type ImageAddRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	Name           string    `json:"name"`
	Url            string    `json:"url"`
	Size           int64     `json:"size"`
	IsBlocked      bool      `json:"isBlocked"`
	IsPrimary      bool      `json:"isPrimary"`
	IsPrivate      bool      `json:"isPrivate"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
