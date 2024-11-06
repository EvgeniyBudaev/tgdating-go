package request

import (
	"time"
)

type FilterUpdateRequestRepositoryDto struct {
	SessionId    string    `json:"sessionId"`
	SearchGender string    `json:"searchGender"`
	AgeFrom      byte      `json:"ageFrom"`
	AgeTo        byte      `json:"ageTo"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
