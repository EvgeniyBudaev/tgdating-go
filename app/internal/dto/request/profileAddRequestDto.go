package request

import "time"

type ProfileAddRequestDto struct {
	SessionId               string    `json:"sessionId"`
	DisplayName             string    `json:"displayName"`
	Birthday                time.Time `json:"birthday"`
	Gender                  string    `json:"gender"`
	SearchGender            string    `json:"searchGender"`
	Location                string    `json:"location"`
	Description             string    `json:"description"`
	Height                  float64   `json:"height"`
	Weight                  float64   `json:"weight"`
	LookingFor              string    `json:"lookingFor"`
	TelegramUserId          uint64    `json:"telegramUserId"`
	TelegramUsername        string    `json:"telegramUsername"`
	TelegramFirstName       string    `json:"telegramFirstName"`
	TelegramLastName        string    `json:"telegramLastName"`
	TelegramLanguageCode    string    `json:"telegramLanguageCode"`
	TelegramAllowsWriteToPm bool      `json:"telegramAllowsWriteToPm"`
	TelegramQueryId         string    `json:"telegramQueryId"`
	TelegramChatId          uint64    `json:"telegramChatId"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	AgeFrom                 byte      `json:"ageFrom"`
	AgeTo                   byte      `json:"ageTo"`
	Distance                float64   `json:"distance"`
	Page                    uint64    `json:"page"`
	Size                    uint64    `json:"size"`
	Image                   []byte    `json:"image"`
}
