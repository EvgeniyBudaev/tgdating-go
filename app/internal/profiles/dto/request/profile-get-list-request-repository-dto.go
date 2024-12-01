package request

type ProfileGetListRequestRepositoryDto struct {
	TelegramUserId string  `json:"telegramUserId"`
	SearchGender   string  `json:"searchGender"`
	LookingFor     string  `json:"lookingFor"`
	AgeFrom        uint64  `json:"ageFrom"`
	AgeTo          uint64  `json:"ageTo"`
	Distance       float64 `json:"distance"`
	Page           uint64  `json:"page"`
	Size           uint64  `json:"size"`
}
