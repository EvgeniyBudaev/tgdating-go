package entity

import "time"

type ProfileBlockEntity struct {
	ID                   uint64    `json:"id"`
	SessionID            string    `json:"sessionId"`
	BlockedUserSessionID string    `json:"blockedUserSessionId"`
	IsBlocked            bool      `json:"isBlocked"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
