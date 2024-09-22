package request

import "time"

type NavigatorUpdateRequestDto struct {
	SessionId string    `json:"sessionId"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	UpdatedAt time.Time `json:"updatedAt"`
}
