package entity

import "time"

type LikeEntity struct {
	Id                  uint64    `json:"id"`
	TelegramUserId      string    `json:"telegramUserId"`
	LikedTelegramUserId string    `json:"likedTelegramUserId"`
	IsLiked             bool      `json:"isLiked"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}
