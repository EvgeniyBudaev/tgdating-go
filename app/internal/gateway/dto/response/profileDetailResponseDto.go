package response

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/entity"
	"time"
)

type ProfileDetailResponseDto struct {
	SessionId      string                      `json:"sessionId"`
	DisplayName    string                      `json:"displayName"`
	Birthday       time.Time                   `json:"birthday"`
	Gender         string                      `json:"gender"`
	Location       string                      `json:"location"`
	Description    string                      `json:"description"`
	Height         float64                     `json:"height"`
	Weight         float64                     `json:"weight"`
	IsDeleted      bool                        `json:"isDeleted"`
	IsBlocked      bool                        `json:"isBlocked"`
	IsPremium      bool                        `json:"isPremium"`
	IsShowDistance bool                        `json:"isShowDistance"`
	IsInvisible    bool                        `json:"isInvisible"`
	IsOnline       bool                        `json:"isOnline"`
	CreatedAt      time.Time                   `json:"createdAt"`
	UpdatedAt      time.Time                   `json:"updatedAt"`
	LastOnline     time.Time                   `json:"lastOnline"`
	Navigator      *NavigatorDetailResponseDto `json:"navigator"`
	Telegram       *TelegramResponseDto        `json:"telegram"`
	Block          *BlockResponseDto           `json:"block"`
	Like           *LikeResponseDto            `json:"like"`
	Images         []*entity.ImageEntity       `json:"images"`
}
