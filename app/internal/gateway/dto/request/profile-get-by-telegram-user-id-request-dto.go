package request

type ProfileGetByTelegramUserIdRequestDto struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}
