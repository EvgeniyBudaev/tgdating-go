package entity

import "time"

type PaymentEntity struct {
	Id             uint64    `json:"id"`
	TelegramUserId string    `json:"telegramUserId"`
	Price          string    `json:"price"`
	Currency       string    `json:"currency"`
	Tariff         string    `json:"tariff"`
	CreatedAt      time.Time `json:"createdAt"`
	AvailableUntil time.Time `json:"availableUntil"`
}
