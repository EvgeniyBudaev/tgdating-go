package mapper

import (
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/gateway/shared/enum"
)

type ProfileMapper struct {
}

func (pm *ProfileMapper) MapToAddRequest(
	r *request.ProfileAddRequestDto, fileList []*pb.FileMetadata) *pb.ProfileAddRequest {
	return &pb.ProfileAddRequest{
		DisplayName:             r.DisplayName,
		Age:                     r.Age,
		Gender:                  string(r.Gender),
		SearchGender:            string(r.SearchGender),
		Description:             r.Description,
		TelegramUserId:          r.TelegramUserId,
		TelegramUsername:        r.TelegramUsername,
		TelegramFirstName:       r.TelegramFirstName,
		TelegramLastName:        r.TelegramLastName,
		TelegramLanguageCode:    r.TelegramLanguageCode,
		TelegramAllowsWriteToPm: r.TelegramAllowsWriteToPm,
		TelegramQueryId:         r.TelegramQueryId,
		CountryCode:             r.CountryCode,
		CountryName:             r.CountryName,
		City:                    r.City,
		Latitude:                r.Latitude,
		Longitude:               r.Longitude,
		AgeFrom:                 r.AgeFrom,
		AgeTo:                   r.AgeTo,
		Distance:                r.Distance,
		Page:                    r.Page,
		Size:                    r.Size,
		IsLeftHand:              r.IsLeftHand,
		Measurement:             string(r.Measurement),
		Files:                   fileList,
	}
}

func (pm *ProfileMapper) MapToUpdateRequest(
	r *request.ProfileUpdateRequestDto, fileList []*pb.FileMetadata) *pb.ProfileUpdateRequest {
	return &pb.ProfileUpdateRequest{
		DisplayName:             r.DisplayName,
		Age:                     r.Age,
		Gender:                  string(r.Gender),
		SearchGender:            string(r.SearchGender),
		Description:             r.Description,
		TelegramUserId:          r.TelegramUserId,
		TelegramUsername:        r.TelegramUsername,
		TelegramFirstName:       r.TelegramFirstName,
		TelegramLastName:        r.TelegramLastName,
		TelegramLanguageCode:    r.TelegramLanguageCode,
		TelegramAllowsWriteToPm: r.TelegramAllowsWriteToPm,
		TelegramQueryId:         r.TelegramQueryId,
		CountryCode:             r.CountryCode,
		CountryName:             r.CountryName,
		City:                    r.City,
		Latitude:                r.Latitude,
		Longitude:               r.Longitude,
		AgeFrom:                 r.AgeFrom,
		AgeTo:                   r.AgeTo,
		Distance:                r.Distance,
		Page:                    r.Page,
		Size:                    r.Size,
		IsImages:                r.IsImages,
		Measurement:             string(r.Measurement),
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

func (pm *ProfileMapper) MapToGetRequest(
	r *request.ProfileGetByTelegramUserIdRequestDto, telegramUserId string) *pb.ProfileGetRequest {
	return &pb.ProfileGetRequest{
		TelegramUserId: telegramUserId,
		CountryCode:    r.CountryCode,
		CountryName:    r.CountryName,
		City:           r.City,
		Latitude:       r.Latitude,
		Longitude:      r.Longitude,
	}
}

func (pm *ProfileMapper) MapToByTelegramUserIdResponse(
	r *pb.ProfileResponse) *response.ProfileResponseDto {
	var navigatorResponse *response.NavigatorResponseDto
	if r.Navigator != nil {
		location := &entity.PointEntity{
			Latitude:  r.Navigator.Location.Latitude,
			Longitude: r.Navigator.Location.Longitude,
		}
		navigatorResponse = &response.NavigatorResponseDto{
			Location: location,
		}
	}
	images := make([]*response.ImageResponseDto, 0)
	if len(r.Images) > 0 {
		for _, image := range r.Images {
			images = append(images, &response.ImageResponseDto{
				Id:             image.Id,
				TelegramUserId: image.TelegramUserId,
				Name:           image.Name,
				Url:            image.Url,
			})
		}
	}
	return &response.ProfileResponseDto{
		TelegramUserId: r.TelegramUserId,
		DisplayName:    r.DisplayName,
		Age:            r.Age,
		Gender:         r.Gender,
		Description:    r.Description,
		Navigator:      navigatorResponse,
		Filter: &response.FilterResponseDto{
			SearchGender: r.Filter.SearchGender,
			AgeFrom:      r.Filter.AgeFrom,
			AgeTo:        r.Filter.AgeTo,
			Distance:     r.Filter.Distance,
			Page:         r.Filter.Page,
			Size:         r.Filter.Size,
		},
		Status: &response.StatusResponseDto{
			IsBlocked:        r.Status.IsBlocked,
			IsFrozen:         r.Status.IsFrozen,
			IsHiddenAge:      r.Status.IsHiddenAge,
			IsHiddenDistance: r.Status.IsHiddenDistance,
			IsInvisible:      r.Status.IsInvisible,
			IsLeftHand:       r.Status.IsLeftHand,
			IsPremium:        r.Status.IsPremium,
		},
		Settings: &response.SettingsResponseDto{
			Measurement: enum.Measurement(r.Settings.Measurement),
		},
		Images: images,
	}
}

func (pm *ProfileMapper) MapToShortInfoResponse(r *pb.ProfileShortInfoResponse) *response.ProfileShortInfoResponseDto {
	return &response.ProfileShortInfoResponseDto{
		TelegramUserId: r.TelegramUserId,
		IsBlocked:      r.IsBlocked,
		IsFrozen:       r.IsFrozen,
		IsPremium:      r.IsPremium,
		AvailableUntil: r.AvailableUntil.AsTime(),
		SearchGender:   r.SearchGender,
		AgeFrom:        r.AgeFrom,
		AgeTo:          r.AgeTo,
		Distance:       r.Distance,
		Page:           r.Page,
		Size:           r.Size,
		LanguageCode:   r.LanguageCode,
		Measurement:    enum.Measurement(r.Measurement),
	}
}

func (pm *ProfileMapper) MapToGetDetailRequest(
	r *request.ProfileGetDetailRequestDto, viewedTelegramUserId string) *pb.ProfileGetDetailRequest {
	return &pb.ProfileGetDetailRequest{
		TelegramUserId:       r.TelegramUserId,
		CountryCode:          r.CountryCode,
		CountryName:          r.CountryName,
		City:                 r.City,
		Latitude:             r.Latitude,
		Longitude:            r.Longitude,
		ViewedTelegramUserId: viewedTelegramUserId,
	}
}

func (pm *ProfileMapper) MapToDetailResponse(r *pb.ProfileDetailResponse) *response.ProfileDetailResponseDto {
	var navigatorResponse *response.NavigatorDetailResponseDto
	if r.Navigator != nil {
		navigatorResponse = &response.NavigatorDetailResponseDto{
			CountryName: r.Navigator.CountryName,
			City:        r.Navigator.City,
			Distance:    r.Navigator.Distance,
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
				Id:             image.Id,
				TelegramUserId: image.TelegramUserId,
				Name:           image.Name,
				Url:            image.Url,
			})
		}
	}
	return &response.ProfileDetailResponseDto{
		TelegramUserId: r.TelegramUserId,
		DisplayName:    r.DisplayName,
		Age:            r.Age,
		Gender:         enum.Gender(r.Gender),
		Description:    r.Description,
		LastOnline:     r.LastOnline.AsTime(),
		Navigator:      navigatorResponse,
		Status: &response.StatusResponseDto{
			IsBlocked:        r.Status.IsBlocked,
			IsFrozen:         r.Status.IsFrozen,
			IsHiddenAge:      r.Status.IsHiddenAge,
			IsHiddenDistance: r.Status.IsHiddenDistance,
			IsInvisible:      r.Status.IsInvisible,
			IsLeftHand:       r.Status.IsLeftHand,
			IsPremium:        r.Status.IsPremium,
		},
		Settings: &response.SettingsResponseDto{
			Measurement: enum.Measurement(r.Settings.Measurement),
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
		CountryCode:    r.CountryCode,
		CountryName:    r.CountryName,
		City:           r.City,
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
				IsLiked:        c.IsLiked,
				LastOnline:     c.LastOnline.AsTime(),
			})
		}
	}
	return &response.ProfileListResponseDto{
		PaginationEntity: paginationEntity,
		Content:          profileContent,
	}
}

func (pm *ProfileMapper) MapToCheckProfileExistsRequest(telegramUserId string) *pb.CheckProfileExistsRequest {
	return &pb.CheckProfileExistsRequest{
		TelegramUserId: telegramUserId,
	}
}

func (pm *ProfileMapper) MapToCheckProfileExistsResponse(isExists bool) *response.CheckProfileExistsResponseDto {
	return &response.CheckProfileExistsResponseDto{
		IsExists: isExists,
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

func (pm *ProfileMapper) MapToFilterGetRequest(telegramUserId string) *pb.FilterGetRequest {
	return &pb.FilterGetRequest{
		TelegramUserId: telegramUserId,
	}
}

func (pm *ProfileMapper) MapToTelegramGetRequest(telegramUserId string) *pb.TelegramGetRequest {
	return &pb.TelegramGetRequest{
		TelegramUserId: telegramUserId,
	}
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

func (pm *ProfileMapper) MapToGetBlockedListRequest(telegramUserId string) *pb.GetBlockedListRequest {
	return &pb.GetBlockedListRequest{
		TelegramUserId: telegramUserId,
	}
}

func (pm *ProfileMapper) MapToUnblockRequest(r *request.UnblockRequestDto) *pb.UnblockRequest {
	return &pb.UnblockRequest{
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

func (pm *ProfileMapper) MapToGetImageLastRequest(telegramUserId string) *pb.GetImageLastByTelegramUserIdRequest {
	return &pb.GetImageLastByTelegramUserIdRequest{
		TelegramUserId: telegramUserId,
	}
}

func (pm *ProfileMapper) MapToLikeUpdateRequest(r *request.LikeUpdateRequestDto) *pb.LikeUpdateRequest {
	return &pb.LikeUpdateRequest{
		Id:             r.Id,
		TelegramUserId: r.TelegramUserId,
		IsLiked:        r.IsLiked,
	}
}

func (pm *ProfileMapper) MapToLikeGetLastRequest(telegramUserId string) *pb.LikeGetLastRequest {
	return &pb.LikeGetLastRequest{
		TelegramUserId: telegramUserId,
	}
}

func (pm *ProfileMapper) MapToLikeGetLastResponse(r *pb.LikeGetLastResponse) *entity.LikeEntity {
	if r.Like == nil {
		return nil
	}
	return &entity.LikeEntity{
		Id:                  r.Like.Id,
		TelegramUserId:      r.Like.TelegramUserId,
		LikedTelegramUserId: r.Like.LikedTelegramUserId,
		IsLiked:             r.Like.IsLiked,
		CreatedAt:           r.Like.CreatedAt.AsTime(),
		UpdatedAt:           r.Like.UpdatedAt.AsTime(),
	}
}

func (pm *ProfileMapper) MapToComplaintAddRequest(r *request.ComplaintAddRequestDto) *pb.ComplaintAddRequest {
	return &pb.ComplaintAddRequest{
		TelegramUserId:         r.TelegramUserId,
		CriminalTelegramUserId: r.CriminalTelegramUserId,
		Reason:                 r.Reason,
	}
}

func (pm *ProfileMapper) MapToPaymentAddRequest(r *request.PaymentAddRequestDto) *pb.PaymentAddRequest {
	return &pb.PaymentAddRequest{
		TelegramUserId: r.TelegramUserId,
		Price:          r.Price,
		Currency:       r.Currency,
		Tariff:         string(r.Tariff),
	}
}

func (pm *ProfileMapper) MapToCheckPremiumRequest(telegramUserId string) *pb.CheckPremiumRequest {
	return &pb.CheckPremiumRequest{
		TelegramUserId: telegramUserId,
	}
}

func (pm *ProfileMapper) MapToCheckPremiumResponse(r *pb.CheckPremiumResponse) *response.PremiumResponseDto {
	return &response.PremiumResponseDto{
		IsPremium:      r.IsPremium,
		AvailableUntil: r.AvailableUntil.AsTime(),
	}
}

func (pm *ProfileMapper) MapToUpdateSettingsRequest(
	r *request.ProfileUpdateSettingsRequestDto) *pb.UpdateSettingsRequest {
	return &pb.UpdateSettingsRequest{
		TelegramUserId: r.TelegramUserId,
		IsHiddenAge:    r.IsHiddenAge,
		Measurement:    string(r.Measurement),
	}
}

func (pm *ProfileMapper) MapToUpdateCoordinatesRequest(
	r *request.NavigatorUpdateRequestDto) *pb.NavigatorUpdateRequest {
	return &pb.NavigatorUpdateRequest{
		TelegramUserId: r.TelegramUserId,
		CountryCode:    r.CountryCode,
		CountryName:    r.CountryName,
		City:           r.City,
		Latitude:       r.Latitude,
		Longitude:      r.Longitude,
	}
}
