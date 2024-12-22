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
		Age:                     in.Age,
		Gender:                  in.Gender,
		SearchGender:            in.SearchGender,
		Location:                in.Location,
		Description:             in.Description,
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
		Age:                     in.Age,
		Gender:                  in.Gender,
		SearchGender:            in.SearchGender,
		Location:                in.Location,
		Description:             in.Description,
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

func (pm *ProfileControllerMapper) MapControllerResponse(
	r *response.ProfileResponseDto) *pb.ProfileResponse {
	var navigatorResponse *pb.NavigatorResponse
	if r.Navigator != nil {
		location := &pb.Point{
			Latitude:  r.Navigator.Location.Latitude,
			Longitude: r.Navigator.Location.Longitude,
		}
		navigatorResponse = &pb.NavigatorResponse{
			Location: location,
		}
	}
	images := make([]*pb.ImageResponse, 0)
	if len(r.Images) > 0 {
		for _, image := range r.Images {
			images = append(images, &pb.ImageResponse{
				Id:             image.Id,
				TelegramUserId: image.TelegramUserId,
				Name:           image.Name,
				Url:            image.Url,
			})
		}
	}
	return &pb.ProfileResponse{
		TelegramUserId: r.TelegramUserId,
		DisplayName:    r.DisplayName,
		Age:            r.Age,
		Gender:         r.Gender,
		Location:       r.Location,
		Description:    r.Description,
		Navigator:      navigatorResponse,
		Filter: &pb.FilterResponse{
			SearchGender: r.Filter.SearchGender,
			AgeFrom:      r.Filter.AgeFrom,
			AgeTo:        r.Filter.AgeTo,
			Distance:     r.Filter.Distance,
			Page:         r.Filter.Page,
			Size:         r.Filter.Size,
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
		likeUpdatedAtTimestamp := timestamppb.New(r.Like.UpdatedAt)
		likeResponse = &pb.LikeResponse{
			Id:        r.Like.Id,
			IsLiked:   r.Like.IsLiked,
			UpdatedAt: likeUpdatedAtTimestamp,
		}
	}
	images := make([]*pb.ImageResponse, 0)
	if len(r.Images) > 0 {
		for _, image := range r.Images {
			images = append(images, &pb.ImageResponse{
				Id:             image.Id,
				TelegramUserId: image.TelegramUserId,
				Name:           image.Name,
				Url:            image.Url,
			})
		}
	}
	return &pb.ProfileDetailResponse{
		TelegramUserId: r.TelegramUserId,
		DisplayName:    r.DisplayName,
		Age:            r.Age,
		Location:       r.Location,
		Description:    r.Description,
		Navigator:      navigatorResponse,
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
				IsLiked:        c.IsLiked,
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
	return &pb.ImageResponse{
		Id:             r.Id,
		TelegramUserId: r.TelegramUserId,
		Name:           r.Name,
		Url:            r.Url,
	}
}

func (pm *ProfileControllerMapper) MapControllerToFilterResponse(
	r *response.FilterResponseDto) *pb.FilterResponse {
	return &pb.FilterResponse{
		SearchGender: r.SearchGender,
		AgeFrom:      r.AgeFrom,
		AgeTo:        r.AgeTo,
		Distance:     r.Distance,
		Page:         r.Page,
		Size:         r.Size,
	}
}

func (pm *ProfileControllerMapper) MapControllerToBlockAddResponse(r *response.ResponseDto) *pb.BlockAddResponse {
	return &pb.BlockAddResponse{
		Success: r.Success,
	}
}

func (pm *ProfileControllerMapper) MapControllerToLikeAddResponse(r *response.ResponseDto) *pb.LikeAddResponse {
	return &pb.LikeAddResponse{
		Success: r.Success,
	}
}

func (pm *ProfileControllerMapper) MapControllerToLikeUpdateResponse(
	r *response.ResponseDto) *pb.LikeUpdateResponse {
	return &pb.LikeUpdateResponse{
		Success: r.Success,
	}
}

func (pm *ProfileControllerMapper) MapControllerToComplaintAddResponse(
	r *response.ResponseDto) *pb.ComplaintAddResponse {
	return &pb.ComplaintAddResponse{
		Success: r.Success,
	}
}

func (pm *ProfileControllerMapper) MapControllerToUpdateCoordinatesResponse(
	r *response.NavigatorResponseDto) *pb.NavigatorUpdateResponse {
	location := &pb.Point{
		Latitude:  r.Location.Latitude,
		Longitude: r.Location.Longitude,
	}
	return &pb.NavigatorUpdateResponse{
		Location: location,
	}
}
