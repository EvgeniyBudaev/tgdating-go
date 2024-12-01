package response

type ProfileShortInfoResponseDto struct {
	TelegramUserId string `json:"telegramUserId"`
	ImageUrl       string `json:"imageUrl"`
	IsFrozen       bool   `json:"isFrozen"`
	IsBlocked      bool   `json:"isBlocked"`
}
