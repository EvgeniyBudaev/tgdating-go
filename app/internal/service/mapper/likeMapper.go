package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type LikeMapper struct {
}

func (pm *LikeMapper) MapToResponse(
	pe *entity.LikeEntity) *response.LikeResponseDto {
	if pe == nil {
		return nil
	}
	return &response.LikeResponseDto{
		Id:             pe.Id,
		SessionId:      pe.SessionId,
		LikedSessionId: pe.LikedSessionId,
		IsLiked:        pe.IsLiked,
		CreatedAt:      pe.CreatedAt,
		UpdatedAt:      pe.UpdatedAt,
	}
}

func (pm *LikeMapper) MapToAddRequest(
	pr *request.LikeAddRequestDto) *request.LikeAddRequestRepositoryDto {
	return &request.LikeAddRequestRepositoryDto{
		SessionId:      pr.SessionId,
		LikedSessionId: pr.LikedSessionId,
		IsLiked:        true,
		IsDeleted:      false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
