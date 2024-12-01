package request

type BlockAddRequestDto struct {
	TelegramUserId        string `json:"telegramUserId"`
	BlockedTelegramUserId string `json:"blockedTelegramUserId"`
}
