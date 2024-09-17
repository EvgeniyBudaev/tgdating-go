package request

import "time"

type ProfileAddRequestDto struct {
	SessionID               string    `json:"sessionId"`
	DisplayName             string    `json:"displayName"`
	Birthday                time.Time `json:"birthday"`
	Gender                  string    `json:"gender"`
	SearchGender            string    `json:"searchGender"`
	Location                string    `json:"location"`
	Description             string    `json:"description"`
	Height                  float64   `json:"height"`
	Weight                  float64   `json:"weight"`
	LookingFor              string    `json:"lookingFor"`
	TelegramUserID          uint64    `json:"telegramUserId"`
	TelegramUsername        string    `json:"telegramUsername"`
	TelegramFirstname       string    `json:"telegramFirstName"`
	TelegramLastname        string    `json:"telegramLastName"`
	TelegramLanguageCode    string    `json:"telegramLanguageCode"`
	TelegramAllowsWriteToPm bool      `json:"telegramAllowsWriteToPm"`
	TelegramQueryID         string    `json:"telegramQueryId"`
	TelegramChatID          uint64    `json:"telegramChatId"`
	Latitude                float64   `json:"latitude"`
	Longitude               float64   `json:"longitude"`
	AgeFrom                 byte      `json:"ageFrom"`
	AgeTo                   byte      `json:"ageTo"`
	Distance                float64   `json:"distance"`
	Page                    uint64    `json:"page"`
	Size                    uint64    `json:"size"`
	Image                   []byte    `json:"image"`
}
