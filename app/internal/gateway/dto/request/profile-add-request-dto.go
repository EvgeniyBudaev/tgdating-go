package request

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enum"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"time"
)

type ProfileAddRequestDto struct {
	DisplayName             string                 `json:"displayName"`
	Birthday                time.Time              `json:"birthday"`
	Gender                  enum.Gender            `json:"gender"`
	SearchGender            enum.SearchGender      `json:"searchGender"`
	Location                string                 `json:"location"`
	Description             string                 `json:"description"`
	Height                  float64                `json:"height"`
	Weight                  float64                `json:"weight"`
	LookingFor              enum.LookingFor        `json:"lookingFor"`
	TelegramUserId          string                 `json:"telegramUserId"`
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
	Files                   []*entity.FileMetadata `json:"files"`
}
