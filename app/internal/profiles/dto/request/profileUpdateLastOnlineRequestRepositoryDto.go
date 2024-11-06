package request

import (
	"time"
)

type ProfileUpdateLastOnlineRequestRepositoryDto struct {
	SessionId  string    `json:"sessionId"`
	LastOnline time.Time `json:"lastOnline"`
}
