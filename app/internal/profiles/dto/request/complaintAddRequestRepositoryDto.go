package request

import (
	"time"
)

type ComplaintAddRequestRepositoryDto struct {
	SessionId         string    `json:"sessionId"`
	CriminalSessionId string    `json:"criminalSessionId"`
	Reason            string    `json:"reason"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
