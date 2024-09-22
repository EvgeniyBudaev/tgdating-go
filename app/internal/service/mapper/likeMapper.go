package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type LikeMapper struct {
}

func (pm *LikeMapper) MapToAddRequest(
	profileLikeAddRequestDto *request.LikeAddRequestDto) *request.LikeAddRequestRepositoryDto {
	return &request.LikeAddRequestRepositoryDto{
		SessionId:      profileLikeAddRequestDto.SessionId,
		LikedSessionId: profileLikeAddRequestDto.LikedSessionId,
		IsLiked:        true,
		IsDeleted:      false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (pm *LikeMapper) MapToAddResponse(
	profileLikeEntity *entity.LikeEntity) *response.LikeResponseDto {
	return &response.LikeResponseDto{
		Id:             profileLikeEntity.Id,
		SessionId:      profileLikeEntity.SessionId,
		LikedSessionId: profileLikeEntity.LikedSessionId,
		IsLiked:        profileLikeEntity.IsLiked,
		CreatedAt:      profileLikeEntity.CreatedAt,
		UpdatedAt:      profileLikeEntity.UpdatedAt,
	}
}
