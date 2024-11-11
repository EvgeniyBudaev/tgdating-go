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
		Height:         r.Height,
		Weight:         r.Weight,
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
		Height:         r.Height,
		Weight:         r.Weight,
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

func (pm *ProfileMapper) MapToGetShortInfoRequest(
	r *request.ProfileGetShortInfoRequestDto, sessionId string) *pb.ProfileGetShortInfoRequest {
	return &pb.ProfileGetShortInfoRequest{
		SessionId: sessionId,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
	}
}

func (pm *ProfileMapper) MapToListRequest(
	r *request.ProfileGetListRequestDto) *pb.ProfileGetListRequest {
	return &pb.ProfileGetListRequest{
		SessionId: r.SessionId,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
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
				SessionId:  c.SessionId,
				Distance:   c.Distance,
				Url:        c.Url,
				IsOnline:   c.IsOnline,
				LastOnline: c.LastOnline.AsTime(),
			})
		}
	}
	return &response.ProfileListResponseDto{
		PaginationEntity: paginationEntity,
		Content:          profileContent,
	}
}

func (pm *ProfileMapper) MapToImageBySessionIdRequest(
	sessionId, fileName string) *pb.GetImageBySessionIdRequest {
	return &pb.GetImageBySessionIdRequest{
		SessionId: sessionId,
		FileName:  fileName,
	}
}

func (pm *ProfileMapper) MapToImageBySessionIdResponse(
	r *pb.ImageBySessionIdResponse) []byte {
	return r.File
}

func (pm *ProfileMapper) MapToFilterRequest(r *request.FilterGetRequestDto, sessionId string) *pb.FilterGetRequest {
	return &pb.FilterGetRequest{
		SessionId: sessionId,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
	}
}

func (pm *ProfileMapper) MapToFilterUpdateRequest(r *request.FilterUpdateRequestDto) *pb.FilterUpdateRequest {
	return &pb.FilterUpdateRequest{
		SessionId:    r.SessionId,
		SearchGender: r.SearchGender,
		AgeFrom:      r.AgeFrom,
		AgeTo:        r.AgeTo,
	}
}

func (pm *ProfileMapper) MapToBlockAddRequest(r *request.BlockAddRequestDto) *pb.BlockAddRequest {
	return &pb.BlockAddRequest{
		SessionId:            r.SessionId,
		BlockedUserSessionId: r.BlockedUserSessionId,
	}
}

func (pm *ProfileMapper) MapToBlockAddResponse(r *pb.BlockAddResponse) *entity.BlockEntity {
	return &entity.BlockEntity{
		Id:                   r.Id,
		SessionId:            r.SessionId,
		BlockedUserSessionId: r.BlockedUserSessionId,
		IsBlocked:            r.IsBlocked,
		CreatedAt:            r.CreatedAt.AsTime(),
		UpdatedAt:            r.UpdatedAt.AsTime(),
	}
}

func (pm *ProfileMapper) MapToLikeAddRequest(r *request.LikeAddRequestDto, locale string) *pb.LikeAddRequest {
	return &pb.LikeAddRequest{
		SessionId:      r.SessionId,
		LikedSessionId: r.LikedSessionId,
		Locale:         locale,
	}
}

func (pm *ProfileMapper) MapToLikeAddResponse(r *pb.LikeAddResponse) *response.LikeResponseDto {
	return &response.LikeResponseDto{
		Id:             r.Id,
		SessionId:      r.SessionId,
		LikedSessionId: r.LikedSessionId,
		IsLiked:        r.IsLiked,
		CreatedAt:      r.CreatedAt.AsTime(),
		UpdatedAt:      r.UpdatedAt.AsTime(),
	}
}

func (pm *ProfileMapper) MapToLikeUpdateRequest(r *request.LikeUpdateRequestDto) *pb.LikeUpdateRequest {
	return &pb.LikeUpdateRequest{
		Id:        r.Id,
		SessionId: r.SessionId,
		IsLiked:   r.IsLiked,
	}
}

func (pm *ProfileMapper) MapToLikeUpdateResponse(r *pb.LikeUpdateResponse) *response.LikeResponseDto {
	return &response.LikeResponseDto{
		Id:             r.Id,
		SessionId:      r.SessionId,
		LikedSessionId: r.LikedSessionId,
		IsLiked:        r.IsLiked,
		CreatedAt:      r.CreatedAt.AsTime(),
		UpdatedAt:      r.UpdatedAt.AsTime(),
	}
}

func (pm *ProfileMapper) MapToComplaintAddRequest(r *request.ComplaintAddRequestDto) *pb.ComplaintAddRequest {
	return &pb.ComplaintAddRequest{
		SessionId:         r.SessionId,
		CriminalSessionId: r.CriminalSessionId,
		Reason:            r.Reason,
	}
}

func (pm *ProfileMapper) MapToComplaintAddResponse(r *pb.ComplaintAddResponse) *entity.ComplaintEntity {
	return &entity.ComplaintEntity{
		Id:                r.Id,
		SessionId:         r.SessionId,
		CriminalSessionId: r.CriminalSessionId,
		Reason:            r.Reason,
		IsDeleted:         r.IsDeleted,
		CreatedAt:         r.CreatedAt.AsTime(),
		UpdatedAt:         r.UpdatedAt.AsTime(),
	}
}

func (pm *ProfileMapper) MapToUpdateCoordinatesRequest(
	r *request.NavigatorUpdateRequestDto) *pb.NavigatorUpdateRequest {
	return &pb.NavigatorUpdateRequest{
		SessionId: r.SessionId,
		Latitude:  r.Latitude,
		Longitude: r.Longitude,
	}
}
