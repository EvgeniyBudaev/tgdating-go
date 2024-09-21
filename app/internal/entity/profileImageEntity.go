package entity

import "time"

type ProfileImageEntity struct {
	Id        uint64    `json:"id"`
	SessionId string    `json:"sessionId"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Size      int64     `json:"size"`
	IsDeleted bool      `json:"isDeleted"`
	IsBlocked bool      `json:"isBlocked"`
	IsPrimary bool      `json:"isPrimary"`
	IsPrivate bool      `json:"isPrivate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
