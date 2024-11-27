package request

import (
	"time"
)

type BlockAddRequestRepositoryDto struct {
	SessionId            string    `json:"sessionId"`
	BlockedUserSessionId string    `json:"blockedUserSessionId"`
	IsBlocked            bool      `json:"isBlocked"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
