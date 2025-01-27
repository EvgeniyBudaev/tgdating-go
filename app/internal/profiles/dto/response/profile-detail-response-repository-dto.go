package response

type ProfileDetailResponseRepositoryDto struct {
	TelegramUserId string                                  `json:"telegramUserId"`
	DisplayName    string                                  `json:"displayName"`
	Age            uint64                                  `json:"age"`
	Location       string                                  `json:"location"`
	Description    string                                  `json:"description"`
	Navigator      *NavigatorDistanceResponseRepositoryDto `json:"navigator"`
	Status         *StatusResponseRepositoryDto            `json:"status"`
	Block          *BlockResponseDto                       `json:"block"`
	Like           *LikeResponseDto                        `json:"like"`
}
