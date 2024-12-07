package response

import (
	"time"
)

type ProfileDetailResponseDto struct {
	TelegramUserId string                      `json:"telegramUserId"`
	DisplayName    string                      `json:"displayName"`
	Birthday       time.Time                   `json:"birthday"`
	Location       string                      `json:"location"`
	Description    string                      `json:"description"`
	Height         float64                     `json:"height"`
	Weight         float64                     `json:"weight"`
	Navigator      *NavigatorDetailResponseDto `json:"navigator"`
	Status         *StatusResponseDto          `json:"status"`
	Block          *BlockResponseDto           `json:"block"`
	Like           *LikeResponseDto            `json:"like"`
	Images         []*ImageResponseDto         `json:"images"`
}
