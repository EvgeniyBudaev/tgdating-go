package request

type LikeAddRequestDto struct {
	SessionId      string `json:"sessionId"`
	LikedSessionId string `json:"likedSessionId"`
}
