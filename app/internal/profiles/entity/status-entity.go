package entity

import "time"

type StatusEntity struct {
	Id             uint64    `json:"id"`
	TelegramUserId string    `json:"telegramUserId"`
	IsBlocked      bool      `json:"isBlocked"`
	IsFrozen       bool      `json:"isFrozen"`
	IsInvisible    bool      `json:"isInvisible"`
	IsOnline       bool      `json:"isOnline"`
	IsPremium      bool      `json:"isPremium"`
	IsShowDistance bool      `json:"isShowDistance"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
