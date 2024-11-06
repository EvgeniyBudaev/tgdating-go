package request

type FilterUpdateRequestDto struct {
	SessionId    string `json:"sessionId"`
	SearchGender string `json:"searchGender"`
	AgeFrom      byte   `json:"ageFrom"`
	AgeTo        byte   `json:"ageTo"`
}
