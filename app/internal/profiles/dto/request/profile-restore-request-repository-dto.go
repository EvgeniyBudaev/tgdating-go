package request

import "time"

type ProfileRestoreRequestRepositoryDto struct {
	SessionId  string    `json:"sessionId"`
	IsFrozen   bool      `json:"isFrozen"`
	UpdatedAt  time.Time `json:"updatedAt"`
	LastOnline time.Time `json:"lastOnline"`
}
