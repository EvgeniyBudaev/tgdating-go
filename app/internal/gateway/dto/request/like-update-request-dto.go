package request

type LikeUpdateRequestDto struct {
	Id                  uint64 `json:"id"`
	TelegramUserId      string `json:"telegramUserId"`
	LikedTelegramUserId string `json:"likedTelegramUserId"`
	IsLiked             bool   `json:"isLiked"`
}
