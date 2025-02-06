package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"time"
)

type BlockMapper struct {
}

func (pm *BlockMapper) MapToAddRequest(
	pr *request.BlockAddRequestDto, initiatorId *string) *request.BlockAddRequestRepositoryDto {
	return &request.BlockAddRequestRepositoryDto{
		TelegramUserId:        pr.TelegramUserId,
		BlockedTelegramUserId: pr.BlockedTelegramUserId,
		InitiatorId:           initiatorId,
		IsBlocked:             true,
		CreatedAt:             time.Now().UTC(),
		UpdatedAt:             time.Now().UTC(),
	}
}

func (pm *BlockMapper) MapToUpdateRequest(
	pr *request.BlockAddRequestDto, initiatorId *string) *request.BlockUpdateRequestRepositoryDto {
	return &request.BlockUpdateRequestRepositoryDto{
		TelegramUserId:        pr.TelegramUserId,
		BlockedTelegramUserId: pr.BlockedTelegramUserId,
		InitiatorId:           initiatorId,
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
