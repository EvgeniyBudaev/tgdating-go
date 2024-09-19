package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"time"
)

type ProfileImageMapper struct {
}

func (pm *ProfileImageMapper) MapToDeleteRequest(id uint64) *request.ProfileImageDeleteRequestRepositoryDto {
	return &request.ProfileImageDeleteRequestRepositoryDto{
		ID:        id,
		IsDeleted: true,
		UpdatedAt: time.Now().UTC(),
	}
}
