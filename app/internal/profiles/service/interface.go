package service

import (
	"context"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
)

type ProfileRepository interface {
	Add(ctx context.Context, p *request.ProfileAddRequestRepositoryDto) (*response.ResponseDto, error)
	Update(ctx context.Context, p *request.ProfileUpdateRequestRepositoryDto) (*entity.ProfileEntity, error)
	Delete(ctx context.Context, p *request.ProfileDeleteRequestDto) (*response.ResponseDto, error)
	FindByTelegramUserId(ctx context.Context, telegramUserId string) (*entity.ProfileEntity, error)
	GetShortInfo(ctx context.Context, telegramUserId string) (*response.ProfileShortInfoResponseDto, error)
	SelectList(ctx context.Context,
		pr *request.ProfileGetListRequestRepositoryDto) (*response.ProfileListResponseRepositoryDto, error)
	UpdateLastOnline(ctx context.Context, p *request.ProfileUpdateLastOnlineRequestRepositoryDto) error
}

type NavigatorRepository interface {
	Add(ctx context.Context, p *request.NavigatorAddRequestRepositoryDto) (*entity.NavigatorEntity, error)
	Update(ctx context.Context, p *request.NavigatorUpdateRequestRepositoryDto) (*entity.NavigatorEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.NavigatorEntity, error)
	FindByTelegramUserId(ctx context.Context, telegramUserId string) (*entity.NavigatorEntity, error)
	FindDistance(ctx context.Context, pe *entity.NavigatorEntity,
		pve *entity.NavigatorEntity) (*response.NavigatorDistanceResponseRepositoryDto, error)
}

type FilterRepository interface {
	Add(ctx context.Context, p *request.FilterAddRequestRepositoryDto) (*entity.FilterEntity, error)
	Update(ctx context.Context, p *request.FilterUpdateRequestRepositoryDto) (*entity.FilterEntity, error)
	FindByTelegramUserId(ctx context.Context, telegramUserId string) (*entity.FilterEntity, error)
}

type TelegramRepository interface {
	Add(ctx context.Context, p *request.TelegramAddRequestRepositoryDto) (*entity.TelegramEntity, error)
	Update(ctx context.Context, p *request.TelegramUpdateRequestRepositoryDto) (*entity.TelegramEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.TelegramEntity, error)
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
	Add(ctx context.Context, p *request.LikeAddRequestRepositoryDto) (*entity.LikeEntity, error)
	Update(ctx context.Context, p *request.LikeUpdateRequestRepositoryDto) (*entity.LikeEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.LikeEntity, error)
	FindByTelegramUserId(ctx context.Context, telegramUserId string) (*entity.LikeEntity, error)
}

type BlockRepository interface {
	Add(ctx context.Context, p *request.BlockAddRequestRepositoryDto) (*entity.BlockEntity, error)
	Find(ctx context.Context, telegramUserId, blockedTelegramUserId string) (*entity.BlockEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.BlockEntity, error)
}

type ComplaintRepository interface {
	Add(ctx context.Context, p *request.ComplaintAddRequestRepositoryDto) (*entity.ComplaintEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.ComplaintEntity, error)
	GetCountUserComplaintsByToday(ctx context.Context, telegramUserId string) (uint64, error)
}

type StatusRepository interface {
	Add(ctx context.Context, p *request.StatusAddRequestRepositoryDto) (*entity.StatusEntity, error)
	Update(ctx context.Context, p *request.StatusUpdateRequestRepositoryDto) (*entity.StatusEntity, error)
	Block(ctx context.Context, telegramUserId string) (*entity.StatusEntity, error)
	Freeze(ctx context.Context, telegramUserId string) (*entity.StatusEntity, error)
	Restore(ctx context.Context, telegramUserId string) (*entity.StatusEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.StatusEntity, error)
	FindByTelegramUserId(ctx context.Context, telegramUserId string) (*entity.StatusEntity, error)
}
