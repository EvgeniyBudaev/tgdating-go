package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"time"
)

type ImageMapper struct {
}

func (pm *ImageMapper) MapToDeleteRequest(id uint64) *request.ImageDeleteRequestRepositoryDto {
	return &request.ImageDeleteRequestRepositoryDto{
		Id:        id,
		IsDeleted: true,
		UpdatedAt: time.Now().UTC(),
	}
}
