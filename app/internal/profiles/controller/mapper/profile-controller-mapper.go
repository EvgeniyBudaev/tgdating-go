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
		DisplayName:             in.DisplayName,
		Birthday:                in.Birthday.AsTime(),
		Gender:                  in.Gender,
		SearchGender:            in.SearchGender,
		Location:                in.Location,
		Description:             in.Description,
		Height:                  in.Height,
		Weight:                  in.Weight,
		LookingFor:              in.LookingFor,
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

func (pm *ProfileControllerMapper) MapControllerToAddResponse(r *response.ResponseDto) *pb.ProfileAddResponse {
	return &pb.ProfileAddResponse{
		Success: r.Success,
	}
}

func (pm *ProfileControllerMapper) MapControllerToUpdateRequest(
	in *pb.ProfileUpdateRequest, fileList []*entity.FileMetadata) *request.ProfileUpdateRequestDto {
	return &request.ProfileUpdateRequestDto{
		DisplayName:             in.DisplayName,
		Birthday:                in.Birthday.AsTime(),
		Gender:                  in.Gender,
		SearchGender:            in.SearchGender,
		Location:                in.Location,
		Description:             in.Description,
		Height:                  in.Height,
		Weight:                  in.Weight,
		LookingFor:              in.LookingFor,
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

func (pm *ProfileControllerMapper) MapControllerToByTelegramUserIdResponse(
	r *response.ProfileResponseDto) *pb.ProfileByTelegramUserIdResponse {
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
			TelegramUserId: r.Navigator.TelegramUserId,
			Location:       location,
		}
	}
	images := make([]*pb.ImageResponse, 0)
	if len(r.Images) > 0 {
		for _, image := range r.Images {
			createdAtImageTimestamp := timestamppb.New(image.CreatedAt)
			updatedAtImageTimestamp := timestamppb.New(image.UpdatedAt)
			status := &pb.ImageStatusResponse{
				IsBlocked: image.Status.IsBlocked,
				IsPrimary: image.Status.IsPrimary,
				IsPrivate: image.Status.IsPrivate,
			}
			images = append(images, &pb.ImageResponse{
				Id:             image.Id,
				TelegramUserId: image.TelegramUserId,
				Name:           image.Name,
				Url:            image.Url,
				Size:           image.Size,
				Status:         status,
				CreatedAt:      createdAtImageTimestamp,
				UpdatedAt:      updatedAtImageTimestamp,
			})
		}
	}
	return &pb.ProfileByTelegramUserIdResponse{
		TelegramUserId: r.TelegramUserId,
		DisplayName:    r.DisplayName,
		Birthday:       birthdayTimestamp,
		Gender:         r.Gender,
		Location:       r.Location,
		Description:    r.Description,
		Height:         r.Height,
		Weight:         r.Weight,
		CreatedAt:      createdAtTimestamp,
		UpdatedAt:      updatedAtTimestamp,
		LastOnline:     lastOnlineTimestamp,
		Navigator:      navigatorResponse,
		Filter: &pb.FilterResponse{
			TelegramUserId: r.Filter.TelegramUserId,
			SearchGender:   r.Filter.SearchGender,
			LookingFor:     r.Filter.LookingFor,
			AgeFrom:        r.Filter.AgeFrom,
			AgeTo:          r.Filter.AgeTo,
			Distance:       r.Filter.Distance,
			Page:           r.Filter.Page,
			Size:           r.Filter.Size,
		},
		Telegram: &pb.TelegramResponse{
			UserId:          r.Telegram.UserId,
			Username:        r.Telegram.Username,
			FirstName:       r.Telegram.FirstName,
			LastName:        r.Telegram.LastName,
			LanguageCode:    r.Telegram.LanguageCode,
			AllowsWriteToPm: r.Telegram.AllowsWriteToPm,
			QueryId:         r.Telegram.QueryId,
		},
		Status: &pb.StatusResponse{
			IsBlocked:      r.Status.IsBlocked,
			IsFrozen:       r.Status.IsFrozen,
			IsInvisible:    r.Status.IsInvisible,
			IsOnline:       r.Status.IsOnline,
			IsPremium:      r.Status.IsPremium,
			IsShowDistance: r.Status.IsShowDistance,
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
			Id:                  r.Like.Id,
			TelegramUserId:      r.Like.TelegramUserId,
			LikedTelegramUserId: r.Like.LikedTelegramUserId,
			IsLiked:             r.Like.IsLiked,
			CreatedAt:           likeCreatedAtTimestamp,
			UpdatedAt:           likeUpdatedAtTimestamp,
		}
	}
	images := make([]*pb.ImageResponse, 0)
	if len(r.Images) > 0 {
		for _, image := range r.Images {
			createdAtImageTimestamp := timestamppb.New(image.CreatedAt)
			updatedAtImageTimestamp := timestamppb.New(image.UpdatedAt)
			status := &pb.ImageStatusResponse{
				IsBlocked: image.Status.IsBlocked,
				IsPrimary: image.Status.IsPrimary,
				IsPrivate: image.Status.IsPrivate,
			}
			images = append(images, &pb.ImageResponse{
				Id:             image.Id,
				TelegramUserId: image.TelegramUserId,
				Name:           image.Name,
				Url:            image.Url,
				Size:           image.Size,
				Status:         status,
				CreatedAt:      createdAtImageTimestamp,
				UpdatedAt:      updatedAtImageTimestamp,
			})
		}
	}
	return &pb.ProfileDetailResponse{
		TelegramUserId: r.TelegramUserId,
		DisplayName:    r.DisplayName,
		Birthday:       birthdayTimestamp,
		Gender:         r.Gender,
		Location:       r.Location,
		Description:    r.Description,
		Height:         r.Height,
		Weight:         r.Weight,
		CreatedAt:      createdAtTimestamp,
		UpdatedAt:      updatedAtTimestamp,
		LastOnline:     lastOnlineTimestamp,
		Navigator:      navigatorResponse,
		Telegram: &pb.TelegramResponse{
			UserId:          r.Telegram.UserId,
			Username:        r.Telegram.Username,
			FirstName:       r.Telegram.FirstName,
			LastName:        r.Telegram.LastName,
			LanguageCode:    r.Telegram.LanguageCode,
			AllowsWriteToPm: r.Telegram.AllowsWriteToPm,
			QueryId:         r.Telegram.QueryId,
		},
		Status: &pb.StatusResponse{
			IsBlocked:      r.Status.IsBlocked,
			IsFrozen:       r.Status.IsFrozen,
			IsInvisible:    r.Status.IsInvisible,
			IsOnline:       r.Status.IsOnline,
			IsPremium:      r.Status.IsPremium,
			IsShowDistance: r.Status.IsShowDistance,
		},
		Block:  blockResponse,
		Like:   likeResponse,
		Images: images,
	}
}

func (pm *ProfileControllerMapper) MapControllerToShortInfoResponse(
	r *response.ProfileShortInfoResponseDto) *pb.ProfileShortInfoResponse {
	return &pb.ProfileShortInfoResponse{
		TelegramUserId: r.TelegramUserId,
		IsBlocked:      r.IsBlocked,
		IsFrozen:       r.IsFrozen,
		SearchGender:   r.SearchGender,
		LookingFor:     r.LookingFor,
		AgeFrom:        r.AgeFrom,
		AgeTo:          r.AgeTo,
		Distance:       r.Distance,
		Page:           r.Page,
		Size:           r.Size,
	}
}

func (pm *ProfileControllerMapper) MapControllerToListResponse(
	r *response.ProfileListResponseDto) *pb.ProfileListResponse {
	contentList := make([]*pb.ProfileListItemResponse, 0)
	if len(r.Content) > 0 {
		for _, c := range r.Content {
			lastOnlineTimestamp := timestamppb.New(c.LastOnline)
			contentList = append(contentList, &pb.ProfileListItemResponse{
				TelegramUserId: c.TelegramUserId,
				Distance:       c.Distance,
				Url:            c.Url,
				IsOnline:       c.IsOnline,
				LastOnline:     lastOnlineTimestamp,
			})
		}
	}
	return &pb.ProfileListResponse{
		HasPrevious:   r.HasPrevious,
		HasNext:       r.HasNext,
		Page:          r.Page,
		Size:          r.Size,
		TotalEntities: r.TotalEntities,
		TotalPages:    r.TotalPages,
		Content:       contentList,
	}
}

func (pm *ProfileControllerMapper) MapControllerToImageResponse(r *response.ImageResponseDto) *pb.ImageResponse {
	createdAtImageTimestamp := timestamppb.New(r.CreatedAt)
	updatedAtImageTimestamp := timestamppb.New(r.UpdatedAt)
	status := &pb.ImageStatusResponse{
		IsBlocked: r.Status.IsBlocked,
		IsPrimary: r.Status.IsPrimary,
		IsPrivate: r.Status.IsPrivate,
	}
	return &pb.ImageResponse{
		Id:             r.Id,
		TelegramUserId: r.TelegramUserId,
		Name:           r.Name,
		Url:            r.Url,
		Size:           r.Size,
		Status:         status,
		CreatedAt:      createdAtImageTimestamp,
		UpdatedAt:      updatedAtImageTimestamp,
	}
}

func (pm *ProfileControllerMapper) MapControllerToFilterUpdateResponse(
	r *response.FilterResponseDto) *pb.FilterUpdateResponse {
	return &pb.FilterUpdateResponse{
		TelegramUserId: r.TelegramUserId,
		SearchGender:   r.SearchGender,
		LookingFor:     r.LookingFor,
		AgeFrom:        r.AgeFrom,
		AgeTo:          r.AgeTo,
		Distance:       r.Distance,
		Page:           r.Page,
		Size:           r.Size,
	}
}

func (pm *ProfileControllerMapper) MapControllerToBlockAddResponse(r *entity.BlockEntity) *pb.BlockAddResponse {
	createdAtTimestamp := timestamppb.New(r.CreatedAt)
	updatedAtTimestamp := timestamppb.New(r.UpdatedAt)
	return &pb.BlockAddResponse{
		Id:                    r.Id,
		TelegramUserId:        r.TelegramUserId,
		BlockedTelegramUserId: r.BlockedTelegramUserId,
		IsBlocked:             r.IsBlocked,
		CreatedAt:             createdAtTimestamp,
		UpdatedAt:             updatedAtTimestamp,
	}
}

func (pm *ProfileControllerMapper) MapControllerToLikeAddResponse(r *response.LikeResponseDto) *pb.LikeAddResponse {
	createdAtTimestamp := timestamppb.New(r.CreatedAt)
	updatedAtTimestamp := timestamppb.New(r.UpdatedAt)
	return &pb.LikeAddResponse{
		Id:                  r.Id,
		TelegramUserId:      r.TelegramUserId,
		LikedTelegramUserId: r.LikedTelegramUserId,
		IsLiked:             r.IsLiked,
		CreatedAt:           createdAtTimestamp,
		UpdatedAt:           updatedAtTimestamp,
	}
}

func (pm *ProfileControllerMapper) MapControllerToLikeUpdateResponse(
	r *response.LikeResponseDto) *pb.LikeUpdateResponse {
	createdAtTimestamp := timestamppb.New(r.CreatedAt)
	updatedAtTimestamp := timestamppb.New(r.UpdatedAt)
	return &pb.LikeUpdateResponse{
		Id:                  r.Id,
		TelegramUserId:      r.TelegramUserId,
		LikedTelegramUserId: r.LikedTelegramUserId,
		IsLiked:             r.IsLiked,
		CreatedAt:           createdAtTimestamp,
		UpdatedAt:           updatedAtTimestamp,
	}
}

func (pm *ProfileControllerMapper) MapControllerToComplaintAddResponse(
	r *entity.ComplaintEntity) *pb.ComplaintAddResponse {
	createdAtTimestamp := timestamppb.New(r.CreatedAt)
	updatedAtTimestamp := timestamppb.New(r.UpdatedAt)
	return &pb.ComplaintAddResponse{
		Id:                     r.Id,
		TelegramUserId:         r.TelegramUserId,
		CriminalTelegramUserId: r.CriminalTelegramUserId,
		Reason:                 r.Reason,
		CreatedAt:              createdAtTimestamp,
		UpdatedAt:              updatedAtTimestamp,
	}
}

func (pm *ProfileControllerMapper) MapControllerToUpdateCoordinatesResponse(
	r *response.NavigatorResponseDto) *pb.NavigatorUpdateResponse {
	location := &pb.Point{
		Latitude:  r.Location.Latitude,
		Longitude: r.Location.Longitude,
	}
	return &pb.NavigatorUpdateResponse{
		TelegramUserId: r.TelegramUserId,
		Location:       location,
	}
}
