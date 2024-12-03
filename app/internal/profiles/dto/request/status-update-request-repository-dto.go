package request

import "time"

type StatusUpdateRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	IsFrozen       bool      `json:"isFrozen"`
	IsBlocked      bool      `json:"isBlocked"`
	IsPremium      bool      `json:"isPremium"`
	IsShowDistance bool      `json:"isShowDistance"`
	IsInvisible    bool      `json:"isInvisible"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
