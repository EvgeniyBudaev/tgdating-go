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
		DisplayName:             r.DisplayName,
		Birthday:                newTimestampBirthday,
		Gender:                  string(r.Gender),
		SearchGender:            string(r.SearchGender),
		Location:                r.Location,
		Description:             r.Description,
		Height:                  r.Height,
		Weight:                  r.Weight,
		LookingFor:              string(r.LookingFor),
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
		DisplayName:             r.DisplayName,
		Birthday:                newTimestampBirthday,
		Gender:                  string(r.Gender),
		SearchGender:            string(r.SearchGender),
		Location:                r.Location,
		Description:             r.Description,
		Height:                  r.Height,
		Weight:                  r.Weight,
		LookingFor:              string(r.LookingFor),
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

func (pm *ProfileMapper) MapToFreezeRequest(r *request.ProfileFreezeRequestDto) *pb.ProfileFreezeRequest {
	return &pb.ProfileFreezeRequest{
		TelegramUserId: r.TelegramUserId,
	}
}

func (pm *ProfileMapper) MapToRestoreRequest(r *request.ProfileRestoreRequestDto) *pb.ProfileRestoreRequest {
	return &pb.ProfileRestoreRequest{
		TelegramUserId: r.TelegramUserId,
	}
}

func (pm *ProfileMapper) MapToDeleteRequest(r *request.ProfileDeleteRequestDto) *pb.ProfileDeleteRequest {
	return &pb.ProfileDeleteRequest{
		TelegramUserId: r.TelegramUserId,
	}
}

func (pm *ProfileMapper) MapToGetByTelegramUserIdRequest(
	r *request.ProfileGetByTelegramUserIdRequestDto, telegramUserId string) *pb.ProfileGetByTelegramUserIdRequest {
	return &pb.ProfileGetByTelegramUserIdRequest{
		TelegramUserId: telegramUserId,
		Latitude:       r.Latitude,
		Longitude:      r.Longitude,
	}
}

func (pm *ProfileMapper) MapToByTelegramUserIdResponse(
	r *pb.ProfileByTelegramUserIdResponse) *response.ProfileResponseDto {
	var navigatorResponse *response.NavigatorResponseDto
	if r.Navigator != nil {
		location := &entity.PointEntity{
			Latitude:  r.Navigator.Location.Latitude,
			Longitude: r.Navigator.Location.Longitude,
		}
		navigatorResponse = &response.NavigatorResponseDto{
			TelegramUserId: r.Navigator.TelegramUserId,
			Location:       location,
		}
	}
	images := make([]*entity.ImageEntity, 0)
	if len(r.Images) > 0 {
		for _, image := range r.Images {
			images = append(images, &entity.ImageEntity{
				Id:             image.Id,
				TelegramUserId: image.TelegramUserId,
				Name:           image.Name,
				Url:            image.Url,
				Size:           image.Size,
				CreatedAt:      image.CreatedAt.AsTime(),
				UpdatedAt:      image.UpdatedAt.AsTime(),
			})
		}
	}
	return &response.ProfileResponseDto{
		TelegramUserId: r.TelegramUserId,
		DisplayName:    r.DisplayName,
		Birthday:       r.Birthday.AsTime(),
		Gender:         r.Gender,
		Location:       r.Location,
		Description:    r.Description,
		Height:         r.Height,
		Weight:         r.Weight,
		CreatedAt:      r.CreatedAt.AsTime(),
		UpdatedAt:      r.UpdatedAt.AsTime(),
		LastOnline:     r.LastOnline.AsTime(),
		Navigator:      navigatorResponse,
		Filter: &response.FilterResponseDto{
			TelegramUserId: r.Filter.TelegramUserId,
			SearchGender:   r.Filter.SearchGender,
			LookingFor:     r.Filter.LookingFor,
			AgeFrom:        r.Filter.AgeFrom,
			AgeTo:          r.Filter.AgeTo,
			Distance:       r.Filter.Distance,
			Page:           r.Filter.Page,
			Size:           r.Filter.Size,
		},
		Telegram: &response.TelegramResponseDto{
			UserId:          r.Telegram.UserId,
			Username:        r.Telegram.Username,
			FirstName:       r.Telegram.FirstName,
			LastName:        r.Telegram.LastName,
			LanguageCode:    r.Telegram.LanguageCode,
			AllowsWriteToPm: r.Telegram.AllowsWriteToPm,
			QueryId:         r.Telegram.QueryId,
		},
		Status: &response.StatusResponseDto{
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

func (pm *ProfileMapper) MapToShortInfoResponse(r *pb.ProfileShortInfoResponse) *response.ProfileShortInfoResponseDto {
	return &response.ProfileShortInfoResponseDto{
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

func (pm *ProfileMapper) MapToGetDetailRequest(
	r *request.ProfileGetDetailRequestDto, viewedTelegramUserId string) *pb.ProfileGetDetailRequest {
	return &pb.ProfileGetDetailRequest{
		TelegramUserId:       r.TelegramUserId,
		Latitude:             r.Latitude,
		Longitude:            r.Longitude,
		ViewedTelegramUserId: viewedTelegramUserId,
	}
}

func (pm *ProfileMapper) MapToDetailResponse(r *pb.ProfileDetailResponse) *response.ProfileDetailResponseDto {
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
			Id:        r.Like.Id,
			IsLiked:   r.Like.IsLiked,
			UpdatedAt: r.Like.UpdatedAt.AsTime(),
		}
	}
	images := make([]*response.ImageResponseDto, 0)
	if len(r.Images) > 0 {
		for _, image := range r.Images {
			images = append(images, &response.ImageResponseDto{
				Id:   image.Id,
				Name: image.Name,
				Url:  image.Url,
			})
		}
	}
	return &response.ProfileDetailResponseDto{
		TelegramUserId: r.TelegramUserId,
		DisplayName:    r.DisplayName,
		Birthday:       r.Birthday.AsTime(),
		Location:       r.Location,
		Description:    r.Description,
		Height:         r.Height,
		Weight:         r.Weight,
		Navigator:      navigatorResponse,
		Status: &response.StatusResponseDto{
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

func (pm *ProfileMapper) MapToGetShortInfoRequest(telegramUserId string) *pb.ProfileGetShortInfoRequest {
	return &pb.ProfileGetShortInfoRequest{
		TelegramUserId: telegramUserId,
	}
}

func (pm *ProfileMapper) MapToListRequest(
	r *request.ProfileGetListRequestDto) *pb.ProfileGetListRequest {
	return &pb.ProfileGetListRequest{
		TelegramUserId: r.TelegramUserId,
		SearchGender:   r.SearchGender,
		LookingFor:     r.LookingFor,
		AgeFrom:        r.AgeFrom,
		AgeTo:          r.AgeTo,
		Distance:       r.Distance,
		Page:           r.Page,
		Size:           r.Size,
		Latitude:       r.Latitude,
		Longitude:      r.Longitude,
	}
}

func (pm *ProfileMapper) MapToListResponse(r *pb.ProfileListResponse) *response.ProfileListResponseDto {
	paginationEntity := &entity.PaginationEntity{
		HasPrevious:   r.HasPrevious,
		HasNext:       r.HasNext,
		Page:          r.Page,
		Size:          r.Size,
		TotalEntities: r.TotalEntities,
		TotalPages:    r.TotalPages,
	}
	profileContent := make([]*response.ProfileListItemResponseDto, 0)
	if len(r.Content) > 0 {
		for _, c := range r.Content {
			profileContent = append(profileContent, &response.ProfileListItemResponseDto{
				TelegramUserId: c.TelegramUserId,
				Distance:       c.Distance,
				Url:            c.Url,
				IsOnline:       c.IsOnline,
				LastOnline:     c.LastOnline.AsTime(),
			})
		}
	}
	return &response.ProfileListResponseDto{
		PaginationEntity: paginationEntity,
		Content:          profileContent,
	}
}

func (pm *ProfileMapper) MapToImageByTelegramUserIdRequest(
	telegramUserId, fileName string) *pb.GetImageByTelegramUserIdRequest {
	return &pb.GetImageByTelegramUserIdRequest{
		TelegramUserId: telegramUserId,
		FileName:       fileName,
	}
}

func (pm *ProfileMapper) MapToImageByTelegramUserIdResponse(
	r *pb.ImageByTelegramUserIdResponse) []byte {
	return r.File
}

func (pm *ProfileMapper) MapToFilterUpdateRequest(r *request.FilterUpdateRequestDto) *pb.FilterUpdateRequest {
	return &pb.FilterUpdateRequest{
		TelegramUserId: r.TelegramUserId,
		SearchGender:   r.SearchGender,
		AgeFrom:        r.AgeFrom,
		AgeTo:          r.AgeTo,
	}
}

func (pm *ProfileMapper) MapToBlockAddRequest(r *request.BlockAddRequestDto) *pb.BlockAddRequest {
	return &pb.BlockAddRequest{
		TelegramUserId:        r.TelegramUserId,
		BlockedTelegramUserId: r.BlockedTelegramUserId,
	}
}

func (pm *ProfileMapper) MapToLikeAddRequest(r *request.LikeAddRequestDto, locale string) *pb.LikeAddRequest {
	return &pb.LikeAddRequest{
		TelegramUserId:      r.TelegramUserId,
		LikedTelegramUserId: r.LikedTelegramUserId,
		Locale:              locale,
	}
}

func (pm *ProfileMapper) MapToLikeAddResponse(r *pb.LikeAddResponse) *response.LikeResponseDto {
	if r == nil {
		return nil
	}
	return &response.LikeResponseDto{
		Id:        r.Id,
		IsLiked:   r.IsLiked,
		UpdatedAt: r.UpdatedAt.AsTime(),
	}
}

func (pm *ProfileMapper) MapToLikeUpdateRequest(r *request.LikeUpdateRequestDto) *pb.LikeUpdateRequest {
	return &pb.LikeUpdateRequest{
		Id:             r.Id,
		TelegramUserId: r.TelegramUserId,
		IsLiked:        r.IsLiked,
	}
}

func (pm *ProfileMapper) MapToLikeUpdateResponse(r *pb.LikeUpdateResponse) *response.LikeResponseDto {
	return &response.LikeResponseDto{
		Id:        r.Id,
		IsLiked:   r.IsLiked,
		UpdatedAt: r.UpdatedAt.AsTime(),
	}
}

func (pm *ProfileMapper) MapToComplaintAddRequest(r *request.ComplaintAddRequestDto) *pb.ComplaintAddRequest {
	return &pb.ComplaintAddRequest{
		TelegramUserId:         r.TelegramUserId,
		CriminalTelegramUserId: r.CriminalTelegramUserId,
		Reason:                 r.Reason,
	}
}

func (pm *ProfileMapper) MapToComplaintAddResponse(r *pb.ComplaintAddResponse) *entity.ComplaintEntity {
	return &entity.ComplaintEntity{
		Id:                     r.Id,
		TelegramUserId:         r.TelegramUserId,
		CriminalTelegramUserId: r.CriminalTelegramUserId,
		Reason:                 r.Reason,
		CreatedAt:              r.CreatedAt.AsTime(),
		UpdatedAt:              r.UpdatedAt.AsTime(),
	}
}

func (pm *ProfileMapper) MapToUpdateCoordinatesRequest(
	r *request.NavigatorUpdateRequestDto) *pb.NavigatorUpdateRequest {
	return &pb.NavigatorUpdateRequest{
		TelegramUserId: r.TelegramUserId,
		Latitude:       r.Latitude,
		Longitude:      r.Longitude,
	}
}
