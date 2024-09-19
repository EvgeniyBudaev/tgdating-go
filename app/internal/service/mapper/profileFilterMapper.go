package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type ProfileFilterMapper struct {
}

func (pm *ProfileFilterMapper) MapToResponse(
	profileFilterEntity *entity.ProfileFilterEntity) *response.ProfileFilterResponseDto {
	return &response.ProfileFilterResponseDto{
		SessionID:    profileFilterEntity.SessionID,
		SearchGender: profileFilterEntity.SearchGender,
		LookingFor:   profileFilterEntity.LookingFor,
		AgeFrom:      profileFilterEntity.AgeFrom,
		AgeTo:        profileFilterEntity.AgeTo,
		Distance:     profileFilterEntity.Distance,
		Page:         profileFilterEntity.Page,
		Size:         profileFilterEntity.Size,
	}
}

func (pm *ProfileFilterMapper) MapToAddRequest(
	profileAddRequestDto *request.ProfileAddRequestDto) *request.ProfileFilterAddRequestRepositoryDto {
	return &request.ProfileFilterAddRequestRepositoryDto{
		SessionID:    profileAddRequestDto.SessionID,
		SearchGender: profileAddRequestDto.SearchGender,
		LookingFor:   profileAddRequestDto.LookingFor,
		AgeFrom:      profileAddRequestDto.AgeFrom,
		AgeTo:        profileAddRequestDto.AgeTo,
		Distance:     profileAddRequestDto.Distance,
		Page:         profileAddRequestDto.Page,
		Size:         profileAddRequestDto.Size,
		IsDeleted:    false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (pm *ProfileFilterMapper) MapToUpdateRequest(
	profileUpdateRequestDto *request.ProfileUpdateRequestDto) *request.ProfileFilterUpdateRequestRepositoryDto {
	return &request.ProfileFilterUpdateRequestRepositoryDto{
		SessionID:    profileUpdateRequestDto.SessionID,
		SearchGender: profileUpdateRequestDto.SearchGender,
		LookingFor:   profileUpdateRequestDto.LookingFor,
		AgeFrom:      profileUpdateRequestDto.AgeFrom,
		AgeTo:        profileUpdateRequestDto.AgeTo,
		Distance:     profileUpdateRequestDto.Distance,
		Page:         profileUpdateRequestDto.Page,
		Size:         profileUpdateRequestDto.Size,
		UpdatedAt:    time.Now(),
	}
}

func (pm *ProfileFilterMapper) MapToDeleteRequest(sessionId string) *request.ProfileFilterDeleteRequestRepositoryDto {
	return &request.ProfileFilterDeleteRequestRepositoryDto{
		SessionID: sessionId,
		IsDeleted: true,
		UpdatedAt: time.Now(),
	}
}
