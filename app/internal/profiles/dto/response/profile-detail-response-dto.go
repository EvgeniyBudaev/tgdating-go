package response

type ProfileDetailResponseDto struct {
	TelegramUserId string                      `json:"telegramUserId"`
	DisplayName    string                      `json:"displayName"`
	Age            uint64                      `json:"age"`
	Location       string                      `json:"location"`
	Description    string                      `json:"description"`
	Navigator      *NavigatorDetailResponseDto `json:"navigator"`
	Status         *StatusResponseDto          `json:"status"`
	Block          *BlockResponseDto           `json:"block"`
	Like           *LikeResponseDto            `json:"like"`
	Images         []*ImageResponseDto         `json:"images"`
}
