package response

import "time"

type LikeResponseDto struct {
	Id        uint64    `json:"id"`
	IsLiked   bool      `json:"isLiked"`
	UpdatedAt time.Time `json:"updatedAt"`
}
