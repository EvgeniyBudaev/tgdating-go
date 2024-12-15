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
	p *response.ProfileResponseRepositoryDto, i []*response.ImageResponseDto,
) *response.ProfileResponseDto {
	return &response.ProfileResponseDto{
		TelegramUserId: p.TelegramUserId,
		DisplayName:    p.DisplayName,
		Age:            p.Age,
		Gender:         p.Gender,
		Location:       p.Location,
		Description:    p.Description,
		Navigator:      p.Navigator,
		Filter:         p.Filter,
		Status:         p.Status,
		Images:         i,
	}
}

func (pm *ProfileMapper) MapToDetailResponse(
	p *response.ProfileDetailResponseRepositoryDto,
	il []*response.ImageResponseDto,
) *response.ProfileDetailResponseDto {
	navigator := &response.NavigatorDetailResponseDto{
		Distance: p.Navigator.Distance,
	}
	return &response.ProfileDetailResponseDto{
		TelegramUserId: p.TelegramUserId,
		DisplayName:    p.DisplayName,
		Age:            p.Age,
		Location:       p.Location,
		Description:    p.Description,
		Navigator:      navigator,
		Status:         p.Status,
		Block:          p.Block,
		Like:           p.Like,
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
		Age:            pr.Age,
		Gender:         pr.Gender,
		Location:       pr.Location,
		Description:    pr.Description,
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
		Age:            pr.Age,
		Gender:         pr.Gender,
		Location:       pr.Location,
		Description:    pr.Description,
		UpdatedAt:      time.Now().UTC(),
		LastOnline:     time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapToListRequest(
	pr *request.ProfileGetListRequestDto) *request.ProfileGetListRequestRepositoryDto {
	return &request.ProfileGetListRequestRepositoryDto{
		TelegramUserId: pr.TelegramUserId,
		SearchGender:   pr.SearchGender,
		AgeFrom:        pr.AgeFrom,
		AgeTo:          pr.AgeTo,
		Distance:       pr.Distance,
		Page:           pr.Page,
		Size:           pr.Size,
	}
}
