package controller

import (
	"context"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
)

type ProfileService interface {
	AddProfile(ctx context.Context, pr *request.ProfileAddRequestDto) (*response.ResponseDto, error)
	UpdateProfile(ctx context.Context, pr *request.ProfileUpdateRequestDto) (*response.ProfileResponseDto, error)
	FreezeProfile(ctx context.Context, pr *request.ProfileFreezeRequestDto) (*response.ResponseDto, error)
	RestoreProfile(ctx context.Context, pr *request.ProfileRestoreRequestDto) (*response.ResponseDto, error)
	DeleteProfile(ctx context.Context, pr *request.ProfileDeleteRequestDto) (*response.ResponseDto, error)
	GetProfile(ctx context.Context, telegramUserId string,
		pr *request.ProfileGetRequestDto) (*response.ProfileResponseDto, error)
	GetProfileDetail(ctx context.Context, telegramUserId string,
		pr *request.ProfileGetDetailRequestDto) (*response.ProfileDetailResponseDto, error)
	GetProfileShortInfo(ctx context.Context, telegramUserId string) (*response.ProfileShortInfoResponseDto, error)
	GetProfileList(ctx context.Context, pr *request.ProfileGetListRequestDto) (*response.ProfileListResponseDto, error)
	CheckProfileExists(ctx context.Context, telegramUserId string) error
	GetImageByTelegramUserId(ctx context.Context, telegramUserId, fileName string) ([]byte, error)
	GetImageLastByTelegramUserId(ctx context.Context, telegramUserId string) (*response.ImageResponseDto, error)
	GetImageById(ctx context.Context, imageId uint64) (*response.ImageResponseDto, error)
	DeleteImage(ctx context.Context, id uint64) (*response.ResponseDto, error)
	GetFilter(ctx context.Context, telegramUserId string) (*response.FilterResponseDto, error)
	UpdateFilter(ctx context.Context, fr *request.FilterUpdateRequestDto) (*response.FilterResponseDto, error)
	GetTelegram(ctx context.Context, telegramUserId string) (*response.TelegramResponseDto, error)
	AddBlock(ctx context.Context, pr *request.BlockAddRequestDto) (*response.ResponseDto, error)
	GetBlockedList(ctx context.Context, telegramUserId string) (*response.BlockedListResponseDto, error)
	Unblock(ctx context.Context, p *request.UnblockRequestDto) (*response.ResponseDto, error)
	AddLike(ctx context.Context, pr *request.LikeAddRequestDto, locale string) (*response.ResponseDto, error)
	UpdateLike(ctx context.Context, pr *request.LikeUpdateRequestDto) (*response.ResponseDto, error)
	GetLastLike(
		ctx context.Context, telegramUserId string) (*entity.LikeEntity, error)
	AddComplaint(ctx context.Context, pr *request.ComplaintAddRequestDto) (*response.ResponseDto, error)
	UpdateCoordinates(
		ctx context.Context, pr *request.NavigatorUpdateRequestDto) (*response.ResponseDto, error)
	AddPayment(ctx context.Context, pr *request.PaymentAddRequestDto) (*response.ResponseDto, error)
	GetPaymentLastByTelegramUserId(ctx context.Context, telegramUserId string) (*entity.PaymentEntity, error)
	CheckPremium(ctx context.Context, telegramUserId string) (*response.PremiumResponseDto, error)
	UpdateSettings(ctx context.Context, pr *request.ProfileUpdateSettingsRequestDto) (*response.ResponseDto, error)
}
