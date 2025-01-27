package response

import "time"

type ProfileShortInfoResponseDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	IsFrozen       bool      `json:"isFrozen"`
	IsBlocked      bool      `json:"isBlocked"`
	IsPremium      bool      `json:"isPremium"`
	AvailableUntil time.Time `json:"availableUntil"`
	SearchGender   string    `json:"searchGender"`
	AgeFrom        uint64    `json:"ageFrom"`
	AgeTo          uint64    `json:"ageTo"`
	Distance       float64   `json:"distance"`
	Page           uint64    `json:"page"`
	Size           uint64    `json:"size"`
	LanguageCode   string    `json:"languageCode"`
}
