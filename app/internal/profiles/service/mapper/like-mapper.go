package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
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
		Id:                  pe.Id,
		TelegramUserId:      pe.TelegramUserId,
		LikedTelegramUserId: pe.LikedTelegramUserId,
		IsLiked:             pe.IsLiked,
		CreatedAt:           pe.CreatedAt,
		UpdatedAt:           pe.UpdatedAt,
	}
}

func (pm *LikeMapper) MapToAddRequest(
	pr *request.LikeAddRequestDto) *request.LikeAddRequestRepositoryDto {
	return &request.LikeAddRequestRepositoryDto{
		TelegramUserId:      pr.TelegramUserId,
		LikedTelegramUserId: pr.LikedTelegramUserId,
		IsLiked:             true,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
}

func (pm *LikeMapper) MapToUpdateRequest(
	pr *request.LikeUpdateRequestDto) *request.LikeUpdateRequestRepositoryDto {
	return &request.LikeUpdateRequestRepositoryDto{
		Id:        pr.Id,
		IsLiked:   pr.IsLiked,
		UpdatedAt: time.Now(),
	}
}
