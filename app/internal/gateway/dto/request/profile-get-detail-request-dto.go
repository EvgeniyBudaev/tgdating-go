package request

type ProfileGetDetailRequestDto struct {
	TelegramUserId string   `json:"telegramUserId"`
	Latitude       *float64 `json:"latitude"`
	Longitude      *float64 `json:"longitude"`
}
