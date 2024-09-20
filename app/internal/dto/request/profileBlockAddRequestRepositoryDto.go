package request

import (
	"time"
)

type ProfileBlockAddRequestRepositoryDto struct {
	SessionID            string    `json:"sessionId"`
	BlockedUserSessionID string    `json:"blockedUserSessionId"`
	IsBlocked            bool      `json:"isBlocked"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
