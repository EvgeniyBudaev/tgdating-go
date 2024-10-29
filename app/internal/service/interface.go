package service

import (
	"context"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
)

type ProfileRepository interface {
	Add(ctx context.Context, p *request.ProfileAddRequestRepositoryDto) (*entity.ProfileEntity, error)
	Update(ctx context.Context, p *request.ProfileUpdateRequestRepositoryDto) (*entity.ProfileEntity, error)
	Delete(ctx context.Context, p *request.ProfileDeleteRequestRepositoryDto) (*entity.ProfileEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.ProfileEntity, error)
	FindBySessionId(ctx context.Context, sessionId string) (*entity.ProfileEntity, error)
	SelectListBySessionId(ctx context.Context,
		pr *request.ProfileGetListRequestRepositoryDto) (*response.ProfileListResponseRepositoryDto, error)
	UpdateLastOnline(ctx context.Context, p *request.ProfileUpdateLastOnlineRequestRepositoryDto) error
}

type NavigatorRepository interface {
	Add(ctx context.Context, p *request.NavigatorAddRequestRepositoryDto) (*entity.NavigatorEntity, error)
	Update(ctx context.Context, p *request.NavigatorUpdateRequestRepositoryDto) (*entity.NavigatorEntity, error)
	Delete(ctx context.Context, p *request.NavigatorDeleteRequestRepositoryDto) (*entity.NavigatorEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.NavigatorEntity, error)
	FindBySessionId(ctx context.Context, sessionId string) (*entity.NavigatorEntity, error)
	FindDistance(ctx context.Context, pe *entity.NavigatorEntity,
		pve *entity.NavigatorEntity) (*response.NavigatorDistanceResponseRepositoryDto, error)
}

type FilterRepository interface {
	Add(ctx context.Context, p *request.FilterAddRequestRepositoryDto) (*entity.FilterEntity, error)
	Update(ctx context.Context, p *request.FilterUpdateRequestRepositoryDto) (*entity.FilterEntity, error)
	Delete(ctx context.Context, p *request.FilterDeleteRequestRepositoryDto) (*entity.FilterEntity, error)
	FindBySessionId(ctx context.Context, sessionId string) (*entity.FilterEntity, error)
}

type TelegramRepository interface {
	Add(ctx context.Context, p *request.TelegramAddRequestRepositoryDto) (*entity.TelegramEntity, error)
	Update(ctx context.Context, p *request.TelegramUpdateRequestRepositoryDto) (*entity.TelegramEntity, error)
	Delete(ctx context.Context, p *request.TelegramDeleteRequestRepositoryDto) (*entity.TelegramEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.TelegramEntity, error)
	FindBySessionId(ctx context.Context, sessionId string) (*entity.TelegramEntity, error)
}

type ImageRepository interface {
	Add(ctx context.Context, p *request.ImageAddRequestRepositoryDto) (*entity.ImageEntity, error)
	Update(ctx context.Context, p *request.ImageUpdateRequestRepositoryDto) (*entity.ImageEntity, error)
	Delete(ctx context.Context, p *request.ImageDeleteRequestRepositoryDto) (*entity.ImageEntity, error)
	FindById(ctx context.Context, imageId uint64) (*entity.ImageEntity, error)
	FindLastBySessionId(ctx context.Context, sessionId string) (*entity.ImageEntity, error)
	SelectListPublicBySessionId(ctx context.Context, sessionId string) ([]*entity.ImageEntity, error)
	SelectListBySessionId(ctx context.Context, sessionId string) ([]*entity.ImageEntity, error)
}

type LikeRepository interface {
	Add(ctx context.Context, p *request.LikeAddRequestRepositoryDto) (*entity.LikeEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.LikeEntity, error)
	FindBySessionId(ctx context.Context, sessionId string) (*entity.LikeEntity, error)
}

type BlockRepository interface {
	Add(ctx context.Context, p *request.BlockAddRequestRepositoryDto) (*entity.BlockEntity, error)
	Find(ctx context.Context, sessionId, blockedUserSessionId string) (*entity.BlockEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.BlockEntity, error)
}

type ComplaintRepository interface {
	Add(ctx context.Context, p *request.ComplaintAddRequestRepositoryDto) (*entity.ComplaintEntity, error)
	FindById(ctx context.Context, id uint64) (*entity.ComplaintEntity, error)
}
