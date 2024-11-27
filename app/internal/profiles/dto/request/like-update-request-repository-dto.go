package request

import "time"

type LikeUpdateRequestRepositoryDto struct {
	Id        uint64    `json:"id"`
	IsLiked   bool      `json:"isLiked"`
	UpdatedAt time.Time `json:"updatedAt"`
}
