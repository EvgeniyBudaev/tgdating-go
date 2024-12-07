package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
)

type ImageMapper struct {
}

func (pm *ImageMapper) MapToResponse(i *response.ImageResponseRepositoryDto) *response.ImageResponseDto {
	return &response.ImageResponseDto{
		Id:   i.Id,
		Name: i.Name,
		Url:  i.Url,
	}
}
