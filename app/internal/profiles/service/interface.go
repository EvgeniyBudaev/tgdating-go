package service

import (
	"context"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
)

type ProfileRepository interface {
	Add(ctx context.Context, p *request.ProfileAddRequestRepositoryDto) (*response.ResponseDto, error)
	Update(ctx context.Context,
		p *request.ProfileUpdateRequestRepositoryDto) (*response.ProfileResponseRepositoryDto, error)
	Delete(ctx context.Context, p *request.ProfileDeleteRequestDto) (*response.ResponseDto, error)
	GetProfile(ctx context.Context, telegramUserId string) (*response.ProfileResponseRepositoryDto, error)
	GetDetail(ctx context.Context,
		telegramUserId, viewedTelegramUserId string) (*response.ProfileDetailResponseRepositoryDto, error)
	GetShortInfo(ctx context.Context, telegramUserId string) (*response.ProfileShortInfoResponseDto, error)
	SelectList(ctx context.Context,
		pr *request.ProfileGetListRequestRepositoryDto) (*response.ProfileListResponseRepositoryDto, error)
	UpdateLastOnline(ctx context.Context, p *request.ProfileUpdateLastOnlineRequestRepositoryDto) error
}

type NavigatorRepository interface {
	Add(ctx context.Context, p *request.NavigatorAddRequestRepositoryDto) (*response.ResponseDto, error)
	Update(ctx context.Context, p *request.NavigatorUpdateRequestRepositoryDto) (*entity.NavigatorEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.NavigatorEntity, error)
	FindByTelegramUserId(ctx context.Context, telegramUserId string) (*entity.NavigatorEntity, error)
	CheckNavigatorExists(ctx context.Context, telegramUserId string) (*response.ResponseDto, error)
}

type FilterRepository interface {
	Add(ctx context.Context, p *request.FilterAddRequestRepositoryDto) (*response.ResponseDto, error)
	Update(ctx context.Context, p *request.FilterUpdateRequestRepositoryDto) (*entity.FilterEntity, error)
	FindByTelegramUserId(ctx context.Context, telegramUserId string) (*entity.FilterEntity, error)
}

type TelegramRepository interface {
	Add(ctx context.Context, p *request.TelegramAddRequestRepositoryDto) (*response.ResponseDto, error)
	Update(ctx context.Context, p *request.TelegramUpdateRequestRepositoryDto) (*entity.TelegramEntity, error)
	FindByTelegramUserId(ctx context.Context, telegramUserId string) (*entity.TelegramEntity, error)
}

type ImageRepository interface {
	Add(ctx context.Context, p *request.ImageAddRequestRepositoryDto) (uint64, error)
	Delete(ctx context.Context, id uint64) (*response.ResponseDto, error)
	FindById(ctx context.Context, imageId uint64) (*response.ImageResponseRepositoryDto, error)
	FindLastByTelegramUserId(ctx context.Context, telegramUserId string) (*response.ImageResponseDto, error)
	SelectListAllByTelegramUserId(ctx context.Context, telegramUserId string) ([]*response.ImageResponseDto, error)
	SelectListPublicByTelegramUserId(ctx context.Context, telegramUserId string) ([]*response.ImageResponseDto, error)
	SelectListByTelegramUserId(ctx context.Context, telegramUserId string) ([]*response.ImageResponseDto, error)
}

type ImageStatusRepository interface {
	Add(ctx context.Context, p *request.ImageStatusAddRequestRepositoryDto) (*response.ResponseDto, error)
	FindById(ctx context.Context, id uint64) (*entity.ImageStatusEntity, error)
}

type LikeRepository interface {
	Add(ctx context.Context, p *request.LikeAddRequestRepositoryDto) (*response.ResponseDto, error)
	Update(ctx context.Context, p *request.LikeUpdateRequestRepositoryDto) (*response.ResponseDto, error)
	DeleteRelatedProfiles(ctx context.Context, id string) (*response.ResponseDto, error)
	FindById(ctx context.Context, id uint64) (*entity.LikeEntity, error)
	FindLastLike(ctx context.Context, telegramUserId string) (*entity.LikeEntity, error)
}

type BlockRepository interface {
	Add(ctx context.Context, p *request.BlockAddRequestRepositoryDto) (*response.ResponseDto, error)
	DeleteRelatedProfiles(ctx context.Context, id string) (*response.ResponseDto, error)
}

type ComplaintRepository interface {
	Add(ctx context.Context, p *request.ComplaintAddRequestRepositoryDto) (*response.ResponseDto, error)
	DeleteRelatedProfiles(ctx context.Context, id string) (*response.ResponseDto, error)
	GetCountUserComplaintsByToday(ctx context.Context, telegramUserId string) (uint64, error)
}

type StatusRepository interface {
	Add(ctx context.Context, p *request.StatusAddRequestRepositoryDto) (*response.ResponseDto, error)
	Update(ctx context.Context, p *request.StatusUpdateRequestRepositoryDto) (*entity.StatusEntity, error)
	Block(ctx context.Context, telegramUserId string) (*entity.StatusEntity, error)
	Freeze(ctx context.Context, telegramUserId string) (*entity.StatusEntity, error)
	Restore(ctx context.Context, telegramUserId string) (*entity.StatusEntity, error)
	FindByTelegramUserId(ctx context.Context, telegramUserId string) (*entity.StatusEntity, error)
	CheckProfileExists(ctx context.Context, telegramUserId string) (*response.CheckExistsDto, error)
}
