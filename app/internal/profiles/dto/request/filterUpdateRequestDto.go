package request

type FilterUpdateRequestDto struct {
	SessionId    string `json:"sessionId"`
	SearchGender string `json:"searchGender"`
	AgeFrom      uint64 `json:"ageFrom"`
	AgeTo        uint64 `json:"ageTo"`
}
