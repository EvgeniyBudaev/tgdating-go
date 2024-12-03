package request

import "time"

type ProfileAddRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	DisplayName    string    `json:"displayName"`
	Birthday       time.Time `json:"birthday"`
	Gender         string    `json:"gender"`
	Location       string    `json:"location"`
	Description    string    `json:"description"`
	Height         float64   `json:"height"`
	Weight         float64   `json:"weight"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	LastOnline     time.Time `json:"lastOnline"`
}
