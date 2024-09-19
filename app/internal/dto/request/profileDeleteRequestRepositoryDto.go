package request

import "time"

type ProfileDeleteRequestRepositoryDto struct {
	SessionID  string    `json:"sessionId"`
	IsDeleted  bool      `json:"isDeleted"`
	UpdatedAt  time.Time `json:"updatedAt"`
	LastOnline time.Time `json:"lastOnline"`
}
