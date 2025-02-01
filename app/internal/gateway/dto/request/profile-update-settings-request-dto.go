package request

import "github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enum"

type ProfileUpdateSettingsRequestDto struct {
	TelegramUserId string           `json:"telegramUserId"`
	IsHiddenAge    bool             `json:"isHiddenAge"`
	Measurement    enum.Measurement `json:"measurement"`
}
