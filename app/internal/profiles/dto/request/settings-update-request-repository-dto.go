package request

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/shared/enum"
	"time"
)

type SettingsUpdateRequestRepositoryDto struct {
	TelegramUserId string           `json:"telegramUserId"`
	Measurement    enum.Measurement `json:"measurement"`
	UpdatedAt      time.Time        `json:"updatedAt"`
}
