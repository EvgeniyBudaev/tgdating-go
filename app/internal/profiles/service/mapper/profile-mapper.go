package mapper

import (
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"time"
)

type ProfileMapper struct {
}

func (pm *ProfileMapper) MapToResponse(
	pe *entity.ProfileEntity, nr *response.NavigatorResponseDto, fr *response.FilterResponseDto,
	tr *response.TelegramResponseDto, sr *response.StatusResponseDto, il []*response.ImageResponseDto,
) *response.ProfileResponseDto {
	return &response.ProfileResponseDto{
		TelegramUserId: pe.TelegramUserId,
		DisplayName:    pe.DisplayName,
		Birthday:       pe.Birthday,
		Gender:         pe.Gender,
		Location:       pe.Location,
		Description:    pe.Description,
		Height:         pe.Height,
		Weight:         pe.Weight,
		CreatedAt:      pe.CreatedAt,
		UpdatedAt:      pe.UpdatedAt,
		LastOnline:     pe.LastOnline,
		Navigator:      nr,
		Filter:         fr,
		Telegram:       tr,
		Status:         sr,
		Images:         il,
	}
}

func (pm *ProfileMapper) MapToDetailResponse(
	pe *entity.ProfileEntity, nr *response.NavigatorDetailResponseDto, br *response.BlockResponseDto,
	lr *response.LikeResponseDto, tr *response.TelegramResponseDto, sr *response.StatusResponseDto,
	il []*response.ImageResponseDto,
) *response.ProfileDetailResponseDto {
	return &response.ProfileDetailResponseDto{
		TelegramUserId: pe.TelegramUserId,
		DisplayName:    pe.DisplayName,
		Birthday:       pe.Birthday,
		Gender:         pe.Gender,
		Location:       pe.Location,
		Description:    pe.Description,
		Height:         pe.Height,
		Weight:         pe.Weight,
		CreatedAt:      pe.CreatedAt,
		UpdatedAt:      pe.UpdatedAt,
		LastOnline:     pe.LastOnline,
		Navigator:      nr,
		Telegram:       tr,
		Status:         sr,
		Block:          br,
		Like:           lr,
		Images:         il,
	}
}

func (pm *ProfileMapper) MapToAddResponse(pe *entity.ProfileEntity) *response.ProfileAddResponseDto {
	return &response.ProfileAddResponseDto{
		TelegramUserId: pe.TelegramUserId,
	}
}

func (pm *ProfileMapper) MapToAddRequest(
	pr *request.ProfileAddRequestDto) *request.ProfileAddRequestRepositoryDto {
	return &request.ProfileAddRequestRepositoryDto{
		TelegramUserId: pr.TelegramUserId,
		DisplayName:    pr.DisplayName,
		Birthday:       pr.Birthday,
		Gender:         pr.Gender,
		Location:       pr.Location,
		Description:    pr.Description,
		Height:         pr.Height,
		Weight:         pr.Weight,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		LastOnline:     time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapToUpdateRequest(
	pr *request.ProfileUpdateRequestDto) *request.ProfileUpdateRequestRepositoryDto {
	return &request.ProfileUpdateRequestRepositoryDto{
		TelegramUserId: pr.TelegramUserId,
		DisplayName:    pr.DisplayName,
		Birthday:       pr.Birthday,
		Gender:         pr.Gender,
		Location:       pr.Location,
		Description:    pr.Description,
		Height:         pr.Height,
		Weight:         pr.Weight,
		UpdatedAt:      time.Now().UTC(),
		LastOnline:     time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapToListRequest(
	pr *request.ProfileGetListRequestDto) *request.ProfileGetListRequestRepositoryDto {
	return &request.ProfileGetListRequestRepositoryDto{
		TelegramUserId: pr.TelegramUserId,
		SearchGender:   pr.SearchGender,
		LookingFor:     pr.LookingFor,
		AgeFrom:        pr.AgeFrom,
		AgeTo:          pr.AgeTo,
		Distance:       pr.Distance,
		Page:           pr.Page,
		Size:           pr.Size,
	}
}
