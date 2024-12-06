package response

type StatusResponseDto struct {
	IsBlocked      bool `json:"isBlocked"`
	IsFrozen       bool `json:"isFrozen"`
	IsOnline       bool `json:"isOnline"`
	IsPremium      bool `json:"isPremium"`
	IsShowDistance bool `json:"isShowDistance"`
	IsInvisible    bool `json:"isInvisible"`
}
