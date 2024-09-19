package request

import (
	"time"
)

type ProfileTelegramDeleteRequestRepositoryDto struct {
	SessionID string    `json:"sessionId"`
	IsDeleted bool      `json:"isDeleted"`
	UpdatedAt time.Time `json:"updatedAt"`
}
