package request

import (
	"time"
)

type TelegramDeleteRequestRepositoryDto struct {
	SessionId string    `json:"sessionId"`
	IsDeleted bool      `json:"isDeleted"`
	UpdatedAt time.Time `json:"updatedAt"`
}
