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
	tr *response.TelegramResponseDto, il []*entity.ImageEntity, isOnline bool,
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
		IsFrozen:       pe.IsFrozen,
		IsBlocked:      pe.IsBlocked,
		IsPremium:      pe.IsPremium,
		IsShowDistance: pe.IsShowDistance,
		IsInvisible:    pe.IsInvisible,
		IsOnline:       isOnline,
		CreatedAt:      pe.CreatedAt,
		UpdatedAt:      pe.UpdatedAt,
		LastOnline:     pe.LastOnline,
		Navigator:      nr,
		Filter:         fr,
		Telegram:       tr,
		Images:         il,
	}
}

func (pm *ProfileMapper) MapToDetailResponse(
	pe *entity.ProfileEntity, nr *response.NavigatorDetailResponseDto, br *response.BlockResponseDto,
	lr *response.LikeResponseDto, tr *response.TelegramResponseDto, il []*entity.ImageEntity, isOnline bool,
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
		IsFrozen:       pe.IsFrozen,
		IsBlocked:      pe.IsBlocked,
		IsPremium:      pe.IsPremium,
		IsShowDistance: pe.IsShowDistance,
		IsInvisible:    pe.IsInvisible,
		IsOnline:       isOnline,
		CreatedAt:      pe.CreatedAt,
		UpdatedAt:      pe.UpdatedAt,
		LastOnline:     pe.LastOnline,
		Navigator:      nr,
		Block:          br,
		Like:           lr,
		Telegram:       tr,
		Images:         il,
	}
}

func (pm *ProfileMapper) MapToShortInfoResponse(pe *entity.ProfileEntity, imageUrl string) *response.ProfileShortInfoResponseDto {
	return &response.ProfileShortInfoResponseDto{
		TelegramUserId: pe.TelegramUserId,
		ImageUrl:       imageUrl,
		IsFrozen:       pe.IsFrozen,
		IsBlocked:      pe.IsBlocked,
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
		IsFrozen:       false,
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

func (pm *ProfileMapper) MapToFreezeRequest(telegramUserId string) *request.ProfileFreezeRequestRepositoryDto {
	return &request.ProfileFreezeRequestRepositoryDto{
		TelegramUserId: telegramUserId,
		IsFrozen:       true,
		UpdatedAt:      time.Now().UTC(),
		LastOnline:     time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapToRestoreRequest(telegramUserId string) *request.ProfileRestoreRequestRepositoryDto {
	return &request.ProfileRestoreRequestRepositoryDto{
		TelegramUserId: telegramUserId,
		IsFrozen:       false,
		UpdatedAt:      time.Now().UTC(),
		LastOnline:     time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapToListRequest(
	pr *entity.FilterEntity) *request.ProfileGetListRequestRepositoryDto {
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
