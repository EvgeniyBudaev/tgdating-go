package request

import "time"

type ProfileNavigatorUpdateRequestDto struct {
	SessionID string    `json:"sessionId"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	UpdatedAt time.Time `json:"updatedAt"`
}
