package response

type ProfileShortInfoResponseDto struct {
	TelegramUserId string  `json:"telegramUserId"`
	IsBlocked      bool    `json:"isBlocked"`
	IsFrozen       bool    `json:"isFrozen"`
	SearchGender   string  `json:"searchGender"`
	LookingFor     string  `json:"lookingFor"`
	AgeFrom        uint64  `json:"ageFrom"`
	AgeTo          uint64  `json:"ageTo"`
	Distance       float64 `json:"distance"`
	Page           uint64  `json:"page"`
	Size           uint64  `json:"size"`
}
