package request

import "time"

type ProfileNavigatorUpdateRequestDto struct {
	SessionId string    `json:"sessionId"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	UpdatedAt time.Time `json:"updatedAt"`
}
