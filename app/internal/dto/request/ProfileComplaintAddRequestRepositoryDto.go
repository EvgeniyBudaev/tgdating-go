package request

import (
	"time"
)

type ProfileComplaintAddRequestRepositoryDto struct {
	SessionId         string    `json:"sessionId"`
	CriminalSessionId string    `json:"criminalSessionId"`
	Reason            string    `json:"reason"`
	IsDeleted         bool      `json:"isDeleted"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
