package response

type ProfileResponseRepositoryDto struct {
	TelegramUserId string                       `json:"telegramUserId"`
	DisplayName    string                       `json:"displayName"`
	Age            uint64                       `json:"age"`
	Gender         string                       `json:"gender"`
	Location       string                       `json:"location"`
	Description    string                       `json:"description"`
	Navigator      *NavigatorResponseDto        `json:"navigator"`
	Filter         *FilterResponseDto           `json:"filter"`
	Status         *StatusResponseRepositoryDto `json:"status"`
}
