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
	sessionId string, longitude float64, latitude float64) *response.ProfileNavigatorResponseDto {
	location := &entity.PointEntity{
		Longitude: longitude,
		Latitude:  latitude,
	}
	return &response.ProfileNavigatorResponseDto{
		SessionId: sessionId,
		Location:  location,
	}
}

func (pm *ProfileNavigatorMapper) MapToAddRequest(
	pr *request.ProfileAddRequestDto) *request.ProfileNavigatorAddRequestRepositoryDto {
	point := &entity.PointEntity{
		Longitude: pr.Longitude,
		Latitude:  pr.Latitude,
	}
	return &request.ProfileNavigatorAddRequestRepositoryDto{
		SessionId: pr.SessionId,
		Location:  point,
		IsDeleted: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (pm *ProfileNavigatorMapper) MapToUpdateRequest(
	sessionId string, longitude float64, latitude float64) *request.ProfileNavigatorUpdateRequestDto {
	return &request.ProfileNavigatorUpdateRequestDto{
		SessionId: sessionId,
		Longitude: longitude,
		Latitude:  latitude,
		UpdatedAt: time.Now().UTC(),
	}
}

func (pm *ProfileNavigatorMapper) MapToDeleteRequest(sessionId string) *request.ProfileNavigatorDeleteRequestDto {
	return &request.ProfileNavigatorDeleteRequestDto{
		SessionId: sessionId,
		IsDeleted: true,
		UpdatedAt: time.Now().UTC(),
	}
}
