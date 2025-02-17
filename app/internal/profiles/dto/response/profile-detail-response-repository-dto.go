package response

import "time"

type ProfileDetailResponseRepositoryDto struct {
	TelegramUserId string                          `json:"telegramUserId"`
	DisplayName    string                          `json:"displayName"`
	Age            uint64                          `json:"age"`
	Gender         string                          `json:"gender"`
	Description    string                          `json:"description"`
	LastOnline     time.Time                       `json:"lastOnline"`
	Navigator      *NavigatorResponseRepositoryDto `json:"navigator"`
	Status         *StatusResponseRepositoryDto    `json:"status"`
	Settings       *SettingsResponseRepositoryDto  `json:"settings"`
	Block          *BlockResponseDto               `json:"block"`
	Like           *LikeResponseDto                `json:"like"`
}
