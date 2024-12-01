package entity

type HubContent struct {
	LikedTelegramUserId string `json:"likedTelegramUserId"`
	Message             string `json:"message"`
	Type                string `json:"type"`
	UserImageUrl        string `json:"userImageUrl"`
	Username            string `json:"username"`
}
