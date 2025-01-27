package request

type UnblockRequestDto struct {
	TelegramUserId        string `json:"telegramUserId"`
	BlockedTelegramUserId string `json:"blockedTelegramUserId"`
}
