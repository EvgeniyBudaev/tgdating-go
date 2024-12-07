package response

type CheckExistsDto struct {
	TelegramUserId string `json:"telegramUserId"`
	IsFrozen       bool   `json:"isFrozen"`
}
