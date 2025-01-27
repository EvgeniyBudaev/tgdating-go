package response

type BlockedListResponseDto struct {
	Content []*BlockedListItemResponseDto `json:"content"`
}
