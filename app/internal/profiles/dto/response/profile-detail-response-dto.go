package response

import "time"

type ProfileDetailResponseDto struct {
	TelegramUserId string                      `json:"telegramUserId"`
	DisplayName    string                      `json:"displayName"`
	Age            uint64                      `json:"age"`
	Gender         string                      `json:"gender"`
	Description    string                      `json:"description"`
	LastOnline     time.Time                   `json:"lastOnline"`
	Navigator      *NavigatorDetailResponseDto `json:"navigator"`
	Status         *StatusResponseDto          `json:"status"`
	Settings       *SettingsResponseDto        `json:"settings"`
	Block          *BlockResponseDto           `json:"block"`
	Like           *LikeResponseDto            `json:"like"`
	Images         []*ImageResponseDto         `json:"images"`
}
