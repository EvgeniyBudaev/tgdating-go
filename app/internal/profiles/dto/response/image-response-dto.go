package response

import "time"

type ImageResponseDto struct {
	Id             uint64                  `json:"id"`
	TelegramUserId string                  `json:"telegramUserId"`
	Name           string                  `json:"name"`
	Url            string                  `json:"url"`
	Size           int64                   `json:"size"`
	Status         *ImageStatusResponseDto `json:"status"`
	CreatedAt      time.Time               `json:"createdAt"`
	UpdatedAt      time.Time               `json:"updatedAt"`
}
