package response

type FilterResponseDto struct {
	SearchGender string  `json:"searchGender"`
	AgeFrom      uint64  `json:"ageFrom"`
	AgeTo        uint64  `json:"ageTo"`
	Distance     float64 `json:"distance"`
	Page         uint64  `json:"page"`
	Size         uint64  `json:"size"`
}
