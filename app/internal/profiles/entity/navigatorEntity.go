package entity

import "time"

type NavigatorEntity struct {
	Id        uint64       `json:"id"`
	SessionId string       `json:"sessionId"`
	Location  *PointEntity `json:"location"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}
