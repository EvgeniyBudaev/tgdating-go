package entity

import (
	"time"
)

type ProfileEntity struct {
	ID             uint64    `json:"id"`
	SessionID      string    `json:"sessionId"`
	DisplayName    string    `json:"displayName"`
	Birthday       time.Time `json:"birthday"`
	Gender         string    `json:"gender"`
	Location       string    `json:"location"`
	Height         float64   `json:"height"`
	Weight         float64   `json:"weight"`
	Description    string    `json:"description"`
	IsDeleted      bool      `json:"isDeleted"`
	IsBlocked      bool      `json:"isBlocked"`
	IsPremium      bool      `json:"isPremium"`
	IsShowDistance bool      `json:"isShowDistance"`
	IsInvisible    bool      `json:"isInvisible"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	LastOnline     time.Time `json:"lastOnline"`
}
