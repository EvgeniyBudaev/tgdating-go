package request

import "time"

type StatusUpdateRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	IsBlocked      bool      `json:"isBlocked"`
	IsFrozen       bool      `json:"isFrozen"`
	IsInvisible    bool      `json:"isInvisible"`
	IsOnline       bool      `json:"isOnline"`
	IsPremium      bool      `json:"isPremium"`
	IsShowDistance bool      `json:"isShowDistance"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
