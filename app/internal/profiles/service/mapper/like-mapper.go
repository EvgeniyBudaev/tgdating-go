package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"time"
)

type LikeMapper struct {
}

func (pm *LikeMapper) MapToAddRequest(
	pr *request.LikeAddRequestDto) *request.LikeAddRequestRepositoryDto {
	return &request.LikeAddRequestRepositoryDto{
		TelegramUserId:      pr.TelegramUserId,
		LikedTelegramUserId: pr.LikedTelegramUserId,
		IsLiked:             true,
		CreatedAt:           time.Now().UTC(),
		UpdatedAt:           time.Now().UTC(),
	}
}

func (pm *LikeMapper) MapToUpdateRequest(
	pr *request.LikeUpdateRequestDto) *request.LikeUpdateRequestRepositoryDto {
	return &request.LikeUpdateRequestRepositoryDto{
		Id:        pr.Id,
		IsLiked:   pr.IsLiked,
		UpdatedAt: time.Now().UTC(),
	}
}
