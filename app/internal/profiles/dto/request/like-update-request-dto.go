package request

type LikeUpdateRequestDto struct {
	Id             uint64 `json:"id"`
	TelegramUserId string `json:"telegramUserId"`
	IsLiked        bool   `json:"isLiked"`
}
