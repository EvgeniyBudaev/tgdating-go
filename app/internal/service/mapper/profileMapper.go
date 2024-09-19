package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type ProfileMapper struct {
}

func (pm *ProfileMapper) MapToResponse(
	profileEntity *entity.ProfileEntity,
	navigatorResponse *response.ProfileNavigatorResponseDto,
	filterResponse *response.ProfileFilterResponseDto,
	telegramResponse *response.ProfileTelegramResponseDto,
	isOnline bool,
) *response.ProfileUpdateResponseDto {
	return &response.ProfileUpdateResponseDto{
		SessionID:      profileEntity.SessionID,
		DisplayName:    profileEntity.DisplayName,
		Birthday:       profileEntity.Birthday,
		Gender:         profileEntity.Gender,
		Location:       profileEntity.Location,
		Description:    profileEntity.Description,
		Height:         profileEntity.Height,
		Weight:         profileEntity.Weight,
		IsDeleted:      profileEntity.IsDeleted,
		IsBlocked:      profileEntity.IsBlocked,
		IsPremium:      profileEntity.IsPremium,
		IsShowDistance: profileEntity.IsShowDistance,
		IsInvisible:    profileEntity.IsInvisible,
		IsOnline:       isOnline,
		CreatedAt:      profileEntity.CreatedAt,
		UpdatedAt:      profileEntity.UpdatedAt,
		LastOnline:     profileEntity.LastOnline,
		Navigator:      navigatorResponse,
		Filter:         filterResponse,
		Telegram:       telegramResponse,
	}
}

func (pm *ProfileMapper) MapToAddResponse(profileEntity *entity.ProfileEntity) *response.ProfileAddResponseDto {
	return &response.ProfileAddResponseDto{
		SessionID: profileEntity.SessionID,
	}
}

func (pm *ProfileMapper) MapToAddRequest(
	profileAddRequestDto *request.ProfileAddRequestDto) *request.ProfileAddRequestRepositoryDto {
	return &request.ProfileAddRequestRepositoryDto{
		SessionID:      profileAddRequestDto.SessionID,
		DisplayName:    profileAddRequestDto.DisplayName,
		Birthday:       profileAddRequestDto.Birthday,
		Gender:         profileAddRequestDto.Gender,
		Location:       profileAddRequestDto.Location,
		Description:    profileAddRequestDto.Description,
		Height:         profileAddRequestDto.Height,
		Weight:         profileAddRequestDto.Weight,
		IsDeleted:      false,
		IsBlocked:      false,
		IsPremium:      false,
		IsShowDistance: true,
		IsInvisible:    false,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		LastOnline:     time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapToUpdateRequest(
	profileUpdateRequestDto *request.ProfileUpdateRequestDto) *request.ProfileUpdateRequestRepositoryDto {
	return &request.ProfileUpdateRequestRepositoryDto{
		SessionID:   profileUpdateRequestDto.SessionID,
		DisplayName: profileUpdateRequestDto.DisplayName,
		Birthday:    profileUpdateRequestDto.Birthday,
		Gender:      profileUpdateRequestDto.Gender,
		Location:    profileUpdateRequestDto.Location,
		Description: profileUpdateRequestDto.Description,
		Height:      profileUpdateRequestDto.Height,
		Weight:      profileUpdateRequestDto.Weight,
		UpdatedAt:   time.Now().UTC(),
		LastOnline:  time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapToDeleteRequest(sessionId string) *request.ProfileDeleteRequestRepositoryDto {
	return &request.ProfileDeleteRequestRepositoryDto{
		SessionID:  sessionId,
		IsDeleted:  true,
		UpdatedAt:  time.Now().UTC(),
		LastOnline: time.Now().UTC(),
	}
}
