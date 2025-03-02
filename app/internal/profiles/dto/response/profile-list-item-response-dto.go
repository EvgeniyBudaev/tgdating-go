package response

import (
	"time"
)

type ProfileListItemResponseDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	Distance       *float64  `json:"distance"`
	Url            string    `json:"url"`
	IsLiked        bool      `json:"isLiked"`
	LastOnline     time.Time `json:"lastOnline"`
}
