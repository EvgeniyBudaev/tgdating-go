package request

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enums"
	"time"
)

type ProfileUpdateRequestDto struct {
	SessionId               string                 `json:"sessionId"`
	DisplayName             string                 `json:"displayName"`
	Birthday                time.Time              `json:"birthday"`
	Gender                  enums.Gender           `json:"gender"`
	SearchGender            enums.SearchGender     `json:"searchGender"`
	Location                string                 `json:"location"`
	Description             string                 `json:"description"`
	Height                  float64                `json:"height"`
	Weight                  float64                `json:"weight"`
	LookingFor              enums.LookingFor       `json:"lookingFor"`
	TelegramUserId          uint64                 `json:"telegramUserId"`
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
	IsImages                bool                   `json:"IsImages"`
	Files                   []*entity.FileMetadata `json:"files"`
}
