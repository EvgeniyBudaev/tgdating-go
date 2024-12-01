package entity

import "time"

type ImageEntity struct {
	Id             uint64    `json:"id"`
	TelegramUserId string    `json:"telegramUserId"`
	Name           string    `json:"name"`
	Url            string    `json:"url"`
	Size           int64     `json:"size"`
	IsDeleted      bool      `json:"isDeleted"`
	IsBlocked      bool      `json:"isBlocked"`
	IsPrimary      bool      `json:"isPrimary"`
	IsPrivate      bool      `json:"isPrivate"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
