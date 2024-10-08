package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"math"
	"time"
)

type NavigatorMapper struct {
}

func (pm *NavigatorMapper) MapToResponse(
	sessionId string, longitude float64, latitude float64) *response.NavigatorResponseDto {
	location := &entity.PointEntity{
		Longitude: longitude,
		Latitude:  latitude,
	}
	return &response.NavigatorResponseDto{
		SessionId: sessionId,
		Location:  location,
	}
}

func (pm *NavigatorMapper) MapToDetailResponse(distance float64) *response.NavigatorDetailResponseDto {
	roundedDistance := uint64(math.Ceil(distance))
	return &response.NavigatorDetailResponseDto{
		Distance: roundedDistance,
	}
}

func (pm *NavigatorMapper) MapToAddRequest(
	pr *request.ProfileAddRequestDto) *request.NavigatorAddRequestRepositoryDto {
	point := &entity.PointEntity{
		Longitude: pr.Longitude,
		Latitude:  pr.Latitude,
	}
	return &request.NavigatorAddRequestRepositoryDto{
		SessionId: pr.SessionId,
		Location:  point,
		IsDeleted: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (pm *NavigatorMapper) MapToUpdateRequest(
	sessionId string, longitude float64, latitude float64) *request.NavigatorUpdateRequestRepositoryDto {
	return &request.NavigatorUpdateRequestRepositoryDto{
		SessionId: sessionId,
		Longitude: longitude,
		Latitude:  latitude,
		UpdatedAt: time.Now().UTC(),
	}
}

func (pm *NavigatorMapper) MapToDeleteRequest(sessionId string) *request.NavigatorDeleteRequestRepositoryDto {
	return &request.NavigatorDeleteRequestRepositoryDto{
		SessionId: sessionId,
		IsDeleted: true,
		UpdatedAt: time.Now().UTC(),
	}
}
