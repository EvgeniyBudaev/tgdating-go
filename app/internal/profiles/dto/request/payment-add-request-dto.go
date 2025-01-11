package request

type PaymentAddRequestDto struct {
	TelegramUserId string `json:"telegramUserId"`
	Price          string `json:"price"`
	Currency       string `json:"currency"`
	Tariff         string `json:"tariff"`
}
