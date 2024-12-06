package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"time"
)

type ImageStatusMapper struct {
}

func (pm *ImageStatusMapper) MapToAddRequest(imageId uint64) *request.ImageStatusAddRequestRepositoryDto {
	return &request.ImageStatusAddRequestRepositoryDto{
		ImageId:   imageId,
		IsBlocked: false,
		IsPrimary: false,
		IsPrivate: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
