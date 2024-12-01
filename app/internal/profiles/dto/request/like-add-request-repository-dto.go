package request

import (
	"time"
)

type LikeAddRequestRepositoryDto struct {
	TelegramUserId      string    `json:"telegramUserId"`
	LikedTelegramUserId string    `json:"likedTelegramUserId"`
	IsLiked             bool      `json:"isLiked"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}
