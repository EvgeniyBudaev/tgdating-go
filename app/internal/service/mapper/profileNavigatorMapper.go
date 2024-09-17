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

func (pm *ProfileNavigatorMapper) MapToAddRequest(
	profileAddRequestDto *request.ProfileAddRequestDto) *entity.ProfileNavigatorEntity {
	point := &entity.PointEntity{
		Latitude:  profileAddRequestDto.Latitude,
		Longitude: profileAddRequestDto.Longitude,
	}
	return &entity.ProfileNavigatorEntity{
		SessionID: profileAddRequestDto.SessionID,
		Location:  point,
		IsDeleted: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (pm *ProfileNavigatorMapper) MapToUpdateRequest(
	profileEntity *entity.ProfileEntity,
	profileUpdateRequestDto *request.ProfileUpdateRequestDto) *request.ProfileNavigatorUpdateRequestDto {
	return &request.ProfileNavigatorUpdateRequestDto{
		SessionID: profileEntity.SessionID,
		Longitude: profileUpdateRequestDto.Longitude,
		Latitude:  profileUpdateRequestDto.Latitude,
		UpdatedAt: time.Now().UTC(),
	}
}
