package entity

import "time"

type ProfileNavigatorEntity struct {
	ID        uint64       `json:"id"`
	SessionID string       `json:"sessionId"`
	Location  *PointEntity `json:"location"`
	IsDeleted bool         `json:"isDeleted"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}
