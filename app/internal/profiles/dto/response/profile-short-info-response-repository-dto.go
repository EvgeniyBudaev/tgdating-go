package response

import "github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/shared/enum"

type ProfileShortInfoResponseRepositoryDto struct {
	TelegramUserId string             `json:"telegramUserId"`
	IsBlocked      bool               `json:"isBlocked"`
	IsFrozen       bool               `json:"isFrozen"`
	LanguageCode   string             `json:"languageCode"`
	Measurement    enum.Measurement   `json:"measurement"`
	Filter         *FilterResponseDto `json:"filter"`
}
