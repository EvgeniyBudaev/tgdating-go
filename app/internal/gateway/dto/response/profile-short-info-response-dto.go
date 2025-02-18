package response

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enum"
	"time"
)

type ProfileShortInfoResponseDto struct {
	TelegramUserId string             `json:"telegramUserId"`
	IsFrozen       bool               `json:"isFrozen"`
	IsBlocked      bool               `json:"isBlocked"`
	IsPremium      bool               `json:"isPremium"`
	AvailableUntil time.Time          `json:"availableUntil"`
	LanguageCode   string             `json:"languageCode"`
	Measurement    enum.Measurement   `json:"measurement"`
	Filter         *FilterResponseDto `json:"filter"`
}
