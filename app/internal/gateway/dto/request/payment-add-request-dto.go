package request

import "github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enum"

type PaymentAddRequestDto struct {
	TelegramUserId string      `json:"telegramUserId"`
	Price          string      `json:"price"`
	Currency       string      `json:"currency"`
	Tariff         enum.Tariff `json:"tariff"`
}
