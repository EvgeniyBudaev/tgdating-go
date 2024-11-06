package request

import (
	"time"
)

type LikeAddRequestRepositoryDto struct {
	SessionId      string    `json:"sessionId"`
	LikedSessionId string    `json:"likedSessionId"`
	IsLiked        bool      `json:"isLiked"`
	IsDeleted      bool      `json:"isDeleted"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
