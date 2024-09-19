package request

import (
	"time"
)

type ProfileTelegramAddRequestRepositoryDto struct {
	SessionID       string    `json:"sessionId"`
	UserID          uint64    `json:"userId"`
	UserName        string    `json:"username"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	LanguageCode    string    `json:"languageCode"`
	AllowsWriteToPm bool      `json:"allowsWriteToPm"`
	QueryID         string    `json:"queryId"`
	ChatID          uint64    `json:"chatId"`
	IsDeleted       bool      `json:"isDeleted"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
