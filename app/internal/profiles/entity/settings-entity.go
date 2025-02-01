package entity

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/shared/enum"
	"time"
)

type SettingsEntity struct {
	Id             uint64           `json:"id"`
	TelegramUserId string           `json:"telegramUserId"`
	Measurement    enum.Measurement `json:"measurement"`
	CreatedAt      time.Time        `json:"createdAt"`
	UpdatedAt      time.Time        `json:"updatedAt"`
}
