package request

type StatusUpdateSettingsRequestRepositoryDto struct {
	TelegramUserId string `json:"telegramUserId"`
	IsHiddenAge    bool   `json:"isHiddenAge"`
}
