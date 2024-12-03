package response

type StatusResponseDto struct {
	IsFrozen       bool `json:"isFrozen"`
	IsBlocked      bool `json:"isBlocked"`
	IsPremium      bool `json:"isPremium"`
	IsShowDistance bool `json:"isShowDistance"`
	IsInvisible    bool `json:"isInvisible"`
}
