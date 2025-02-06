package request

import (
	"time"
)

type BlockAddRequestRepositoryDto struct {
	TelegramUserId        string    `json:"telegramUserId"`
	BlockedTelegramUserId string    `json:"blockedTelegramUserId"`
	InitiatorId           *string   `json:"initiatorId"`
	IsBlocked             bool      `json:"isBlocked"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}
