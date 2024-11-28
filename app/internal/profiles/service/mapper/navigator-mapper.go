package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"math"
	"time"
)

type NavigatorMapper struct {
}

func (pm *NavigatorMapper) MapToResponse(
	sessionId string, longitude float64, latitude float64) *response.NavigatorResponseDto {
	var navigatorResponse *response.NavigatorResponseDto
	if longitude != 0 || latitude != 0 {
		location := &entity.PointEntity{
			Longitude: longitude,
			Latitude:  latitude,
		}
		navigatorResponse = &response.NavigatorResponseDto{
			SessionId: sessionId,
			Location:  location,
		}
	}

	return navigatorResponse
}

func (pm *NavigatorMapper) MapToDetailResponse(distance float64) *response.NavigatorDetailResponseDto {
	roundedDistance := uint64(math.Ceil(distance))
	return &response.NavigatorDetailResponseDto{
		Distance: roundedDistance,
	}
}

func (pm *NavigatorMapper) MapToAddRequest(
	sessionId string, longitude, latitude float64) *request.NavigatorAddRequestRepositoryDto {
	point := &entity.PointEntity{
		Longitude: longitude,
		Latitude:  latitude,
	}
	return &request.NavigatorAddRequestRepositoryDto{
		SessionId: sessionId,
		Location:  point,
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