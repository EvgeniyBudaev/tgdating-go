package request

type ProfileGetShortInfoRequestDto struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}
