package entity

import "time"

type FilterEntity struct {
	Id           uint64    `json:"id"`
	SessionId    string    `json:"sessionId"`
	SearchGender string    `json:"searchGender"`
	LookingFor   string    `json:"lookingFor"`
	AgeFrom      byte      `json:"ageFrom"`
	AgeTo        byte      `json:"ageTo"`
	Distance     float64   `json:"distance"`
	Page         uint64    `json:"page"`
	Size         uint64    `json:"size"`
	IsDeleted    bool      `json:"isDeleted"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
