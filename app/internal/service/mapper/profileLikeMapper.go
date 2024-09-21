package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type ProfileLikeMapper struct {
}

func (pm *ProfileLikeMapper) MapToAddRequest(
	profileLikeAddRequestDto *request.ProfileLikeAddRequestDto) *request.ProfileLikeAddRequestRepositoryDto {
	return &request.ProfileLikeAddRequestRepositoryDto{
		SessionId:      profileLikeAddRequestDto.SessionId,
		LikedSessionId: profileLikeAddRequestDto.LikedSessionId,
		IsLiked:        true,
		IsDeleted:      false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (pm *ProfileLikeMapper) MapToAddResponse(
	profileLikeEntity *entity.ProfileLikeEntity) *response.ProfileLikeResponseDto {
	return &response.ProfileLikeResponseDto{
		Id:             profileLikeEntity.Id,
		SessionId:      profileLikeEntity.SessionId,
		LikedSessionId: profileLikeEntity.LikedSessionId,
		IsLiked:        profileLikeEntity.IsLiked,
		CreatedAt:      profileLikeEntity.CreatedAt,
		UpdatedAt:      profileLikeEntity.UpdatedAt,
	}
}
