package response

type ImageStatusResponseDto struct {
	IsBlocked bool `json:"isBlocked"`
	IsPrimary bool `json:"isPrimary"`
	IsPrivate bool `json:"isPrivate"`
}
