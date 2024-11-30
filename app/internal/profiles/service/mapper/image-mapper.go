package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
)

type ImageMapper struct {
}

func (pm *ImageMapper) MapToDeleteRequest(id uint64) *request.ImageDeleteRequestRepositoryDto {
	return &request.ImageDeleteRequestRepositoryDto{
		Id: id,
	}
}
