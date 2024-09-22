package request

import (
	"time"
)

type TelegramUpdateRequestRepositoryDto struct {
	SessionId       string    `json:"sessionId"`
	UserId          uint64    `json:"userId"`
	UserName        string    `json:"username"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	LanguageCode    string    `json:"languageCode"`
	AllowsWriteToPm bool      `json:"allowsWriteToPm"`
	QueryId         string    `json:"queryId"`
	ChatId          uint64    `json:"chatId"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
