package response

type StatusResponseDto struct {
	IsBlocked      bool `json:"isBlocked"`
	IsFrozen       bool `json:"isFrozen"`
	IsInvisible    bool `json:"isInvisible"`
	IsOnline       bool `json:"isOnline"`
	IsPremium      bool `json:"isPremium"`
	IsShowDistance bool `json:"isShowDistance"`
}
