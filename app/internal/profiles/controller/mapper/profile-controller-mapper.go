package mapper

import (
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/shared/enum"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProfileControllerMapper struct {
}

func (pm *ProfileControllerMapper) MapControllerToAddRequest(
	in *pb.ProfileAddRequest, fileList []*entity.FileMetadata) *request.ProfileAddRequestDto {
	return &request.ProfileAddRequestDto{
		DisplayName:             in.DisplayName,
		Age:                     in.Age,
		Gender:                  enum.Gender(in.Gender),
		SearchGender:            enum.SearchGender(in.SearchGender),
		Description:             in.Description,
		TelegramUserId:          in.TelegramUserId,
		TelegramUsername:        in.TelegramUsername,
		TelegramFirstName:       in.TelegramFirstName,
		TelegramLastName:        in.TelegramLastName,
		TelegramLanguageCode:    in.TelegramLanguageCode,
		TelegramAllowsWriteToPm: in.TelegramAllowsWriteToPm,
		TelegramQueryId:         in.TelegramQueryId,
		CountryCode:             in.CountryCode,
		CountryName:             in.CountryName,
		City:                    in.City,
		Latitude:                in.Latitude,
		Longitude:               in.Longitude,
		AgeFrom:                 in.AgeFrom,
		AgeTo:                   in.AgeTo,
		Distance:                in.Distance,
		Page:                    in.Page,
		Size:                    in.Size,
		IsLeftHand:              in.IsLeftHand,
		Measurement:             enum.Measurement(in.Measurement),
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
		Gender:                  enum.Gender(in.Gender),
		SearchGender:            enum.SearchGender(in.SearchGender),
		Description:             in.Description,
		TelegramUserId:          in.TelegramUserId,
		TelegramUsername:        in.TelegramUsername,
		TelegramFirstName:       in.TelegramFirstName,
		TelegramLastName:        in.TelegramLastName,
		TelegramLanguageCode:    in.TelegramLanguageCode,
		TelegramAllowsWriteToPm: in.TelegramAllowsWriteToPm,
		TelegramQueryId:         in.TelegramQueryId,
		CountryCode:             in.CountryCode,
		CountryName:             in.CountryName,
		City:                    in.City,
		Latitude:                in.Latitude,
		Longitude:               in.Longitude,
		AgeFrom:                 in.AgeFrom,
		AgeTo:                   in.AgeTo,
		Distance:                in.Distance,
		Page:                    in.Page,
		Size:                    in.Size,
		IsLiked:                 in.IsLiked,
		IsOnline:                in.IsOnline,
		IsImages:                in.IsImages,
		Measurement:             enum.Measurement(in.Measurement),
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
		Description:    r.Description,
		Navigator:      navigatorResponse,
		Filter: &pb.FilterResponse{
			SearchGender: r.Filter.SearchGender,
			AgeFrom:      r.Filter.AgeFrom,
			AgeTo:        r.Filter.AgeTo,
			Distance:     r.Filter.Distance,
			Page:         r.Filter.Page,
			Size:         r.Filter.Size,
			IsLiked:      r.Filter.IsLiked,
			IsOnline:     r.Filter.IsOnline,
		},
		Status: &pb.StatusResponse{
			IsBlocked:        r.Status.IsBlocked,
			IsFrozen:         r.Status.IsFrozen,
			IsHiddenAge:      r.Status.IsHiddenAge,
			IsHiddenDistance: r.Status.IsHiddenDistance,
			IsInvisible:      r.Status.IsInvisible,
			IsLeftHand:       r.Status.IsLeftHand,
			IsPremium:        r.Status.IsPremium,
		},
		Settings: &pb.SettingsResponse{
			Measurement: string(r.Settings.Measurement),
		},
		Images: images,
	}
}

func (pm *ProfileControllerMapper) MapControllerToDetailResponse(
	r *response.ProfileDetailResponseDto) *pb.ProfileDetailResponse {
	lastOnlineTimestamp := timestamppb.New(r.LastOnline)
	var navigatorResponse *pb.NavigatorDetailResponse
	if r.Navigator != nil {
		navigatorResponse = &pb.NavigatorDetailResponse{
			CountryName: r.Navigator.CountryName,
			City:        r.Navigator.City,
			Distance:    r.Navigator.Distance,
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
		Gender:         r.Gender,
		Description:    r.Description,
		LastOnline:     lastOnlineTimestamp,
		Navigator:      navigatorResponse,
		Status: &pb.StatusResponse{
			IsBlocked:        r.Status.IsBlocked,
			IsFrozen:         r.Status.IsFrozen,
			IsHiddenAge:      r.Status.IsHiddenAge,
			IsHiddenDistance: r.Status.IsHiddenDistance,
			IsInvisible:      r.Status.IsInvisible,
			IsLeftHand:       r.Status.IsLeftHand,
			IsPremium:        r.Status.IsPremium,
		},
		Settings: &pb.SettingsResponse{
			Measurement: string(r.Settings.Measurement),
		},
		Block:  blockResponse,
		Like:   likeResponse,
		Images: images,
	}
}

func (pm *ProfileControllerMapper) MapControllerToShortInfoResponse(
	r *response.ProfileShortInfoResponseDto) *pb.ProfileShortInfoResponse {
	availableUntilTimestamp := timestamppb.New(r.AvailableUntil)
	return &pb.ProfileShortInfoResponse{
		TelegramUserId: r.TelegramUserId,
		IsBlocked:      r.IsBlocked,
		IsFrozen:       r.IsFrozen,
		IsPremium:      r.IsPremium,
		AvailableUntil: availableUntilTimestamp,
		LanguageCode:   r.LanguageCode,
		Measurement:    string(r.Measurement),
		Filter: &pb.FilterResponse{
			SearchGender: r.Filter.SearchGender,
			AgeFrom:      r.Filter.AgeFrom,
			AgeTo:        r.Filter.AgeTo,
			Distance:     r.Filter.Distance,
			Page:         r.Filter.Page,
			Size:         r.Filter.Size,
			IsLiked:      r.Filter.IsLiked,
			IsOnline:     r.Filter.IsOnline,
		},
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

func (pm *ProfileControllerMapper) MapControllerToCheckProfileExistsResponse() *pb.CheckProfileExistsResponse {
	return &pb.CheckProfileExistsResponse{
		IsExists: true,
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
		IsLiked:      r.IsLiked,
		IsOnline:     r.IsOnline,
	}
}

func (pm *ProfileControllerMapper) MapControllerToTelegramResponse(
	r *response.TelegramResponseDto) *pb.TelegramResponse {
	return &pb.TelegramResponse{
		UserId:          r.UserId,
		Username:        r.UserName,
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		LanguageCode:    r.LanguageCode,
		AllowsWriteToPm: r.AllowsWriteToPm,
		QueryId:         r.QueryId,
	}
}

func (pm *ProfileControllerMapper) MapControllerToBlockAddResponse(r *response.ResponseDto) *pb.BlockAddResponse {
	return &pb.BlockAddResponse{
		Success: r.Success,
	}
}

func (pm *ProfileControllerMapper) MapControllerToGetBlockedListResponse(r *response.BlockedListResponseDto) *pb.GetBlockedListResponse {
	content := make([]*pb.BlockedListItemResponse, 0)
	if len(r.Content) > 0 {
		for _, c := range r.Content {
			content = append(content, &pb.BlockedListItemResponse{
				BlockedTelegramUserId: c.BlockedTelegramUserId,
				Url:                   c.Url,
			})
		}
	}
	return &pb.GetBlockedListResponse{
		Content: content,
	}
}

func (pm *ProfileControllerMapper) MapControllerToUnblockRequest(r *pb.UnblockRequest) *request.UnblockRequestDto {
	return &request.UnblockRequestDto{
		TelegramUserId:        r.TelegramUserId,
		BlockedTelegramUserId: r.BlockedTelegramUserId,
	}
}
func (pm *ProfileControllerMapper) MapControllerToUnblockResponse(r *response.ResponseDto) *pb.UnblockResponse {
	return &pb.UnblockResponse{
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

func (pm *ProfileControllerMapper) MapControllerToLikeGetLastResponse(
	r *entity.LikeEntity) *pb.LikeGetLastResponse {
	var like *pb.LikeEntity
	if r != nil {
		createdAtTimestamp := timestamppb.New(r.CreatedAt)
		updatedAtTimestamp := timestamppb.New(r.UpdatedAt)
		like = &pb.LikeEntity{
			Id:                  r.Id,
			TelegramUserId:      r.TelegramUserId,
			LikedTelegramUserId: r.LikedTelegramUserId,
			IsLiked:             r.IsLiked,
			CreatedAt:           createdAtTimestamp,
			UpdatedAt:           updatedAtTimestamp,
		}
	}
	return &pb.LikeGetLastResponse{
		Like: like,
	}
}

func (pm *ProfileControllerMapper) MapControllerToComplaintAddResponse(
	r *response.ResponseDto) *pb.ComplaintAddResponse {
	return &pb.ComplaintAddResponse{
		Success: r.Success,
	}
}

func (pm *ProfileControllerMapper) MapControllerToPaymentAddResponse(
	r *response.ResponseDto) *pb.PaymentAddResponse {
	return &pb.PaymentAddResponse{
		Success: r.Success,
	}
}

func (pm *ProfileControllerMapper) MapControllerToCheckPremiumResponse(
	r *response.PremiumResponseDto) *pb.CheckPremiumResponse {
	availableUntilTimestamp := timestamppb.New(r.AvailableUntil)
	return &pb.CheckPremiumResponse{
		IsPremium:      r.IsPremium,
		AvailableUntil: availableUntilTimestamp,
	}
}

func (pm *ProfileControllerMapper) MapControllerToUpdateSettingsRequest(
	in *pb.UpdateSettingsRequest) *request.ProfileUpdateSettingsRequestDto {
	return &request.ProfileUpdateSettingsRequestDto{
		TelegramUserId: in.TelegramUserId,
		IsHiddenAge:    in.IsHiddenAge,
		Measurement:    enum.Measurement(in.Measurement),
	}
}

func (pm *ProfileControllerMapper) MapControllerToUpdateSettingsResponse(
	r *response.ResponseDto) *pb.UpdateSettingsResponse {
	return &pb.UpdateSettingsResponse{
		Success: r.Success,
	}
}

func (pm *ProfileControllerMapper) MapControllerToUpdateCoordinatesRequest(
	in *pb.NavigatorUpdateRequest) *request.NavigatorUpdateRequestDto {
	return &request.NavigatorUpdateRequestDto{
		TelegramUserId: in.TelegramUserId,
		CountryCode:    in.CountryCode,
		CountryName:    in.CountryName,
		City:           in.City,
		Latitude:       in.Latitude,
		Longitude:      in.Longitude,
	}
}

func (pm *ProfileControllerMapper) MapControllerToUpdateCoordinatesResponse(
	r *response.ResponseDto) *pb.NavigatorUpdateResponse {
	return &pb.NavigatorUpdateResponse{
		Success: r.Success,
	}
}
