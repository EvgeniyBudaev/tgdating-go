package response

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enum"
	"time"
)

type ProfileListItemResponseDto struct {
	TelegramUserId string           `json:"telegramUserId"`
	Distance       *float64         `json:"distance"`
	Url            string           `json:"url"`
	IsOnline       bool             `json:"isOnline"`
	IsLiked        bool             `json:"isLiked"`
	LastOnline     time.Time        `json:"lastOnline"`
	Measurement    enum.Measurement `json:"measurement"`
}
