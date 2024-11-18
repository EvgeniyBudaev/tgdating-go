package request

type LikeUpdateRequestDto struct {
	Id        uint64 `json:"id"`
	SessionId string `json:"sessionId"`
	IsLiked   bool   `json:"isLiked"`
}
