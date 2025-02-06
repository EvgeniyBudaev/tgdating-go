package request

type BlockUpdateRequestRepositoryDto struct {
	TelegramUserId        string  `json:"telegramUserId"`
	BlockedTelegramUserId string  `json:"blockedTelegramUserId"`
	InitiatorId           *string `json:"initiatorId"`
}
