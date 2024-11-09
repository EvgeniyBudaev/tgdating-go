package response

import "time"

type LikeResponseDto struct {
	Id             uint64    `json:"id"`
	SessionId      string    `json:"sessionId"`
	LikedSessionId string    `json:"likedSessionId"`
	IsLiked        bool      `json:"isLiked"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
