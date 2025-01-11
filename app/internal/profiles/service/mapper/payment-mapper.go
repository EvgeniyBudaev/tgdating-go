package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"time"
)

type PaymentMapper struct {
}

func (pm *PaymentMapper) MapToAddRequest(
	pr *request.PaymentAddRequestDto) *request.PaymentAddRequestRepositoryDto {
	return &request.PaymentAddRequestRepositoryDto{
		TelegramUserId: pr.TelegramUserId,
		Price:          pr.Price,
		Currency:       pr.Currency,
		Tariff:         pr.Tariff,
		CreatedAt:      time.Now().UTC(),
	}
}
