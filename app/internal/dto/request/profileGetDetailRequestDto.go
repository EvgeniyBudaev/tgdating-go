package request

type ProfileGetDetailRequestDto struct {
	ViewedSessionId string  `json:"viewedSessionId"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
}
