package entity

type HubContent struct {
	LikedUserId  uint64 `json:"likedUserId"`
	Message      string `json:"message"`
	Type         string `json:"type"`
	UserImageUrl string `json:"userImageUrl"`
	Username     string `json:"username"`
}
