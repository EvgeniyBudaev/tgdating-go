package request

import "time"

type PaymentAddRequestRepositoryDto struct {
	TelegramUserId string    `json:"telegramUserId"`
	Price          string    `json:"price"`
	Currency       string    `json:"currency"`
	Tariff         string    `json:"tariff"`
	CreatedAt      time.Time `json:"createdAt"`
}
