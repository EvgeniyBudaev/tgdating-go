package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"time"
)

type NavigatorMapper struct {
}

func (pm *NavigatorMapper) MapToResponse(
	telegramUserId string, longitude float64, latitude float64) *response.NavigatorResponseDto {
	var navigatorResponse *response.NavigatorResponseDto
	if longitude != 0 || latitude != 0 {
		location := &entity.PointEntity{
			Longitude: longitude,
			Latitude:  latitude,
		}
		navigatorResponse = &response.NavigatorResponseDto{
			Location: location,
		}
	}

	return navigatorResponse
}

func (pm *NavigatorMapper) MapToAddRequest(telegramUserId string, countryCode *string, longitude,
	latitude float64) *request.NavigatorAddRequestRepositoryDto {
	point := &entity.PointEntity{
		Longitude: longitude,
		Latitude:  latitude,
	}
	return &request.NavigatorAddRequestRepositoryDto{
		TelegramUserId: telegramUserId,
		CountryCode:    countryCode,
		Location:       point,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
}

func (pm *NavigatorMapper) MapToUpdateRequest(telegramUserId string, countryCode *string, longitude float64,
	latitude float64) *request.NavigatorUpdateRequestRepositoryDto {
	return &request.NavigatorUpdateRequestRepositoryDto{
		TelegramUserId: telegramUserId,
		CountryCode:    countryCode,
		Longitude:      longitude,
		Latitude:       latitude,
		UpdatedAt:      time.Now().UTC(),
	}
}
