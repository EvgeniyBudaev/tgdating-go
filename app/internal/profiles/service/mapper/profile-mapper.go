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
	p *response.ProfileResponseRepositoryDto, i []*response.ImageResponseDto, isPremium bool,
) *response.ProfileResponseDto {
	s := &response.StatusResponseDto{
		IsBlocked:        p.Status.IsBlocked,
		IsFrozen:         p.Status.IsFrozen,
		IsHiddenAge:      p.Status.IsHiddenAge,
		IsHiddenDistance: p.Status.IsHiddenDistance,
		IsInvisible:      p.Status.IsInvisible,
		IsLeftHand:       p.Status.IsLeftHand,
		IsOnline:         p.Status.IsOnline,
		IsPremium:        isPremium,
	}
	st := &response.SettingsResponseDto{
		Measurement: p.Settings.Measurement,
	}
	return &response.ProfileResponseDto{
		TelegramUserId: p.TelegramUserId,
		DisplayName:    p.DisplayName,
		Age:            p.Age,
		Gender:         p.Gender,
		Description:    p.Description,
		Navigator:      p.Navigator,
		Filter:         p.Filter,
		Status:         s,
		Settings:       st,
		Images:         i,
	}
}

func (pm *ProfileMapper) MapToDetailResponse(
	p *response.ProfileDetailResponseRepositoryDto,
	il []*response.ImageResponseDto,
	isPremium bool,
) *response.ProfileDetailResponseDto {
	navigator := &response.NavigatorDetailResponseDto{
		CountryName: p.Navigator.CountryName,
		City:        p.Navigator.City,
		Distance:    p.Navigator.Distance,
	}
	s := &response.StatusResponseDto{
		IsBlocked:        p.Status.IsBlocked,
		IsFrozen:         p.Status.IsFrozen,
		IsHiddenAge:      p.Status.IsHiddenAge,
		IsHiddenDistance: p.Status.IsHiddenDistance,
		IsInvisible:      p.Status.IsInvisible,
		IsLeftHand:       p.Status.IsLeftHand,
		IsOnline:         p.Status.IsOnline,
		IsPremium:        isPremium,
	}
	st := &response.SettingsResponseDto{
		Measurement: p.Settings.Measurement,
	}
	return &response.ProfileDetailResponseDto{
		TelegramUserId: p.TelegramUserId,
		DisplayName:    p.DisplayName,
		Age:            p.Age,
		Description:    p.Description,
		Navigator:      navigator,
		Status:         s,
		Settings:       st,
		Block:          p.Block,
		Like:           p.Like,
		Images:         il,
	}
}
func (pm *ProfileMapper) MapToShortInfoResponse(
	p *response.ProfileShortInfoResponseRepositoryDto, pr *response.PremiumResponseDto) *response.ProfileShortInfoResponseDto {
	return &response.ProfileShortInfoResponseDto{
		TelegramUserId: p.TelegramUserId,
		IsBlocked:      p.IsBlocked,
		IsFrozen:       p.IsFrozen,
		IsPremium:      pr.IsPremium,
		AvailableUntil: pr.AvailableUntil,
		SearchGender:   p.SearchGender,
		AgeFrom:        p.AgeFrom,
		AgeTo:          p.AgeTo,
		Distance:       p.Distance,
		Page:           p.Page,
		Size:           p.Size,
		LanguageCode:   p.LanguageCode,
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
		Gender:         string(pr.Gender),
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
		Gender:         string(pr.Gender),
		Description:    pr.Description,
		UpdatedAt:      time.Now().UTC(),
		LastOnline:     time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapToListRequest(
	pr *request.ProfileGetListRequestDto, f *entity.FilterEntity) *request.ProfileGetListRequestRepositoryDto {
	return &request.ProfileGetListRequestRepositoryDto{
		TelegramUserId: pr.TelegramUserId,
		SearchGender:   f.SearchGender,
		AgeFrom:        f.AgeFrom,
		AgeTo:          f.AgeTo,
		Distance:       f.Distance,
		Page:           f.Page,
		Size:           f.Size,
	}
}
