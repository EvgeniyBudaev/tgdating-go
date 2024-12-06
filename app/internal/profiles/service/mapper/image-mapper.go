package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
)

type ImageMapper struct {
}

func (pm *ImageMapper) MapToResponse(i *response.ImageResponseRepositoryDto) *response.ImageResponseDto {
	status := &response.ImageStatusResponseDto{
		IsBlocked: i.IsBlocked,
		IsPrimary: i.IsPrimary,
		IsPrivate: i.IsPrivate,
	}
	return &response.ImageResponseDto{
		Id:             i.Id,
		TelegramUserId: i.TelegramUserId,
		Name:           i.Name,
		Url:            i.Url,
		Size:           i.Size,
		Status:         status,
		CreatedAt:      i.CreatedAt,
		UpdatedAt:      i.UpdatedAt,
	}
}
