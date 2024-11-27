package request

type ProfileGetDetailRequestDto struct {
	SessionId string   `json:"sessionId"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}
