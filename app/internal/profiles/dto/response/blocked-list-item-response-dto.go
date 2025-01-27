package response

type BlockedListItemResponseDto struct {
	BlockedTelegramUserId string `json:"blockedTelegramUserId"`
	Url                   string `json:"url"`
}
