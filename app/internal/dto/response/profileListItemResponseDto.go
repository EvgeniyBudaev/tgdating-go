package response

import "time"

type ProfileListItemResponseDto struct {
	SessionId  string    `json:"sessionId"`
	Distance   float64   `json:"distance"`
	Url        string    `json:"url"`
	IsOnline   bool      `json:"isOnline"`
	LastOnline time.Time `json:"lastOnline"`
}
