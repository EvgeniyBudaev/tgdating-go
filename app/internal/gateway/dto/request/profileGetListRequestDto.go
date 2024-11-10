package request

type ProfileGetListRequestDto struct {
	SessionId string   `json:"sessionId"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}
