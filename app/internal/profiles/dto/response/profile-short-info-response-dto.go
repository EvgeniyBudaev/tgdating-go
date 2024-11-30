package response

type ProfileShortInfoResponseDto struct {
	SessionId string `json:"sessionId"`
	ImageUrl  string `json:"imageUrl"`
	IsFrozen  bool   `json:"isFrozen"`
	IsBlocked bool   `json:"isBlocked"`
}
