package response

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"time"
)

type ProfileDetailResponseRepositoryDto struct {
	TelegramUserId string                      `json:"telegramUserId"`
	DisplayName    string                      `json:"displayName"`
	Birthday       time.Time                   `json:"birthday"`
	Location       string                      `json:"location"`
	Description    string                      `json:"description"`
	Height         float64                     `json:"height"`
	Weight         float64                     `json:"weight"`
	Navigator      *NavigatorDetailResponseDto `json:"navigator"`
	Telegram       *TelegramResponseDto        `json:"telegram"`
	IsBlocked      bool                        `json:"isBlocked"`
	IsFrozen       bool                        `json:"isFrozen"`
	IsOnline       bool                        `json:"isOnline"`
	Block          *BlockResponseDto           `json:"block"`
	Like           *LikeResponseDto            `json:"like"`
	Images         []*entity.ImageEntity       `json:"images"`
}
