package request

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/shared/enum"
)

type ProfileAddRequestDto struct {
	TelegramUserId          string                 `json:"telegramUserId"`
	DisplayName             string                 `json:"displayName"`
	Age                     uint64                 `json:"age"`
	Gender                  enum.Gender            `json:"gender"`
	SearchGender            enum.SearchGender      `json:"searchGender"`
	Location                string                 `json:"location"`
	Description             string                 `json:"description"`
	TelegramUsername        string                 `json:"telegramUsername"`
	TelegramFirstName       string                 `json:"telegramFirstName"`
	TelegramLastName        string                 `json:"telegramLastName"`
	TelegramLanguageCode    string                 `json:"telegramLanguageCode"`
	TelegramAllowsWriteToPm bool                   `json:"telegramAllowsWriteToPm"`
	TelegramQueryId         string                 `json:"telegramQueryId"`
	Latitude                *float64               `json:"latitude"`
	Longitude               *float64               `json:"longitude"`
	AgeFrom                 uint64                 `json:"ageFrom"`
	AgeTo                   uint64                 `json:"ageTo"`
	Distance                float64                `json:"distance"`
	Page                    uint64                 `json:"page"`
	Size                    uint64                 `json:"size"`
	IsLeftHand              bool                   `json:"isLeftHand"`
	Files                   []*entity.FileMetadata `json:"files"`
}
