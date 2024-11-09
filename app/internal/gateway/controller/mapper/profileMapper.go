package mapper

import (
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProfileMapper struct {
}

func (pm *ProfileMapper) MapToAddRequest(
	r *request.ProfileAddRequestDto, fileList []*pb.FileMetadata) *pb.ProfileAddRequest {
	newTimestampBirthday := timestamppb.New(r.Birthday)
	return &pb.ProfileAddRequest{
		SessionId:               r.SessionId,
		DisplayName:             r.DisplayName,
		Birthday:                newTimestampBirthday,
		Gender:                  r.Gender,
		SearchGender:            r.SearchGender,
		Location:                r.Location,
		Description:             r.Description,
		Height:                  r.Height,
		Weight:                  r.Weight,
		TelegramUserId:          r.TelegramUserId,
		TelegramUsername:        r.TelegramUsername,
		TelegramFirstName:       r.TelegramFirstName,
		TelegramLastName:        r.TelegramLastName,
		TelegramLanguageCode:    r.TelegramLanguageCode,
		TelegramAllowsWriteToPm: r.TelegramAllowsWriteToPm,
		TelegramQueryId:         r.TelegramQueryId,
		Latitude:                r.Latitude,
		Longitude:               r.Longitude,
		AgeFrom:                 r.AgeFrom,
		AgeTo:                   r.AgeTo,
		Distance:                r.Distance,
		Page:                    r.Page,
		Size:                    r.Size,
		Files:                   fileList,
	}
}

func (pm *ProfileMapper) MapToUpdateRequest(
	r *request.ProfileUpdateRequestDto, fileList []*pb.FileMetadata) *pb.ProfileUpdateRequest {
	newTimestampBirthday := timestamppb.New(r.Birthday)
	return &pb.ProfileUpdateRequest{
		SessionId:               r.SessionId,
		DisplayName:             r.DisplayName,
		Birthday:                newTimestampBirthday,
		Gender:                  r.Gender,
		SearchGender:            r.SearchGender,
		Location:                r.Location,
		Description:             r.Description,
		Height:                  r.Height,
		Weight:                  r.Weight,
		TelegramUserId:          r.TelegramUserId,
		TelegramUsername:        r.TelegramUsername,
		TelegramFirstName:       r.TelegramFirstName,
		TelegramLastName:        r.TelegramLastName,
		TelegramLanguageCode:    r.TelegramLanguageCode,
		TelegramAllowsWriteToPm: r.TelegramAllowsWriteToPm,
		TelegramQueryId:         r.TelegramQueryId,
		Latitude:                r.Latitude,
		Longitude:               r.Longitude,
		AgeFrom:                 r.AgeFrom,
		AgeTo:                   r.AgeTo,
		Distance:                r.Distance,
		Page:                    r.Page,
		Size:                    r.Size,
		IsImages:                r.IsImages,
		Files:                   fileList,
	}
}

func (pm *ProfileMapper) MapToGetBySessionIdRequest(
	r *request.ProfileGetBySessionIdRequestDto, sessionId string) *pb.ProfileGetBySessionIdRequest {
	return &pb.ProfileGetBySessionIdRequest{
		SessionId: sessionId,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
	}
}

func (pm *ProfileMapper) MapToBySessionIdResponse(r *pb.ProfileBySessionIdResponse) *response.ProfileResponseDto {
	var heightResponse *float64
	if r.Height != nil {
		heightResponse = r.Height
	}
	var weightResponse *float64
	if r.Weight != nil {
		weightResponse = r.Weight
	}
	var navigatorResponse *response.NavigatorResponseDto
	if r.Navigator != nil {
		location := &entity.PointEntity{
			Latitude:  r.Navigator.Location.Latitude,
			Longitude: r.Navigator.Location.Longitude,
		}
		navigatorResponse = &response.NavigatorResponseDto{
			SessionId: r.Navigator.SessionId,
			Location:  location,
		}
	}
	images := make([]*entity.ImageEntity, 0)
	if len(r.Images) > 0 {
		for _, image := range r.Images {
			images = append(images, &entity.ImageEntity{
				Id:        image.Id,
				SessionId: image.SessionId,
				Name:      image.Name,
				Url:       image.Url,
				Size:      image.Size,
				IsDeleted: image.IsDeleted,
				IsBlocked: image.IsBlocked,
				IsPrimary: image.IsPrimary,
				IsPrivate: image.IsPrivate,
				CreatedAt: image.CreatedAt.AsTime(),
				UpdatedAt: image.UpdatedAt.AsTime(),
			})
		}
	}
	return &response.ProfileResponseDto{
		SessionId:      r.SessionId,
		DisplayName:    r.DisplayName,
		Birthday:       r.Birthday.AsTime(),
		Gender:         r.Gender,
		Location:       r.Location,
		Description:    r.Description,
		Height:         heightResponse,
		Weight:         weightResponse,
		IsDeleted:      r.IsDeleted,
		IsBlocked:      r.IsBlocked,
		IsPremium:      r.IsPremium,
		IsShowDistance: r.IsShowDistance,
		IsInvisible:    r.IsInvisible,
		IsOnline:       r.IsOnline,
		CreatedAt:      r.CreatedAt.AsTime(),
		UpdatedAt:      r.UpdatedAt.AsTime(),
		LastOnline:     r.LastOnline.AsTime(),
		Navigator:      navigatorResponse,
		Filter: &response.FilterResponseDto{
			SessionId:    r.Filter.SessionId,
			SearchGender: r.Filter.SearchGender,
			LookingFor:   r.Filter.LookingFor,
			AgeFrom:      r.Filter.AgeFrom,
			AgeTo:        r.Filter.AgeTo,
			Distance:     r.Filter.Distance,
			Page:         r.Filter.Page,
			Size:         r.Filter.Size,
		},
		Telegram: &response.TelegramResponseDto{
			SessionId:       r.Telegram.SessionId,
			UserId:          r.Telegram.UserId,
			Username:        r.Telegram.Username,
			FirstName:       r.Telegram.FirstName,
			LastName:        r.Telegram.LastName,
			LanguageCode:    r.Telegram.LanguageCode,
			AllowsWriteToPm: r.Telegram.AllowsWriteToPm,
			QueryId:         r.Telegram.QueryId,
		},
		Images: images,
	}
}

func (pm *ProfileMapper) MapToGetDetailRequest(
	r *request.ProfileGetDetailRequestDto, viewedSessionId string) *pb.ProfileGetDetailRequest {
	return &pb.ProfileGetDetailRequest{
		SessionId:       r.SessionId,
		Latitude:        r.Latitude,
		Longitude:       r.Longitude,
		ViewedSessionId: viewedSessionId,
	}
}

func (pm *ProfileMapper) MapToDetailResponse(r *pb.ProfileDetailResponse) *response.ProfileDetailResponseDto {
	var heightResponse *float64
	if r.Height != nil {
		heightResponse = r.Height
	}
	var weightResponse *float64
	if r.Weight != nil {
		weightResponse = r.Weight
	}
	var navigatorResponse *response.NavigatorDetailResponseDto
	if r.Navigator != nil {
		navigatorResponse = &response.NavigatorDetailResponseDto{
			Distance: r.Navigator.Distance,
		}
	}
	var blockResponse *response.BlockResponseDto
	if r.Block != nil {
		blockResponse = &response.BlockResponseDto{
			IsBlocked: r.Block.IsBlocked,
		}
	}
	var likeResponse *response.LikeResponseDto
	if r.Like != nil {
		likeResponse = &response.LikeResponseDto{
			SessionId:      r.Like.SessionId,
			LikedSessionId: r.Like.LikedSessionId,
			IsLiked:        r.Like.IsLiked,
			CreatedAt:      r.Like.CreatedAt.AsTime(),
			UpdatedAt:      r.Like.UpdatedAt.AsTime(),
		}
	}
	images := make([]*entity.ImageEntity, 0)
	if len(r.Images) > 0 {
		for _, image := range r.Images {
			images = append(images, &entity.ImageEntity{
				Id:        image.Id,
				SessionId: image.SessionId,
				Name:      image.Name,
				Url:       image.Url,
				Size:      image.Size,
				IsDeleted: image.IsDeleted,
				IsBlocked: image.IsBlocked,
				IsPrimary: image.IsPrimary,
				IsPrivate: image.IsPrivate,
				CreatedAt: image.CreatedAt.AsTime(),
				UpdatedAt: image.UpdatedAt.AsTime(),
			})
		}
	}
	return &response.ProfileDetailResponseDto{
		SessionId:      r.SessionId,
		DisplayName:    r.DisplayName,
		Birthday:       r.Birthday.AsTime(),
		Gender:         r.Gender,
		Location:       r.Location,
		Description:    r.Description,
		Height:         heightResponse,
		Weight:         weightResponse,
		IsDeleted:      r.IsDeleted,
		IsBlocked:      r.IsBlocked,
		IsPremium:      r.IsPremium,
		IsShowDistance: r.IsShowDistance,
		IsInvisible:    r.IsInvisible,
		IsOnline:       r.IsOnline,
		CreatedAt:      r.CreatedAt.AsTime(),
		UpdatedAt:      r.UpdatedAt.AsTime(),
		LastOnline:     r.LastOnline.AsTime(),
		Navigator:      navigatorResponse,
		Telegram: &response.TelegramResponseDto{
			SessionId:       r.Telegram.SessionId,
			UserId:          r.Telegram.UserId,
			Username:        r.Telegram.Username,
			FirstName:       r.Telegram.FirstName,
			LastName:        r.Telegram.LastName,
			LanguageCode:    r.Telegram.LanguageCode,
			AllowsWriteToPm: r.Telegram.AllowsWriteToPm,
			QueryId:         r.Telegram.QueryId,
		},
		Block:  blockResponse,
		Like:   likeResponse,
		Images: images,
	}
}
