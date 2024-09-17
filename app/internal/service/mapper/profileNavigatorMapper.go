package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type ProfileNavigatorMapper struct {
}

func (pm *ProfileNavigatorMapper) MapToResponse(
	profileEntity *entity.ProfileEntity,
	navigatorEntity *entity.ProfileNavigatorEntity) *response.ProfileNavigatorResponseDto {
	return &response.ProfileNavigatorResponseDto{
		SessionID: profileEntity.SessionID,
		Location:  navigatorEntity.Location,
	}
}

func (pm *ProfileNavigatorMapper) MapToAddRequest(pr *request.ProfileAddRequestDto) *entity.ProfileNavigatorEntity {
	point := &entity.PointEntity{
		Latitude:  pr.Latitude,
		Longitude: pr.Longitude,
	}
	return &entity.ProfileNavigatorEntity{
		SessionID: pr.SessionID,
		Location:  point,
		IsDeleted: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (pm *ProfileNavigatorMapper) MapToUpdateRequest(
	profileEntity *entity.ProfileEntity,
	pr *request.ProfileUpdateRequestDto) *request.ProfileNavigatorUpdateRequestDto {
	return &request.ProfileNavigatorUpdateRequestDto{
		SessionID: profileEntity.SessionID,
		Longitude: pr.Longitude,
		Latitude:  pr.Latitude,
		UpdatedAt: time.Now().UTC(),
	}
}
