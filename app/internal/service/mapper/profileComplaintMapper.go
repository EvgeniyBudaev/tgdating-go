package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"time"
)

type ProfileComplaintMapper struct {
}

func (pm *ProfileComplaintMapper) MapToAddRequest(
	pr *request.ProfileComplaintAddRequestDto) *request.ProfileComplaintAddRequestRepositoryDto {
	return &request.ProfileComplaintAddRequestRepositoryDto{
		SessionId:         pr.SessionId,
		CriminalSessionId: pr.CriminalSessionId,
		Reason:            pr.Reason,
		IsDeleted:         false,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}
