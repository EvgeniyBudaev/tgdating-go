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
	telegramResponse *response.ProfileTelegramResponseDto,
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
		CreatedAt:      profileEntity.CreatedAt,
		UpdatedAt:      profileEntity.UpdatedAt,
		LastOnline:     profileEntity.LastOnline,
		Navigator:      navigatorResponse,
		Telegram:       telegramResponse,
	}
}

func (pm *ProfileMapper) MapToAddResponse(profileEntity *entity.ProfileEntity) *response.ProfileAddResponseDto {
	return &response.ProfileAddResponseDto{
		SessionID: profileEntity.SessionID,
	}
}

func (pm *ProfileMapper) MapToAddRequest(pr *request.ProfileAddRequestDto) *entity.ProfileEntity {
	return &entity.ProfileEntity{
		SessionID:      pr.SessionID,
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

func (pm *ProfileMapper) MapToUpdateRequest(pr *request.ProfileUpdateRequestDto) *entity.ProfileEntity {
	return &entity.ProfileEntity{
		SessionID:   pr.SessionID,
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
