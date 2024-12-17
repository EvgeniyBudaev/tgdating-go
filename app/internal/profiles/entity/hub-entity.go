package entity

type HubContent struct {
	LikedTelegramUserId string `json:"likedTelegramUserId"`
	Message             string `json:"message"`
	Type                string `json:"type"`
	UserImageUrl        string `json:"userImageUrl"`
	Username            string `json:"username"`
}

type Hub struct {
	Broadcast chan *HubContent
}

func NewHub() *Hub {
	return &Hub{
		Broadcast: make(chan *HubContent, 5),
	}
}
