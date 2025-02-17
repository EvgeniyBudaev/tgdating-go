package response

type StatusResponseDto struct {
	IsBlocked        bool `json:"isBlocked"`
	IsFrozen         bool `json:"isFrozen"`
	IsHiddenAge      bool `json:"isHiddenAge"`
	IsHiddenDistance bool `json:"isHiddenDistance"`
	IsInvisible      bool `json:"isInvisible"`
	IsLeftHand       bool `json:"isLeftHand"`
	IsPremium        bool `json:"isPremium"`
}
