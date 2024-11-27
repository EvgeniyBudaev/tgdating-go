package response

type ProfileShortInfoResponseDto struct {
	SessionId string `json:"sessionId"`
	ImageUrl  string `json:"imageUrl"`
	IsDeleted bool   `json:"isDeleted"`
	IsBlocked bool   `json:"isBlocked"`
}
