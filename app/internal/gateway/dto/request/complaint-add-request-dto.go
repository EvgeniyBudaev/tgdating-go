package request

import "github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enum"

type ComplaintAddRequestDto struct {
	TelegramUserId         string         `json:"telegramUserId"`
	CriminalTelegramUserId string         `json:"criminalTelegramUserId"`
	Type                   enum.Complaint `json:"type"`
	Description            string         `json:"description"`
}
