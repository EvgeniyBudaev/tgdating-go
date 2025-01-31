package request

import "time"

type StatusUpdateRequestRepositoryDto struct {
	TelegramUserId   string    `json:"telegramUserId"`
	IsBlocked        bool      `json:"isBlocked"`
	IsFrozen         bool      `json:"isFrozen"`
	IsInvisible      bool      `json:"isInvisible"`
	IsHiddenDistance bool      `json:"isHiddenDistance"`
	IsHiddenAge      bool      `json:"isHiddenAge"`
	IsLeftHand       bool      `json:"isLeftHand"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
