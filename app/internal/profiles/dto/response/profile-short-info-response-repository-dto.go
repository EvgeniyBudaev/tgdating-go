package response

import "github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/shared/enum"

type ProfileShortInfoResponseRepositoryDto struct {
	TelegramUserId string           `json:"telegramUserId"`
	IsBlocked      bool             `json:"isBlocked"`
	IsFrozen       bool             `json:"isFrozen"`
	SearchGender   string           `json:"searchGender"`
	AgeFrom        uint64           `json:"ageFrom"`
	AgeTo          uint64           `json:"ageTo"`
	Distance       float64          `json:"distance"`
	Page           uint64           `json:"page"`
	Size           uint64           `json:"size"`
	LanguageCode   string           `json:"languageCode"`
	Measurement    enum.Measurement `json:"measurement"`
}
