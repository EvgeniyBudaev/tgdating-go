package mapper

import (
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProfileControllerMapper struct {
}

func (pm *ProfileControllerMapper) MapControllerToAddRequest(
	in *pb.ProfileAddRequest, fileList []*entity.FileMetadata) *request.ProfileAddRequestDto {
	return &request.ProfileAddRequestDto{
		SessionId:               in.SessionId,
		DisplayName:             in.DisplayName,
		Birthday:                in.Birthday.AsTime(),
		Gender:                  in.Gender,
		SearchGender:            in.SearchGender,
		Location:                in.Location,
		Description:             in.Description,
		Height:                  in.Height,
		Weight:                  in.Weight,
		TelegramUserId:          in.TelegramUserId,
		TelegramUsername:        in.TelegramUsername,
		TelegramFirstName:       in.TelegramFirstName,
		TelegramLastName:        in.TelegramLastName,
		TelegramLanguageCode:    in.TelegramLanguageCode,
		TelegramAllowsWriteToPm: in.TelegramAllowsWriteToPm,
		TelegramQueryId:         in.TelegramQueryId,
		Latitude:                in.Latitude,
		Longitude:               in.Longitude,
		AgeFrom:                 in.AgeFrom,
		AgeTo:                   in.AgeTo,
		Distance:                in.Distance,
		Page:                    in.Page,
		Size:                    in.Size,
		Files:                   fileList,
	}
}

func (pm *ProfileControllerMapper) MapControllerToAddResponse(r *response.ProfileAddResponseDto) *pb.ProfileAddResponse {
	return &pb.ProfileAddResponse{
		SessionId: r.SessionId,
	}
}

func (pm *ProfileControllerMapper) MapControllerToUpdateRequest(
	in *pb.ProfileUpdateRequest, fileList []*entity.FileMetadata) *request.ProfileUpdateRequestDto {
	return &request.ProfileUpdateRequestDto{
		SessionId:               in.SessionId,
		DisplayName:             in.DisplayName,
		Birthday:                in.Birthday.AsTime(),
		Gender:                  in.Gender,
		SearchGender:            in.SearchGender,
		Location:                in.Location,
		Description:             in.Description,
		Height:                  in.Height,
		Weight:                  in.Weight,
		TelegramUserId:          in.TelegramUserId,
		TelegramUsername:        in.TelegramUsername,
		TelegramFirstName:       in.TelegramFirstName,
		TelegramLastName:        in.TelegramLastName,
		TelegramLanguageCode:    in.TelegramLanguageCode,
		TelegramAllowsWriteToPm: in.TelegramAllowsWriteToPm,
		TelegramQueryId:         in.TelegramQueryId,
		Latitude:                in.Latitude,
		Longitude:               in.Longitude,
		AgeFrom:                 in.AgeFrom,
		AgeTo:                   in.AgeTo,
		Distance:                in.Distance,
		Page:                    in.Page,
		Size:                    in.Size,
		IsImages:                in.IsImages,
		Files:                   fileList,
	}
}

func (pm *ProfileControllerMapper) MapControllerToBySessionIdResponse(
	r *response.ProfileResponseDto) *pb.ProfileBySessionIdResponse {
	birthdayTimestamp := timestamppb.New(r.Birthday)
	createdAtTimestamp := timestamppb.New(r.CreatedAt)
	updatedAtTimestamp := timestamppb.New(r.UpdatedAt)
	lastOnlineTimestamp := timestamppb.New(r.LastOnline)
	var navigatorResponse *pb.NavigatorResponse
	if r.Navigator != nil {
		location := &pb.Point{
			Latitude:  r.Navigator.Location.Latitude,
			Longitude: r.Navigator.Location.Longitude,
		}
		navigatorResponse = &pb.NavigatorResponse{
			SessionId: r.Navigator.SessionId,
			Location:  location,
		}
	}
	images := make([]*pb.Image, 0)
	if len(r.Images) > 0 {
		for _, image := range r.Images {
			createdAtImageTimestamp := timestamppb.New(image.CreatedAt)
			updatedAtImageTimestamp := timestamppb.New(image.UpdatedAt)
			images = append(images, &pb.Image{
				Id:        image.Id,
				SessionId: image.SessionId,
				Name:      image.Name,
				Url:       image.Url,
				Size:      image.Size,
				IsDeleted: image.IsDeleted,
				IsBlocked: image.IsBlocked,
				IsPrimary: image.IsPrimary,
				IsPrivate: image.IsPrivate,
				CreatedAt: createdAtImageTimestamp,
				UpdatedAt: updatedAtImageTimestamp,
			})
		}
	}
	return &pb.ProfileBySessionIdResponse{
		SessionId:      r.SessionId,
		DisplayName:    r.DisplayName,
		Birthday:       birthdayTimestamp,
		Gender:         r.Gender,
		Location:       r.Location,
		Description:    r.Description,
		Height:         r.Height,
		Weight:         r.Weight,
		IsDeleted:      r.IsDeleted,
		IsBlocked:      r.IsBlocked,
		IsPremium:      r.IsPremium,
		IsShowDistance: r.IsShowDistance,
		IsInvisible:    r.IsInvisible,
		IsOnline:       r.IsOnline,
		CreatedAt:      createdAtTimestamp,
		UpdatedAt:      updatedAtTimestamp,
		LastOnline:     lastOnlineTimestamp,
		Navigator:      navigatorResponse,
		Filter: &pb.FilterResponse{
			SessionId:    r.Filter.SessionId,
			SearchGender: r.Filter.SearchGender,
			LookingFor:   r.Filter.LookingFor,
			AgeFrom:      r.Filter.AgeFrom,
			AgeTo:        r.Filter.AgeTo,
			Distance:     r.Filter.Distance,
			Page:         r.Filter.Page,
			Size:         r.Filter.Size,
		},
		Telegram: &pb.TelegramResponse{
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

func (pm *ProfileControllerMapper) MapControllerToDetailResponse(
	r *response.ProfileDetailResponseDto) *pb.ProfileDetailResponse {
	birthdayTimestamp := timestamppb.New(r.Birthday)
	createdAtTimestamp := timestamppb.New(r.CreatedAt)
	updatedAtTimestamp := timestamppb.New(r.UpdatedAt)
	lastOnlineTimestamp := timestamppb.New(r.LastOnline)
	var navigatorResponse *pb.NavigatorDetailResponse
	if r.Navigator != nil {
		navigatorResponse = &pb.NavigatorDetailResponse{
			Distance: r.Navigator.Distance,
		}
	}
	var blockResponse *pb.BlockResponse
	if r.Block != nil {
		blockResponse = &pb.BlockResponse{
			IsBlocked: r.Block.IsBlocked,
		}
	}
	var likeResponse *pb.LikeResponse
	if r.Like != nil {
		likeCreatedAtTimestamp := timestamppb.New(r.Like.CreatedAt)
		likeUpdatedAtTimestamp := timestamppb.New(r.Like.UpdatedAt)
		likeResponse = &pb.LikeResponse{
			SessionId:      r.Like.SessionId,
			LikedSessionId: r.Like.LikedSessionId,
			IsLiked:        r.Like.IsLiked,
			CreatedAt:      likeCreatedAtTimestamp,
			UpdatedAt:      likeUpdatedAtTimestamp,
		}
	}
	images := make([]*pb.Image, 0)
	if len(r.Images) > 0 {
		for _, image := range r.Images {
			createdAtImageTimestamp := timestamppb.New(image.CreatedAt)
			updatedAtImageTimestamp := timestamppb.New(image.UpdatedAt)
			images = append(images, &pb.Image{
				Id:        image.Id,
				SessionId: image.SessionId,
				Name:      image.Name,
				Url:       image.Url,
				Size:      image.Size,
				IsDeleted: image.IsDeleted,
				IsBlocked: image.IsBlocked,
				IsPrimary: image.IsPrimary,
				IsPrivate: image.IsPrivate,
				CreatedAt: createdAtImageTimestamp,
				UpdatedAt: updatedAtImageTimestamp,
			})
		}
	}
	return &pb.ProfileDetailResponse{
		SessionId:      r.SessionId,
		DisplayName:    r.DisplayName,
		Birthday:       birthdayTimestamp,
		Gender:         r.Gender,
		Location:       r.Location,
		Description:    r.Description,
		Height:         r.Height,
		Weight:         r.Weight,
		IsDeleted:      r.IsDeleted,
		IsBlocked:      r.IsBlocked,
		IsPremium:      r.IsPremium,
		IsShowDistance: r.IsShowDistance,
		IsInvisible:    r.IsInvisible,
		IsOnline:       r.IsOnline,
		CreatedAt:      createdAtTimestamp,
		UpdatedAt:      updatedAtTimestamp,
		LastOnline:     lastOnlineTimestamp,
		Navigator:      navigatorResponse,
		Telegram: &pb.TelegramResponse{
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

func (pm *ProfileControllerMapper) MapControllerToShortInfoResponse(
	r *response.ProfileShortInfoResponseDto) *pb.ProfileShortInfoResponse {
	return &pb.ProfileShortInfoResponse{
		SessionId: r.SessionId,
		ImageUrl:  r.ImageUrl,
		IsDeleted: r.IsDeleted,
		IsBlocked: r.IsBlocked,
	}
}
