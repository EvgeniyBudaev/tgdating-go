package entity

import "time"

type StatusEntity struct {
	Id             uint64    `json:"id"`
	TelegramUserId string    `json:"telegramUserId"`
	IsFrozen       bool      `json:"isFrozen"`
	IsBlocked      bool      `json:"isBlocked"`
	IsPremium      bool      `json:"isPremium"`
	IsShowDistance bool      `json:"isShowDistance"`
	IsInvisible    bool      `json:"isInvisible"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
