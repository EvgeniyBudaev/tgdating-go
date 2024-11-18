package request

import "time"

type ProfileRestoreRequestRepositoryDto struct {
	SessionId  string    `json:"sessionId"`
	IsDeleted  bool      `json:"isDeleted"`
	UpdatedAt  time.Time `json:"updatedAt"`
	LastOnline time.Time `json:"lastOnline"`
}
