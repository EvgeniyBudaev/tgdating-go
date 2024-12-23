package request

import "time"

type StatusAddRequestRepositoryDto struct {
	TelegramUserId   string    `json:"telegramUserId"`
	IsBlocked        bool      `json:"isBlocked"`
	IsFrozen         bool      `json:"isFrozen"`
	IsHiddenAge      bool      `json:"isHiddenAge"`
	IsHiddenDistance bool      `json:"isHiddenDistance"`
	IsInvisible      bool      `json:"isInvisible"`
	IsLeftHand       bool      `json:"isLeftHand"`
	IsOnline         bool      `json:"isOnline"`
	IsPremium        bool      `json:"isPremium"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
