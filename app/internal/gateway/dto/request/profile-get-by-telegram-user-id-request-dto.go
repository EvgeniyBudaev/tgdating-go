package request

type ProfileGetByTelegramUserIdRequestDto struct {
	CountryCode string   `json:"countryCode"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
}
