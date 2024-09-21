package entity

import "time"

type ProfileBlockEntity struct {
	Id                   uint64    `json:"id"`
	SessionId            string    `json:"sessionId"`
	BlockedUserSessionId string    `json:"blockedUserSessionId"`
	IsBlocked            bool      `json:"isBlocked"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
