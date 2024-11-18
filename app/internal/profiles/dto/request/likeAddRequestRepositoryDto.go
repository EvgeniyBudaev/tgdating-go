package request

import (
	"time"
)

type LikeAddRequestRepositoryDto struct {
	SessionId      string    `json:"sessionId"`
	LikedSessionId string    `json:"likedSessionId"`
	IsLiked        bool      `json:"isLiked"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
