package request

type ProfileUpdateSettingsRequestDto struct {
	TelegramUserId string `json:"telegramUserId"`
	IsHiddenAge    bool   `json:"isHiddenAge"`
}
