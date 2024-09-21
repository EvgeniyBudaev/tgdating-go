package entity

import "time"

type ProfileLikeEntity struct {
	Id             uint64    `json:"id"`
	SessionId      string    `json:"sessionId"`
	LikedSessionId string    `json:"likedSessionId"`
	IsLiked        bool      `json:"isLiked"`
	IsDeleted      bool      `json:"isDeleted"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
