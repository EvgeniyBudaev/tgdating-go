package request

type FilterUpdateRequestDto struct {
	TelegramUserId string  `json:"telegramUserId"`
	SearchGender   string  `json:"searchGender"`
	AgeFrom        uint64  `json:"ageFrom"`
	AgeTo          uint64  `json:"ageTo"`
	Distance       float64 `json:"distance"`
	IsLiked        bool    `json:"isLiked"`
	IsOnline       bool    `json:"isOnline"`
}
