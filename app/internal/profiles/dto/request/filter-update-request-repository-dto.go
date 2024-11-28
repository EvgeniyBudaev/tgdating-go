package request

import (
	"time"
)

type FilterUpdateRequestRepositoryDto struct {
	SessionId    string    `json:"sessionId"`
	SearchGender string    `json:"searchGender"`
	LookingFor   string    `json:"lookingFor"`
	AgeFrom      uint64    `json:"ageFrom"`
	AgeTo        uint64    `json:"ageTo"`
	UpdatedAt    time.Time `json:"updatedAt"`
}