package entity

import "time"

type ComplaintEntity struct {
	Id                     uint64    `json:"id"`
	TelegramUserId         string    `json:"telegramUserId"`
	CriminalTelegramUserId string    `json:"criminalTelegramUserId"`
	Reason                 string    `json:"reason"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
}
