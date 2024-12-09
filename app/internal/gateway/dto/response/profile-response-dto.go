package response

import (
	"time"
)

type ProfileResponseDto struct {
	TelegramUserId string                `json:"telegramUserId"`
	DisplayName    string                `json:"displayName"`
	Birthday       time.Time             `json:"birthday"`
	Gender         string                `json:"gender"`
	Location       string                `json:"location"`
	Description    string                `json:"description"`
	Height         float64               `json:"height"`
	Weight         float64               `json:"weight"`
	Navigator      *NavigatorResponseDto `json:"navigator"`
	Filter         *FilterResponseDto    `json:"filter"`
	Status         *StatusResponseDto    `json:"status"`
	Images         []*ImageResponseDto   `json:"images"`
}
