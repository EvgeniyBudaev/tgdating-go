package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"time"
)

type BlockMapper struct {
}

func (pm *BlockMapper) MapToAddRequest(
	pr *request.BlockRequestDto) *request.BlockAddRequestRepositoryDto {
	return &request.BlockAddRequestRepositoryDto{
		SessionId:            pr.SessionId,
		BlockedUserSessionId: pr.BlockedUserSessionId,
		IsBlocked:            true,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}
}
