package entity

import "time"

type LikeEntity struct {
	Id             uint64    `json:"id"`
	SessionId      string    `json:"sessionId"`
	LikedSessionId string    `json:"likedSessionId"`
	IsLiked        bool      `json:"isLiked"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
