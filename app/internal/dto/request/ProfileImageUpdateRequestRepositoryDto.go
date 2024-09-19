package request

import (
	"time"
)

type ProfileImageUpdateRequestRepositoryDto struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Size      int64     `json:"size"`
	IsDeleted bool      `json:"isDeleted"`
	IsBlocked bool      `json:"isBlocked"`
	IsPrimary bool      `json:"isPrimary"`
	IsPrivate bool      `json:"isPrivate"`
	UpdatedAt time.Time `json:"updatedAt"`
}
