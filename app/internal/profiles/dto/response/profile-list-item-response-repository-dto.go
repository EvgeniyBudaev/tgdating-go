package response

import "time"

type ProfileListItemResponseRepositoryDto struct {
	Id             uint64    `json:"id"`
	TelegramUserId string    `json:"telegramUserId"`
	Birthday       time.Time `json:"birthday"`
	Gender         string    `json:"gender"`
	IsBlocked      bool      `json:"isBlocked"`
	IsFrozen       bool      `json:"isFrozen"`
	IsOnline       bool      `json:"isOnline"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	LastOnline     time.Time `json:"lastOnline"`
	Age            uint64    `json:"age"`
	Distance       *float64  `json:"distance"`
	Url            string    `json:"url"`
	TotalEntities  uint64    `json:"total_count"`
}
