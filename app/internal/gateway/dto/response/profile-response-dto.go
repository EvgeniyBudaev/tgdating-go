package response

type ProfileResponseDto struct {
	TelegramUserId string                `json:"telegramUserId"`
	DisplayName    string                `json:"displayName"`
	Age            uint64                `json:"age"`
	Gender         string                `json:"gender"`
	Description    string                `json:"description"`
	Navigator      *NavigatorResponseDto `json:"navigator"`
	Filter         *FilterResponseDto    `json:"filter"`
	Status         *StatusResponseDto    `json:"status"`
	Images         []*ImageResponseDto   `json:"images"`
}
