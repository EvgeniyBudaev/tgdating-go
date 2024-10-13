package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
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

func (pm *BlockMapper) MapToResponse(be *entity.BlockEntity) *response.BlockResponseDto {
	if be == nil {
		return nil
	}
	return &response.BlockResponseDto{
		IsBlocked: be.IsBlocked,
	}
}
