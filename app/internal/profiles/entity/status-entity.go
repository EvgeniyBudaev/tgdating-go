package entity

import "time"

type StatusEntity struct {
	Id               uint64    `json:"id"`
	TelegramUserId   string    `json:"telegramUserId"`
	IsBlocked        bool      `json:"isBlocked"`
	IsFrozen         bool      `json:"isFrozen"`
	IsHiddenAge      bool      `json:"isHiddenAge"`
	IsHiddenDistance bool      `json:"isHiddenDistance"`
	IsInvisible      bool      `json:"isInvisible"`
	IsLeftHand       bool      `json:"isLeftHand"`
	IsOnline         bool      `json:"isOnline"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
