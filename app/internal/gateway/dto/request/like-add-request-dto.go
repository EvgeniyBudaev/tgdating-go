package request

type LikeAddRequestDto struct {
	TelegramUserId      string `json:"telegramUserId"`
	LikedTelegramUserId string `json:"likedTelegramUserId"`
}
