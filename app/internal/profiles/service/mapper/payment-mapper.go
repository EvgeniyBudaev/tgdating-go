package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"time"
)

type PaymentMapper struct {
}

func (pm *PaymentMapper) MapToAddRequest(
	pr *request.PaymentAddRequestDto, pl *entity.PaymentEntity) *request.PaymentAddRequestRepositoryDto {
	tariff := string(pr.Tariff)
	now := time.Now().UTC()
	if pl == nil {
		availableUntil := pm.calculateAvailableUntil(now, tariff)
		return &request.PaymentAddRequestRepositoryDto{
			TelegramUserId: pr.TelegramUserId,
			Price:          pr.Price,
			Currency:       pr.Currency,
			Tariff:         tariff,
			CreatedAt:      now,
			AvailableUntil: availableUntil,
		}
	}
	isBeforeDate := now.Before(pl.AvailableUntil)
	var availableUntil time.Time
	if isBeforeDate {
		availableUntil = pm.calculateAvailableUntil(pl.AvailableUntil, tariff)
	} else {
		availableUntil = pm.calculateAvailableUntil(now, tariff)
	}
	return &request.PaymentAddRequestRepositoryDto{
		TelegramUserId: pr.TelegramUserId,
		Price:          pr.Price,
		Currency:       pr.Currency,
		Tariff:         tariff,
		CreatedAt:      now,
		AvailableUntil: availableUntil,
	}
}

// MapToCheckPremium - checking if a subscription is active
func (pm *PaymentMapper) MapToCheckPremium(pr *entity.PaymentEntity) *response.PremiumResponseDto {
	now := time.Now().UTC()
	if pr == nil {
		return &response.PremiumResponseDto{
			IsPremium:      false,
			AvailableUntil: now,
		}
	}
	availableUntil := pr.AvailableUntil
	isPremium := now.Before(availableUntil)
	return &response.PremiumResponseDto{
		IsPremium:      isPremium,
		AvailableUntil: availableUntil,
	}
}

func (pm *PaymentMapper) calculateAvailableUntil(au time.Time, tariff string) time.Time {
	now := time.Now().UTC()
	switch tariff {
	case "free":
		return au.Add(60 * time.Second)
	case "month":
		return au.AddDate(0, 1, 0)
	case "threeMonths":
		return au.AddDate(0, 3, 0)
	case "year":
		return au.AddDate(1, 0, 0)
	default:
		return now
	}
}
