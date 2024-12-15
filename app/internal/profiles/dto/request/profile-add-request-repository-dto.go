package request

import "time"

type ProfileAddRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	DisplayName    string    `json:"displayName"`
	Age            uint64    `json:"age"`
	Gender         string    `json:"gender"`
	Location       string    `json:"location"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	LastOnline     time.Time `json:"lastOnline"`
}
