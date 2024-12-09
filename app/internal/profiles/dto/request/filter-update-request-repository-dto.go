package request

import (
	"time"
)

type FilterUpdateRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	SearchGender   string    `json:"searchGender"`
	AgeFrom        uint64    `json:"ageFrom"`
	AgeTo          uint64    `json:"ageTo"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
