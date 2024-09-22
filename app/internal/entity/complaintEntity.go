package entity

import "time"

type ComplaintEntity struct {
	Id                uint64    `json:"id"`
	SessionId         string    `json:"sessionId"`
	CriminalSessionId string    `json:"criminalSessionId"`
	Reason            string    `json:"reason"`
	IsDeleted         bool      `json:"isDeleted"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
