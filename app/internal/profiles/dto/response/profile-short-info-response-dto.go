package response

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/shared/enum"
	"time"
)

type ProfileShortInfoResponseDto struct {
	TelegramUserId string             `json:"telegramUserId"`
	IsBlocked      bool               `json:"isBlocked"`
	IsFrozen       bool               `json:"isFrozen"`
	IsPremium      bool               `json:"isPremium"`
	AvailableUntil time.Time          `json:"availableUntil"`
	LanguageCode   string             `json:"languageCode"`
	Measurement    enum.Measurement   `json:"measurement"`
	Filter         *FilterResponseDto `json:"filter"`
}
