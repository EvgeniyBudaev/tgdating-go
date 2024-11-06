package request

type LikeUpdateRequestDto struct {
	Id        uint64 `json:"id"`
	IsLiked   bool   `json:"isLiked"`
	SessionId string `json:"sessionId"`
}
