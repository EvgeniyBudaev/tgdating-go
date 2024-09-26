package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"time"
)

type FilterMapper struct {
}

func (pm *FilterMapper) MapToResponse(
	pe *entity.FilterEntity) *response.FilterResponseDto {
	return &response.FilterResponseDto{
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

func (pm *FilterMapper) MapToAddRequest(
	pr *request.ProfileAddRequestDto) *request.FilterAddRequestRepositoryDto {
	return &request.FilterAddRequestRepositoryDto{
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

func (pm *FilterMapper) MapToUpdateRequest(
	fr *request.FilterUpdateRequestDto) *request.FilterUpdateRequestRepositoryDto {
	return &request.FilterUpdateRequestRepositoryDto{
		SessionId:    fr.SessionId,
		SearchGender: fr.SearchGender,
		LookingFor:   fr.LookingFor,
		AgeFrom:      fr.AgeFrom,
		AgeTo:        fr.AgeTo,
		Distance:     fr.Distance,
		Page:         fr.Page,
		Size:         fr.Size,
		UpdatedAt:    time.Now(),
	}
}

func (pm *FilterMapper) MapProfileToUpdateRequest(
	pr *request.ProfileUpdateRequestDto) *request.FilterUpdateRequestRepositoryDto {
	return &request.FilterUpdateRequestRepositoryDto{
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

func (pm *FilterMapper) MapToDeleteRequest(sessionId string) *request.FilterDeleteRequestRepositoryDto {
	return &request.FilterDeleteRequestRepositoryDto{
		SessionId: sessionId,
		IsDeleted: true,
		UpdatedAt: time.Now(),
	}
}
