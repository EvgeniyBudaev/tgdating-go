package request

import (
	"time"
)

type ComplaintAddRequestRepositoryDto struct {
	TelegramUserId         string    `json:"telegramUserId"`
	CriminalTelegramUserId string    `json:"criminalTelegramUserId"`
	Reason                 string    `json:"reason"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
}
