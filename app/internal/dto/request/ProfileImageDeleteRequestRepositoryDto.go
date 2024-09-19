package request

import (
	"time"
)

type ProfileImageDeleteRequestRepositoryDto struct {
	ID        uint64    `json:"id"`
	IsDeleted bool      `json:"isDeleted"`
	UpdatedAt time.Time `json:"updatedAt"`
}
