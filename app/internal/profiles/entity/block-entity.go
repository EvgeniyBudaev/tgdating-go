package entity

import "time"

type BlockEntity struct {
	Id                    uint64    `json:"id"`
	TelegramUserId        string    `json:"telegramUserId"`
	BlockedTelegramUserId string    `json:"blockedTelegramUserId"`
	IsBlocked             bool      `json:"isBlocked"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}
