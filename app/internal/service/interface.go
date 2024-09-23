package service

import (
	"context"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
)

type ProfileRepository interface {
	AddProfile(ctx context.Context, p *request.ProfileAddRequestRepositoryDto) (*entity.ProfileEntity, error)
	UpdateProfile(ctx context.Context, p *request.ProfileUpdateRequestRepositoryDto) (*entity.ProfileEntity, error)
	DeleteProfile(ctx context.Context, p *request.ProfileDeleteRequestRepositoryDto) (*entity.ProfileEntity, error)
	FindProfileById(ctx context.Context, id uint64) (*entity.ProfileEntity, error)
	FindProfileBySessionId(ctx context.Context, sessionId string) (*entity.ProfileEntity, error)
	SelectProfileListBySessionId(ctx context.Context,
		pr *request.ProfileGetListRequestRepositoryDto) (*response.ProfileListResponseRepositoryDto, error)
	UpdateLastOnline(ctx context.Context, p *request.ProfileUpdateLastOnlineRequestRepositoryDto) error
}

type NavigatorRepository interface {
	AddNavigator(ctx context.Context, p *request.NavigatorAddRequestRepositoryDto) (*entity.NavigatorEntity, error)
	UpdateNavigator(ctx context.Context, p *request.NavigatorUpdateRequestDto) (*entity.NavigatorEntity, error)
	DeleteNavigator(ctx context.Context, p *request.NavigatorDeleteRequestDto) (*entity.NavigatorEntity, error)
	FindNavigatorById(ctx context.Context, id uint64) (*entity.NavigatorEntity, error)
	FindNavigatorBySessionId(ctx context.Context, sessionId string) (*entity.NavigatorEntity, error)
}

type FilterRepository interface {
	AddFilter(ctx context.Context, p *request.FilterAddRequestRepositoryDto) (*entity.FilterEntity, error)
	UpdateFilter(ctx context.Context, p *request.FilterUpdateRequestRepositoryDto) (*entity.FilterEntity, error)
	DeleteFilter(ctx context.Context, p *request.FilterDeleteRequestRepositoryDto) (*entity.FilterEntity, error)
	FindFilterBySessionId(ctx context.Context, sessionId string) (*entity.FilterEntity, error)
}

type TelegramRepository interface {
	AddTelegram(ctx context.Context, p *request.TelegramAddRequestRepositoryDto) (*entity.TelegramEntity, error)
	UpdateTelegram(ctx context.Context, p *request.TelegramUpdateRequestRepositoryDto) (*entity.TelegramEntity, error)
	DeleteTelegram(ctx context.Context, p *request.TelegramDeleteRequestRepositoryDto) (*entity.TelegramEntity, error)
	FindTelegramById(ctx context.Context, id uint64) (*entity.TelegramEntity, error)
	FindTelegramBySessionId(ctx context.Context, sessionId string) (*entity.TelegramEntity, error)
}

type ImageRepository interface {
	AddImage(ctx context.Context, p *request.ImageAddRequestRepositoryDto) (*entity.ImageEntity, error)
	UpdateImage(ctx context.Context, p *request.ImageUpdateRequestRepositoryDto) (*entity.ImageEntity, error)
	DeleteImage(ctx context.Context, p *request.ImageDeleteRequestRepositoryDto) (*entity.ImageEntity, error)
	FindImageById(ctx context.Context, imageId uint64) (*entity.ImageEntity, error)
	FindLastImageBySessionId(ctx context.Context, sessionId string) (*entity.ImageEntity, error)
	SelectImageListPublicBySessionId(ctx context.Context, sessionId string) ([]*entity.ImageEntity, error)
	SelectImageListBySessionId(ctx context.Context, sessionId string) ([]*entity.ImageEntity, error)
}

type LikeRepository interface {
	AddLike(ctx context.Context, p *request.LikeAddRequestRepositoryDto) (*entity.LikeEntity, error)
	FindLikeById(ctx context.Context, id uint64) (*entity.LikeEntity, error)
}

type BlockRepository interface {
	AddBlock(ctx context.Context, p *request.BlockAddRequestRepositoryDto) (*entity.BlockEntity, error)
	FindBlockById(ctx context.Context, id uint64) (*entity.BlockEntity, error)
}

type ComplaintRepository interface {
	AddComplaint(ctx context.Context, p *request.ComplaintAddRequestRepositoryDto) (*entity.ComplaintEntity, error)
	FindComplaintById(ctx context.Context, id uint64) (*entity.ComplaintEntity, error)
}
