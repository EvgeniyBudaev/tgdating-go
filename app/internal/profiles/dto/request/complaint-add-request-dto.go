package request

type ComplaintAddRequestDto struct {
	TelegramUserId         string `json:"telegramUserId"`
	CriminalTelegramUserId string `json:"criminalTelegramUserId"`
	Reason                 string `json:"reason"`
}
