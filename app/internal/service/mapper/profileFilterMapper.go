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
	pe *entity.ProfileFilterEntity) *response.ProfileFilterResponseDto {
	return &response.ProfileFilterResponseDto{
		SessionId:    pe.SessionId,
		SearchGender: pe.SearchGender,
		LookingFor:   pe.LookingFor,
		AgeFrom:      pe.AgeFrom,
		AgeTo:        pe.AgeTo,
		Distance:     pe.Distance,
		Page:         pe.Page,
		Size:         pe.Size,
	}
}

func (pm *ProfileFilterMapper) MapToAddRequest(
	pr *request.ProfileAddRequestDto) *request.ProfileFilterAddRequestRepositoryDto {
	return &request.ProfileFilterAddRequestRepositoryDto{
		SessionId:    pr.SessionId,
		SearchGender: pr.SearchGender,
		LookingFor:   pr.LookingFor,
		AgeFrom:      pr.AgeFrom,
		AgeTo:        pr.AgeTo,
		Distance:     pr.Distance,
		Page:         pr.Page,
		Size:         pr.Size,
		IsDeleted:    false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (pm *ProfileFilterMapper) MapToUpdateRequest(
	pr *request.ProfileUpdateRequestDto) *request.ProfileFilterUpdateRequestRepositoryDto {
	return &request.ProfileFilterUpdateRequestRepositoryDto{
		SessionId:    pr.SessionId,
		SearchGender: pr.SearchGender,
		LookingFor:   pr.LookingFor,
		AgeFrom:      pr.AgeFrom,
		AgeTo:        pr.AgeTo,
		Distance:     pr.Distance,
		Page:         pr.Page,
		Size:         pr.Size,
		UpdatedAt:    time.Now(),
	}
}

func (pm *ProfileFilterMapper) MapToDeleteRequest(sessionId string) *request.ProfileFilterDeleteRequestRepositoryDto {
	return &request.ProfileFilterDeleteRequestRepositoryDto{
		SessionId: sessionId,
		IsDeleted: true,
		UpdatedAt: time.Now(),
	}
}
