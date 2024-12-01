package response

import "time"

type LikeResponseDto struct {
	Id                  uint64    `json:"id"`
	TelegramUserId      string    `json:"telegramUserId"`
	LikedTelegramUserId string    `json:"likedTelegramUserId"`
	IsLiked             bool      `json:"isLiked"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}
