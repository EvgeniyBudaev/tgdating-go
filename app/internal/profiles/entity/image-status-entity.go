package entity

import "time"

type ImageStatusEntity struct {
	Id        uint64    `json:"id"`
	ImageId   uint64    `json:"imageId"`
	IsBlocked bool      `json:"isBlocked"`
	IsPrimary bool      `json:"isPrimary"`
	IsPrivate bool      `json:"isPrivate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
