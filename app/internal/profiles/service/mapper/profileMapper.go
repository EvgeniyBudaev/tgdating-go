package mapper

import (
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ProfileMapper struct {
}

func (pm *ProfileMapper) MapControllerToResponse(r *response.ProfileResponseDto) *pb.ProfileUpdateResponse {
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
	return &pb.ProfileUpdateResponse{
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

func (pm *ProfileMapper) MapToResponse(
	pe *entity.ProfileEntity, nr *response.NavigatorResponseDto, fr *response.FilterResponseDto,
	tr *response.TelegramResponseDto, il []*entity.ImageEntity, isOnline bool,
) *response.ProfileResponseDto {
	return &response.ProfileResponseDto{
		SessionId:      pe.SessionId,
		DisplayName:    pe.DisplayName,
		Birthday:       pe.Birthday,
		Gender:         pe.Gender,
		Location:       pe.Location,
		Description:    pe.Description,
		Height:         pe.Height,
		Weight:         pe.Weight,
		IsDeleted:      pe.IsDeleted,
		IsBlocked:      pe.IsBlocked,
		IsPremium:      pe.IsPremium,
		IsShowDistance: pe.IsShowDistance,
		IsInvisible:    pe.IsInvisible,
		IsOnline:       isOnline,
		CreatedAt:      pe.CreatedAt,
		UpdatedAt:      pe.UpdatedAt,
		LastOnline:     pe.LastOnline,
		Navigator:      nr,
		Filter:         fr,
		Telegram:       tr,
		Images:         il,
	}
}

func (pm *ProfileMapper) MapToDetailResponse(
	pe *entity.ProfileEntity, nr *response.NavigatorDetailResponseDto, br *response.BlockResponseDto,
	lr *response.LikeResponseDto, tr *response.TelegramResponseDto, il []*entity.ImageEntity, isOnline bool,
) *response.ProfileDetailResponseDto {
	return &response.ProfileDetailResponseDto{
		SessionId:      pe.SessionId,
		DisplayName:    pe.DisplayName,
		Birthday:       pe.Birthday,
		Gender:         pe.Gender,
		Location:       pe.Location,
		Description:    pe.Description,
		Height:         pe.Height,
		Weight:         pe.Weight,
		IsDeleted:      pe.IsDeleted,
		IsBlocked:      pe.IsBlocked,
		IsPremium:      pe.IsPremium,
		IsShowDistance: pe.IsShowDistance,
		IsInvisible:    pe.IsInvisible,
		IsOnline:       isOnline,
		CreatedAt:      pe.CreatedAt,
		UpdatedAt:      pe.UpdatedAt,
		LastOnline:     pe.LastOnline,
		Navigator:      nr,
		Block:          br,
		Like:           lr,
		Telegram:       tr,
		Images:         il,
	}
}

func (pm *ProfileMapper) MapToShortInfoResponse(pe *entity.ProfileEntity, imageUrl string) *response.ProfileShortInfoResponseDto {
	return &response.ProfileShortInfoResponseDto{
		SessionId: pe.SessionId,
		ImageUrl:  imageUrl,
		IsDeleted: pe.IsDeleted,
		IsBlocked: pe.IsBlocked,
	}
}

func (pm *ProfileMapper) MapControllerToAddResponse(r *response.ProfileAddResponseDto) *pb.ProfileAddResponse {
	return &pb.ProfileAddResponse{
		SessionId: r.SessionId,
	}
}

func (pm *ProfileMapper) MapToAddResponse(pe *entity.ProfileEntity) *response.ProfileAddResponseDto {
	return &response.ProfileAddResponseDto{
		SessionId: pe.SessionId,
	}
}

func (pm *ProfileMapper) MapControllerToAddRequest(
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

func (pm *ProfileMapper) MapToAddRequest(
	pr *request.ProfileAddRequestDto) *request.ProfileAddRequestRepositoryDto {
	return &request.ProfileAddRequestRepositoryDto{
		SessionId:      pr.SessionId,
		DisplayName:    pr.DisplayName,
		Birthday:       pr.Birthday,
		Gender:         pr.Gender,
		Location:       pr.Location,
		Description:    pr.Description,
		Height:         pr.Height,
		Weight:         pr.Weight,
		IsDeleted:      false,
		IsBlocked:      false,
		IsPremium:      false,
		IsShowDistance: true,
		IsInvisible:    false,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		LastOnline:     time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapControllerToUpdateRequest(
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

func (pm *ProfileMapper) MapToUpdateRequest(
	pr *request.ProfileUpdateRequestDto) *request.ProfileUpdateRequestRepositoryDto {
	return &request.ProfileUpdateRequestRepositoryDto{
		SessionId:   pr.SessionId,
		DisplayName: pr.DisplayName,
		Birthday:    pr.Birthday,
		Gender:      pr.Gender,
		Location:    pr.Location,
		Description: pr.Description,
		Height:      pr.Height,
		Weight:      pr.Weight,
		UpdatedAt:   time.Now().UTC(),
		LastOnline:  time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapToDeleteRequest(sessionId string) *request.ProfileDeleteRequestRepositoryDto {
	return &request.ProfileDeleteRequestRepositoryDto{
		SessionId:  sessionId,
		IsDeleted:  true,
		UpdatedAt:  time.Now().UTC(),
		LastOnline: time.Now().UTC(),
	}
}

func (pm *ProfileMapper) MapToListRequest(
	pr *entity.FilterEntity) *request.ProfileGetListRequestRepositoryDto {
	return &request.ProfileGetListRequestRepositoryDto{
		SessionId:    pr.SessionId,
		SearchGender: pr.SearchGender,
		LookingFor:   pr.LookingFor,
		AgeFrom:      pr.AgeFrom,
		AgeTo:        pr.AgeTo,
		Distance:     pr.Distance,
		Page:         pr.Page,
		Size:         pr.Size,
	}
}
