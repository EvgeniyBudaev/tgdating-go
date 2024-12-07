package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"time"
)

type ComplaintMapper struct {
}

func (pm *ComplaintMapper) MapToAddRequest(
	pr *request.ComplaintAddRequestDto) *request.ComplaintAddRequestRepositoryDto {
	return &request.ComplaintAddRequestRepositoryDto{
		TelegramUserId:         pr.TelegramUserId,
		CriminalTelegramUserId: pr.CriminalTelegramUserId,
		Reason:                 pr.Reason,
		CreatedAt:              time.Now().UTC(),
		UpdatedAt:              time.Now().UTC(),
	}
}
