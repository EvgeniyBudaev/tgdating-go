package entity

import "time"

type NavigatorEntity struct {
	Id             uint64       `json:"id"`
	TelegramUserId string       `json:"telegramUserId"`
	CountryCode    *string      `json:"countryCode"`
	CountryName    *string      `json:"countryName"`
	City           *string      `json:"city"`
	Location       *PointEntity `json:"location"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt"`
}
