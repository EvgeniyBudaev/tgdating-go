package request

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enum"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
)

type ProfileAddRequestDto struct {
	DisplayName             string                 `json:"displayName"`
	Age                     uint64                 `json:"age"`
	Gender                  enum.Gender            `json:"gender"`
	SearchGender            enum.SearchGender      `json:"searchGender"`
	Description             string                 `json:"description"`
	TelegramUserId          string                 `json:"telegramUserId"`
	TelegramUsername        string                 `json:"telegramUsername"`
	TelegramFirstName       string                 `json:"telegramFirstName"`
	TelegramLastName        string                 `json:"telegramLastName"`
	TelegramLanguageCode    string                 `json:"telegramLanguageCode"`
	TelegramAllowsWriteToPm bool                   `json:"telegramAllowsWriteToPm"`
	TelegramQueryId         string                 `json:"telegramQueryId"`
	CountryCode             *string                `json:"countryCode"`
	CountryName             *string                `json:"countryName"`
	City                    *string                `json:"city"`
	Latitude                *float64               `json:"latitude"`
	Longitude               *float64               `json:"longitude"`
	AgeFrom                 uint64                 `json:"ageFrom"`
	AgeTo                   uint64                 `json:"ageTo"`
	Distance                float64                `json:"distance"`
	Page                    uint64                 `json:"page"`
	Size                    uint64                 `json:"size"`
	IsLiked                 bool                   `json:"isLiked"`
	IsOnline                bool                   `json:"isOnline"`
	IsLeftHand              bool                   `json:"isLeftHand"`
	Measurement             enum.Measurement       `json:"measurement"`
	Files                   []*entity.FileMetadata `json:"files"`
}
