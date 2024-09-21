package request

type ProfileLikeAddRequestDto struct {
	SessionId      string `json:"sessionId"`
	LikedSessionId string `json:"likedSessionId"`
}
