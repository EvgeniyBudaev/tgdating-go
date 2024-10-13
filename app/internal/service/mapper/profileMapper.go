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
	pe *entity.ProfileEntity, nr *response.NavigatorResponseDto, fr *response.FilterResponseDto,
	tr *response.TelegramResponseDto, il []*entity.ImageEntity, isOnline bool,
) *response.ProfileResponseDto {
	return &response.ProfileResponseDto{
		SessionId:      pe.SessionId,
		DisplayName:    pe.DisplayName,
		Birthday:       pe.Birthday,
		Gender:         pe.Gender,
		Location:       pe.Location,
		Description:    pe.Description,
		Height:         pe.Height,
		Weight:         pe.Weight,
		IsDeleted:      pe.IsDeleted,
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
		SessionId:      pe.SessionId,
		DisplayName:    pe.DisplayName,
		Birthday:       pe.Birthday,
		Gender:         pe.Gender,
		Location:       pe.Location,
		Description:    pe.Description,
		Height:         pe.Height,
		Weight:         pe.Weight,
		IsDeleted:      pe.IsDeleted,
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
		SessionId: pe.SessionId,
		ImageUrl:  imageUrl,
		IsDeleted: pe.IsDeleted,
		IsBlocked: pe.IsBlocked,
	}
}

func (pm *ProfileMapper) MapToAddResponse(pe *entity.ProfileEntity) *response.ProfileAddResponseDto {
	return &response.ProfileAddResponseDto{
		SessionId: pe.SessionId,
	}
}

func (pm *ProfileMapper) MapToAddRequest(
	pr *request.ProfileAddRequestDto) *request.ProfileAddRequestRepositoryDto {
	return &request.ProfileAddRequestRepositoryDto{
		SessionId:      pr.SessionId,
		DisplayName:    pr.DisplayName,
		Birthday:       pr.Birthday,
		Gender:         pr.Gender,
		Location:       pr.Location,
		Description:    pr.Description,
		Height:         pr.Height,
		Weight:         pr.Weight,
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
	pr *request.ProfileUpdateRequestDto) *request.ProfileUpdateRequestRepositoryDto {
	return &request.ProfileUpdateRequestRepositoryDto{
		SessionId:   pr.SessionId,
		DisplayName: pr.DisplayName,
		Birthday:    pr.Birthday,
		Gender:      pr.Gender,
		Location:    pr.Location,
		Description: pr.Description,
		Height:      pr.Height,
		Weight:      pr.Weight,
		UpdatedAt:   time.Now().UTC(),
		LastOnline:  time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapToDeleteRequest(sessionId string) *request.ProfileDeleteRequestRepositoryDto {
	return &request.ProfileDeleteRequestRepositoryDto{
		SessionId:  sessionId,
		IsDeleted:  true,
		UpdatedAt:  time.Now().UTC(),
		LastOnline: time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapToListRequest(
	pr *request.ProfileGetListRequestDto) *request.ProfileGetListRequestRepositoryDto {
	return &request.ProfileGetListRequestRepositoryDto{
		SessionId:    pr.SessionId,
		SearchGender: pr.SearchGender,
		LookingFor:   pr.LookingFor,
		AgeFrom:      pr.AgeFrom,
		AgeTo:        pr.AgeTo,
		Distance:     pr.Distance,
		Page:         pr.Page,
		Size:         pr.Size,
	}
}
