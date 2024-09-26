package request

type FilterUpdateRequestDto struct {
	SessionId    string  `json:"sessionId"`
	SearchGender string  `json:"searchGender"`
	LookingFor   string  `json:"lookingFor"`
	AgeFrom      byte    `json:"ageFrom"`
	AgeTo        byte    `json:"ageTo"`
	Distance     float64 `json:"distance"`
	Page         uint64  `json:"page"`
	Size         uint64  `json:"size"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}
