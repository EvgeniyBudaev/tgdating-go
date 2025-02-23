package request

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/shared/enum"
	"time"
)

type ComplaintAddRequestRepositoryDto struct {
	TelegramUserId         string         `json:"telegramUserId"`
	CriminalTelegramUserId string         `json:"criminalTelegramUserId"`
	Type                   enum.Complaint `json:"type"`
	Description            string         `json:"description"`
	CreatedAt              time.Time      `json:"createdAt"`
	UpdatedAt              time.Time      `json:"updatedAt"`
}
