package entity

import "time"

type ProfileTelegramEntity struct {
	ID              uint64    `json:"id"`
	SessionID       string    `json:"sessionId"`
	UserID          uint64    `json:"userId"`
	UserName        string    `json:"username"`
	Firstname       string    `json:"firstName"`
	Lastname        string    `json:"lastName"`
	LanguageCode    string    `json:"languageCode"`
	AllowsWriteToPm bool      `json:"allowsWriteToPm"`
	QueryID         string    `json:"queryId"`
	ChatID          uint64    `json:"chatId"`
	IsDeleted       bool      `json:"isDeleted"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
