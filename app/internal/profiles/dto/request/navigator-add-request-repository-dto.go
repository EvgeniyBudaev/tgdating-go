package request

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"time"
)

type NavigatorAddRequestRepositoryDto struct {
	TelegramUserId string              `json:"telegramUserId"`
	CountryCode    *string             `json:"countryCode"`
	CountryName    *string             `json:"countryName"`
	City           *string             `json:"city"`
	Location       *entity.PointEntity `json:"location"`
	CreatedAt      time.Time           `json:"createdAt"`
	UpdatedAt      time.Time           `json:"updatedAt"`
}
