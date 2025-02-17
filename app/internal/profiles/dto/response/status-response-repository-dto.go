package response

type StatusResponseRepositoryDto struct {
	IsBlocked        bool `json:"isBlocked"`
	IsFrozen         bool `json:"isFrozen"`
	IsHiddenAge      bool `json:"isHiddenAge"`
	IsHiddenDistance bool `json:"isHiddenDistance"`
	IsInvisible      bool `json:"isInvisible"`
	IsLeftHand       bool `json:"isLeftHand"`
}
