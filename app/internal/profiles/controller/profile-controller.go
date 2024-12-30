package controller

import (
	"context"
	"fmt"
	pb "github.com/EvgeniyBudaev/tgdating-go/app/contracts/proto/profiles"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/controller/mapper"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/repository/psql"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

const (
	errorFilePath = "internal/profiles/controller/profile-controller.go"
)

type ProfileController struct {
	logger  logger.Logger
	service ProfileService
	pb.UnimplementedProfileServer
}

func NewProfileController(l logger.Logger, ps ProfileService) *ProfileController {
	return &ProfileController{
		logger:  l,
		service: ps,
	}
}

func (pc *ProfileController) AddProfile(ctx context.Context, in *pb.ProfileAddRequest) (*pb.ProfileAddResponse, error) {
	pc.logger.Info("POST /api/v1/profiles")
	fileList := make([]*entity.FileMetadata, 0)
	if len(in.Files) > 0 {
		for _, file := range in.Files {
			fileList = append(fileList, &entity.FileMetadata{
				Filename: file.Filename,
				Size:     file.Size,
				Content:  file.Content,
			})
		}
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	profileRequest := profileMapper.MapControllerToAddRequest(in, fileList)
	profileAdded, err := pc.service.AddProfile(ctx, profileRequest)
	if err != nil {
		return nil, err
	}
	profileResponse := profileMapper.MapControllerToAddResponse(profileAdded)
	return profileResponse, nil
}

func (pc *ProfileController) UpdateProfile(
	ctx context.Context, in *pb.ProfileUpdateRequest) (*pb.ProfileResponse, error) {
	pc.logger.Info("PUT /api/v1/profiles")
	fileList := make([]*entity.FileMetadata, 0)
	if len(in.Files) > 0 {
		for _, file := range in.Files {
			fileList = append(fileList, &entity.FileMetadata{
				Filename: file.Filename,
				Size:     file.Size,
				Content:  file.Content,
			})
		}
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	profileRequest := profileMapper.MapControllerToUpdateRequest(in, fileList)
	profileUpdated, err := pc.service.UpdateProfile(ctx, profileRequest)
	if err != nil {
		return nil, err
	}
	profileResponse := profileMapper.MapControllerResponse(profileUpdated)
	return profileResponse, nil
}

func (pc *ProfileController) FreezeProfile(
	ctx context.Context, in *pb.ProfileFreezeRequest) (*pb.ProfileFreezeResponse, error) {
	pc.logger.Info("POST /api/v1/profiles/freeze")
	req := &request.ProfileFreezeRequestDto{
		TelegramUserId: in.TelegramUserId,
	}
	profileDeleted, err := pc.service.FreezeProfile(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ProfileFreezeResponse{
		Success: profileDeleted.Success,
	}, nil
}

func (pc *ProfileController) RestoreProfile(
	ctx context.Context, in *pb.ProfileRestoreRequest) (*pb.ProfileRestoreResponse, error) {
	pc.logger.Info("POST /api/v1/profiles/restore")
	req := &request.ProfileRestoreRequestDto{
		TelegramUserId: in.TelegramUserId,
	}
	profileDeleted, err := pc.service.RestoreProfile(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ProfileRestoreResponse{
		Success: profileDeleted.Success,
	}, nil
}

func (pc *ProfileController) DeleteProfile(
	ctx context.Context, in *pb.ProfileDeleteRequest) (*pb.ProfileDeleteResponse, error) {
	pc.logger.Info("DELETE /api/v1/profiles")
	req := &request.ProfileDeleteRequestDto{
		TelegramUserId: in.TelegramUserId,
	}
	profileDeleted, err := pc.service.DeleteProfile(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.ProfileDeleteResponse{
		Success: profileDeleted.Success,
	}, nil
}

func (pc *ProfileController) GetProfile(
	ctx context.Context, in *pb.ProfileGetRequest) (*pb.ProfileResponse, error) {
	pc.logger.Info("GET /api/v1/profiles/telegram/:telegramUserId")
	req := &request.ProfileGetRequestDto{
		Latitude:  in.Latitude,
		Longitude: in.Longitude,
	}
	profileByTelegramUserId, err := pc.service.GetProfile(ctx, in.TelegramUserId, req)
	if err != nil {
		if errors.Is(err, psql.ErrNotRowFound) {
			return nil, status.Errorf(codes.NotFound, psql.ErrNotRowFoundMessage)
		}
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	profileResponse := profileMapper.MapControllerResponse(profileByTelegramUserId)
	return profileResponse, nil
}

func (pc *ProfileController) GetProfileDetail(
	ctx context.Context, in *pb.ProfileGetDetailRequest) (*pb.ProfileDetailResponse, error) {
	pc.logger.Info("GET /api/v1/profiles/detail/:viewedTelegramUserId")
	req := &request.ProfileGetDetailRequestDto{
		TelegramUserId: in.TelegramUserId,
		Latitude:       in.Latitude,
		Longitude:      in.Longitude,
	}
	profileDetail, err := pc.service.GetProfileDetail(ctx, in.ViewedTelegramUserId, req)
	if err != nil {
		if errors.Is(err, psql.ErrNotRowFound) {
			return nil, status.Errorf(codes.NotFound, psql.ErrNotRowFoundMessage)
		}
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	profileResponse := profileMapper.MapControllerToDetailResponse(profileDetail)
	return profileResponse, nil
}

func (pc *ProfileController) GetProfileShortInfo(
	ctx context.Context, in *pb.ProfileGetShortInfoRequest) (*pb.ProfileShortInfoResponse, error) {
	pc.logger.Info("GET /api/v1/profiles/short/:telegramUserId")
	profileShortInfo, err := pc.service.GetProfileShortInfo(ctx, in.TelegramUserId)
	if err != nil {
		if errors.Is(err, psql.ErrNotRowFound) {
			return nil, status.Errorf(codes.NotFound, psql.ErrNotRowFoundMessage)
		}
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	profileResponse := profileMapper.MapControllerToShortInfoResponse(profileShortInfo)
	return profileResponse, nil
}

func (pc *ProfileController) GetProfileList(
	ctx context.Context, in *pb.ProfileGetListRequest) (*pb.ProfileListResponse, error) {
	pc.logger.Info("GET api/v1/profiles/list")
	req := &request.ProfileGetListRequestDto{
		TelegramUserId: in.TelegramUserId,
		Latitude:       in.Latitude,
		Longitude:      in.Longitude,
	}
	profileList, err := pc.service.GetProfileList(ctx, req)
	if err != nil {
		if errors.Is(err, psql.ErrNotRowFound) {
			return nil, status.Errorf(codes.NotFound, psql.ErrNotRowFoundMessage)
		}
		if errors.Is(err, psql.ErrNotRowsFound) {
			return nil, status.Errorf(codes.NotFound, psql.ErrNotRowsFoundMessage)
		}
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	profileResponse := profileMapper.MapControllerToListResponse(profileList)
	return profileResponse, nil
}

func (pc *ProfileController) GetImageByTelegramUserId(
	ctx context.Context, in *pb.GetImageByTelegramUserIdRequest) (*pb.ImageByTelegramUserIdResponse, error) {
	pc.logger.Info("GET /api/v1/profiles/:telegramUserId/images/:fileName")
	file, err := pc.service.GetImageByTelegramUserId(ctx, in.TelegramUserId, in.FileName)
	if err != nil {
		return nil, err
	}
	fileResponse := &pb.ImageByTelegramUserIdResponse{
		File: file,
	}
	return fileResponse, nil
}

func (pc *ProfileController) GetImageById(ctx context.Context, in *pb.GetImageByIdRequest) (*pb.ImageResponse, error) {
	pc.logger.Info("GET /api/v1/profiles/images/:id")
	imageById, err := pc.service.GetImageById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	imageResponse := profileMapper.MapControllerToImageResponse(imageById)
	return imageResponse, nil
}

func (pc *ProfileController) DeleteImage(
	ctx context.Context, in *pb.ImageDeleteRequest) (*pb.ImageDeleteResponse, error) {
	pc.logger.Info("DELETE /api/v1/profiles/images/:id")
	req := &request.ImageDeleteRequestDto{
		Id: in.Id,
	}
	fileDeleted, err := pc.service.DeleteImage(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	fileResponse := &pb.ImageDeleteResponse{
		Success: fileDeleted.Success,
	}
	return fileResponse, nil
}

func (pc *ProfileController) GetFilter(
	ctx context.Context, in *pb.FilterGetRequest) (*pb.FilterResponse, error) {
	pc.logger.Info("GET /api/v1/profiles/filters/:telegramUserId")
	filter, err := pc.service.GetFilter(ctx, in.TelegramUserId)
	if err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	filterResponse := profileMapper.MapControllerToFilterResponse(filter)
	return filterResponse, nil
}

func (pc *ProfileController) UpdateFilter(
	ctx context.Context, in *pb.FilterUpdateRequest) (*pb.FilterResponse, error) {
	pc.logger.Info("PUT /api/v1/profiles/filters")
	req := &request.FilterUpdateRequestDto{
		TelegramUserId: in.TelegramUserId,
		SearchGender:   in.SearchGender,
		AgeFrom:        in.AgeFrom,
		AgeTo:          in.AgeTo,
	}
	filterUpdated, err := pc.service.UpdateFilter(ctx, req)
	if err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	filterResponse := profileMapper.MapControllerToFilterResponse(filterUpdated)
	return filterResponse, nil
}

func (pc *ProfileController) AddBlock(ctx context.Context, in *pb.BlockAddRequest) (*pb.BlockAddResponse, error) {
	pc.logger.Info("POST /api/v1/profiles/blocks")
	req := &request.BlockAddRequestDto{
		TelegramUserId:        in.TelegramUserId,
		BlockedTelegramUserId: in.BlockedTelegramUserId,
	}
	block, err := pc.service.AddBlock(ctx, req)
	if err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	blockResponse := profileMapper.MapControllerToBlockAddResponse(block)
	return blockResponse, nil
}

func (pc *ProfileController) AddLike(ctx context.Context, in *pb.LikeAddRequest) (*pb.LikeAddResponse, error) {
	pc.logger.Info("POST /api/v1/profiles/likes")
	req := &request.LikeAddRequestDto{
		TelegramUserId:      in.TelegramUserId,
		LikedTelegramUserId: in.LikedTelegramUserId,
	}
	locale := in.Locale
	likeAdded, err := pc.service.AddLike(ctx, req, locale)
	if err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	likeResponse := profileMapper.MapControllerToLikeAddResponse(likeAdded)
	return likeResponse, nil
}

func (pc *ProfileController) UpdateLike(ctx context.Context, in *pb.LikeUpdateRequest) (*pb.LikeUpdateResponse, error) {
	pc.logger.Info("PUT /api/v1/profiles/likes")
	req := &request.LikeUpdateRequestDto{
		Id:             in.Id,
		TelegramUserId: in.TelegramUserId,
		IsLiked:        in.IsLiked,
	}
	likeUpdated, err := pc.service.UpdateLike(ctx, req)
	if err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	likeResponse := profileMapper.MapControllerToLikeUpdateResponse(likeUpdated)
	return likeResponse, nil
}

func (pc *ProfileController) GetLastLike(
	ctx context.Context, in *pb.LikeGetLastRequest) (*pb.LikeGetLastResponse, error) {
	pc.logger.Info("GET /api/v1/profiles/likes/:telegramUserId/last")
	likeEntity, err := pc.service.GetLastLike(ctx, in.TelegramUserId)
	if err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	likeResponse := profileMapper.MapControllerToLikeGetLastResponse(likeEntity)
	return likeResponse, nil
}

func (pc *ProfileController) AddComplaint(
	ctx context.Context, in *pb.ComplaintAddRequest) (*pb.ComplaintAddResponse, error) {
	pc.logger.Info("POST /api/v1/profiles/complaints")
	req := &request.ComplaintAddRequestDto{
		TelegramUserId:         in.TelegramUserId,
		CriminalTelegramUserId: in.CriminalTelegramUserId,
		Reason:                 in.Reason,
	}
	complaintAdded, err := pc.service.AddComplaint(ctx, req)
	if err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	complaintResponse := profileMapper.MapControllerToComplaintAddResponse(complaintAdded)
	return complaintResponse, nil
}

func (pc *ProfileController) UpdateCoordinates(
	ctx context.Context, in *pb.NavigatorUpdateRequest) (*pb.NavigatorUpdateResponse, error) {
	pc.logger.Info("PUT /api/v1/profiles/navigators")
	req := &request.NavigatorUpdateRequestDto{
		TelegramUserId: in.TelegramUserId,
		Latitude:       in.Latitude,
		Longitude:      in.Longitude,
	}
	updatedCoordinates, err := pc.service.UpdateCoordinates(ctx, req)
	if err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileControllerMapper{}
	updatedCoordinatesResponse := profileMapper.MapControllerToUpdateCoordinatesResponse(updatedCoordinates)
	return updatedCoordinatesResponse, nil
}

func (pc *ProfileController) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}

func (pc *ProfileController) convertToUint64(name, value string) (uint64, error) {
	if value == "" {
		errorMessage := fmt.Sprintf("%s is empty", name)
		return 0, errors.New(errorMessage)
	}
	value64, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		errorMessage := pc.getErrorMessage("convertToUint64", "ParseUint")
		pc.logger.Debug(errorMessage, zap.Error(err))
		return 0, err
	}
	return value64, nil
}
