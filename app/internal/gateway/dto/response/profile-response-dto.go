package response

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/entity"
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
	IsOnline       bool                  `json:"isOnline"`
	CreatedAt      time.Time             `json:"createdAt"`
	UpdatedAt      time.Time             `json:"updatedAt"`
	LastOnline     time.Time             `json:"lastOnline"`
	Navigator      *NavigatorResponseDto `json:"navigator"`
	Filter         *FilterResponseDto    `json:"filter"`
	Telegram       *TelegramResponseDto  `json:"telegram"`
	Status         *StatusResponseDto    `json:"status"`
	Images         []*entity.ImageEntity `json:"images"`
}
