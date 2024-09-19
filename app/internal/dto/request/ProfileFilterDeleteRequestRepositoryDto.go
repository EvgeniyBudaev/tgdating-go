package request

import (
	"time"
)

type ProfileFilterDeleteRequestRepositoryDto struct {
	SessionID string    `json:"sessionId"`
	IsDeleted bool      `json:"isDeleted"`
	UpdatedAt time.Time `json:"updatedAt"`
}
