package request

import (
	"time"
)

type ImageDeleteRequestRepositoryDto struct {
	Id        uint64    `json:"id"`
	IsDeleted bool      `json:"isDeleted"`
	UpdatedAt time.Time `json:"updatedAt"`
}
