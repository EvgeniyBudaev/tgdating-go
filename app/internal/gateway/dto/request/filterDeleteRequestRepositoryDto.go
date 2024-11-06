package request

import (
	"time"
)

type FilterDeleteRequestRepositoryDto struct {
	SessionId string    `json:"sessionId"`
	IsDeleted bool      `json:"isDeleted"`
	UpdatedAt time.Time `json:"updatedAt"`
}
