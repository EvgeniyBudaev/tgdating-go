package request

type ProfileGetDetailRequestDto struct {
	TelegramUserId string   `json:"telegramUserId"`
	CountryCode    *string  `json:"countryCode"`
	CountryName    *string  `json:"countryName"`
	City           *string  `json:"city"`
	Latitude       *float64 `json:"latitude"`
	Longitude      *float64 `json:"longitude"`
}
