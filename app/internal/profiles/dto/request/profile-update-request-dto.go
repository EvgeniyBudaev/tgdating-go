package request

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/shared/enum"
)

type ProfileUpdateRequestDto struct {
	TelegramUserId          string                 `json:"telegramUserId"`
	DisplayName             string                 `json:"displayName"`
	Age                     uint64                 `json:"age"`
	Gender                  enum.Gender            `json:"gender"`
	SearchGender            enum.SearchGender      `json:"searchGender"`
	Description             string                 `json:"description"`
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
	IsImages                bool                   `json:"IsImages"`
	Measurement             enum.Measurement       `json:"measurement"`
	Files                   []*entity.FileMetadata `json:"files"`
}
