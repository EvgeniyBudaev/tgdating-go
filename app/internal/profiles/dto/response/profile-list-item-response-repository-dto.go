package response

import "time"

type ProfileListItemResponseRepositoryDto struct {
	Id             uint64    `json:"id"`
	SessionId      string    `json:"sessionId"`
	DisplayName    string    `json:"displayName"`
	Birthday       time.Time `json:"birthday"`
	Gender         string    `json:"gender"`
	Location       string    `json:"location"`
	Description    string    `json:"description"`
	Height         float64   `json:"height"`
	Weight         float64   `json:"weight"`
	IsFrozen       bool      `json:"isFrozen"`
	IsBlocked      bool      `json:"isBlocked"`
	IsPremium      bool      `json:"isPremium"`
	IsShowDistance bool      `json:"isShowDistance"`
	IsInvisible    bool      `json:"isInvisible"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	LastOnline     time.Time `json:"lastOnline"`
	Age            uint64    `json:"age"`
	Distance       *float64  `json:"distance"`
	LanguageCode   string    `json:"languageCode"`
}
