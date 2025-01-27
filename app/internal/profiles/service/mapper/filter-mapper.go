package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"time"
)

type FilterMapper struct {
}

func (pm *FilterMapper) MapToResponse(
	pe *entity.FilterEntity) *response.FilterResponseDto {
	return &response.FilterResponseDto{
		SearchGender: pe.SearchGender,
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
		TelegramUserId: pr.TelegramUserId,
		SearchGender:   string(pr.SearchGender),
		AgeFrom:        pr.AgeFrom,
		AgeTo:          pr.AgeTo,
		Distance:       pr.Distance,
		Page:           pr.Page,
		Size:           pr.Size,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
}

func (pm *FilterMapper) MapToUpdateRequest(
	fr *request.FilterUpdateRequestDto) *request.FilterUpdateRequestRepositoryDto {
	return &request.FilterUpdateRequestRepositoryDto{
		TelegramUserId: fr.TelegramUserId,
		SearchGender:   fr.SearchGender,
		AgeFrom:        fr.AgeFrom,
		AgeTo:          fr.AgeTo,
		UpdatedAt:      time.Now().UTC(),
	}
}

func (pm *FilterMapper) MapProfileToUpdateRequest(
	pr *request.ProfileUpdateRequestDto) *request.FilterUpdateRequestRepositoryDto {
	return &request.FilterUpdateRequestRepositoryDto{
		TelegramUserId: pr.TelegramUserId,
		SearchGender:   string(pr.SearchGender),
		AgeFrom:        pr.AgeFrom,
		AgeTo:          pr.AgeTo,
		UpdatedAt:      time.Now().UTC(),
	}
}
