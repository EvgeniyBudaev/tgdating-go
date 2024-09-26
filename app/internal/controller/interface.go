package controller

import (
	"context"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type ProfileService interface {
	AddProfile(ctx context.Context, ctf *fiber.Ctx,
		pr *request.ProfileAddRequestDto) (*response.ProfileAddResponseDto, error)
	UpdateProfile(ctx context.Context, ctf *fiber.Ctx,
		pr *request.ProfileUpdateRequestDto) (*response.ProfileResponseDto, error)
	DeleteProfile(ctx context.Context, pr *request.ProfileDeleteRequestDto) (*response.ResponseDto, error)
	GetProfileBySessionId(ctx context.Context, sessionId string,
		pr *request.ProfileGetBySessionIdRequestDto) (*response.ProfileResponseDto, error)
	GetProfileDetail(ctx context.Context, sessionId string,
		pr *request.ProfileGetDetailRequestDto) (*response.ProfileDetailResponseDto, error)
	GetProfileShortInfo(ctx context.Context, sessionId string,
		pr *request.ProfileGetShortInfoRequestDto) (*response.ProfileShortInfoResponseDto, error)
	GetProfileList(ctx context.Context, pr *request.ProfileGetListRequestDto) (*response.ProfileListResponseDto, error)
	DeleteImage(ctx context.Context, id uint64) (*response.ResponseDto, error)
	GetFilterBySessionId(
		ctx context.Context, sessionId string, fr *request.FilterGetRequestDto) (*response.FilterResponseDto, error)
	UpdateFilter(ctx context.Context, fr *request.FilterUpdateRequestDto) (*response.FilterResponseDto, error)
	AddBlock(ctx context.Context, pr *request.BlockRequestDto) (*entity.BlockEntity, error)
	AddLike(ctx context.Context, pr *request.LikeAddRequestDto) (*response.LikeResponseDto, error)
	AddComplaint(ctx context.Context, pr *request.ComplaintAddRequestDto) (*entity.ComplaintEntity, error)
}
