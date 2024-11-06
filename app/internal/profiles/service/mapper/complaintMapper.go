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
		SessionId:         pr.SessionId,
		CriminalSessionId: pr.CriminalSessionId,
		Reason:            pr.Reason,
		IsDeleted:         false,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}
