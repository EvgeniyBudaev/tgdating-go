package response

import "time"

type ProfileListItemResponseRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	FilterDistance *float64  `json:"filter_distance"`
	Distance       *float64  `json:"distance"`
	Page           uint64    `json:"page"`
	Size           uint64    `json:"page_size"`
	Url            string    `json:"url"`
	IsOnline       bool      `json:"isOnline"`
	LastOnline     time.Time `json:"lastOnline"`
}
