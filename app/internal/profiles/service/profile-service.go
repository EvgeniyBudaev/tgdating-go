package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/config"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/service/mapper"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/h2non/bimg"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	errorFilePath = "internal/profiles/service/profile-service.go"
	minDistance   = 100
)

type ProfileService struct {
	logger              logger.Logger
	db                  *sql.DB
	config              *config.Config
	kafkaWriter         *kafka.Writer
	s3                  *config.S3
	uwf                 *UnitOfWorkFactory
	profileRepository   ProfileRepository
	navigatorRepository NavigatorRepository
	filterRepository    FilterRepository
	telegramRepository  TelegramRepository
	imageRepository     ImageRepository
	likeRepository      LikeRepository
	blockRepository     BlockRepository
	complaintRepository ComplaintRepository
}

func NewProfileService(
	l logger.Logger,
	db *sql.DB,
	cfg *config.Config,
	kw *kafka.Writer,
	s3 *config.S3,
	uwf *UnitOfWorkFactory,
	pr ProfileRepository,
	nr NavigatorRepository,
	fr FilterRepository,
	tr TelegramRepository,
	ir ImageRepository, lr LikeRepository, br BlockRepository, cr ComplaintRepository) *ProfileService {
	return &ProfileService{
		logger:              l,
		db:                  db,
		config:              cfg,
		kafkaWriter:         kw,
		s3:                  s3,
		uwf:                 uwf,
		profileRepository:   pr,
		navigatorRepository: nr,
		telegramRepository:  tr,
		filterRepository:    fr,
		imageRepository:     ir,
		likeRepository:      lr,
		blockRepository:     br,
		complaintRepository: cr,
	}
}

func (s *ProfileService) AddProfile(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*response.ProfileAddResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToAddRequest(pr)
	profileCreated, err := unitOfWork.ProfileRepository().Add(ctx, profileRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddProfile", "ProfileRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if pr.Longitude != nil && pr.Latitude != nil {
		longitude := *pr.Longitude
		latitude := *pr.Latitude
		navigatorMapper := &mapper.NavigatorMapper{}
		navigatorRequest := navigatorMapper.MapToAddRequest(pr.TelegramUserId, longitude, latitude)
		_, err = unitOfWork.NavigatorRepository().Add(ctx, navigatorRequest)
		if err != nil {
			errorMessage := s.getErrorMessage("AddProfile",
				"NavigatorRepository().Add")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	profileResponse := profileMapper.MapToAddResponse(profileCreated)
	if err := s.AddImageList(ctx, unitOfWork, profileCreated.TelegramUserId, pr.Files); err != nil {
		errorMessage := s.getErrorMessage("AddProfile", "AddImageList")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	filterMapper := &mapper.FilterMapper{}
	filterRequest := filterMapper.MapToAddRequest(pr)
	_, err = unitOfWork.FilterRepository().Add(ctx, filterRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddProfile", "FilterRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	telegramMapper := &mapper.TelegramMapper{}
	telegramRequest := telegramMapper.MapToAddRequest(pr)
	_, err = unitOfWork.TelegramRepository().Add(ctx, telegramRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddProfile", "TelegramRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("AddProfile", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("AddProfile", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return profileResponse, err
}

func (s *ProfileService) UpdateProfile(
	ctx context.Context, pr *request.ProfileUpdateRequestDto) (*response.ProfileResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	if err := s.checkUserExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile", "checkUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToUpdateRequest(pr)
	profileEntity, err := unitOfWork.ProfileRepository().Update(ctx, profileRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile",
			"ProfileRepository().Update")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if err := s.updateImageList(ctx, unitOfWork, profileEntity.TelegramUserId, pr.Files); err != nil {
		return nil, err
	}
	var navigatorResponse *response.NavigatorResponseDto
	if pr.Longitude != nil && pr.Latitude != nil {
		longitude := *pr.Longitude
		latitude := *pr.Latitude
		navigatorResponse, err = s.updateNavigator(ctx, pr.TelegramUserId, longitude, latitude)
		if err != nil {
			errorMessage := s.getErrorMessage("UpdateProfile", "updateNavigator")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	filterMapper := &mapper.FilterMapper{}
	filterRequest := filterMapper.MapProfileToUpdateRequest(pr)
	filterEntity, err := unitOfWork.FilterRepository().Update(ctx, filterRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile",
			"FilterRepository().Update")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	telegramMapper := &mapper.TelegramMapper{}
	telegramRequest := telegramMapper.MapToUpdateRequest(pr)
	telegramEntity, err := unitOfWork.TelegramRepository().Update(ctx, telegramRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile",
			"telegramRepository.Update")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	telegramResponse := telegramMapper.MapToResponse(telegramEntity)
	imageEntityList, err := unitOfWork.ImageRepository().SelectListByTelegramUserId(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile",
			"ImageRepository().SelectListByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	isOnline := s.checkIsOnline(profileEntity.LastOnline)
	profileResponse := profileMapper.MapToResponse(profileEntity, navigatorResponse, filterResponse, telegramResponse,
		imageEntityList, isOnline)
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("UpdateProfile", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return profileResponse, nil
}

func (s *ProfileService) FreezeProfile(
	ctx context.Context, pr *request.ProfileFreezeRequestDto) (*response.ResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	if err := s.checkUserExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("FreezeProfile", "checkUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToFreezeRequest(pr.TelegramUserId)
	_, err := unitOfWork.ProfileRepository().Freeze(ctx, profileRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("FreezeProfile",
			"ProfileRepository().Freeze")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileResponse := &response.ResponseDto{
		Success: true,
	}
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("FreezeProfile", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("FreezeProfile", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return profileResponse, err
}

func (s *ProfileService) RestoreProfile(
	ctx context.Context, pr *request.ProfileRestoreRequestDto) (*response.ResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	if err := s.checkUserExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("RestoreProfile", "checkUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToRestoreRequest(pr.TelegramUserId)
	_, err := unitOfWork.ProfileRepository().Restore(ctx, profileRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("RestoreProfile",
			"profileRepository.Restore")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileResponse := &response.ResponseDto{
		Success: true,
	}
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("RestoreProfile", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("RestoreProfile", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return profileResponse, err
}

func (s *ProfileService) DeleteProfile(
	ctx context.Context, pr *request.ProfileDeleteRequestDto) (*response.ResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	if err := s.checkUserExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile", "checkUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if err := s.deleteImageListByS3(ctx, unitOfWork, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile", "deleteImageListByS3")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	_, err := unitOfWork.ProfileRepository().Delete(ctx, pr)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile",
			"ProfileRepository().Delete")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileResponse := &response.ResponseDto{
		Success: true,
	}
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("DeleteProfile", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return profileResponse, err
}

func (s *ProfileService) GetProfileByTelegramUserId(ctx context.Context, telegramUserId string,
	pr *request.ProfileGetByTelegramUserIdRequestDto) (*response.ProfileResponseDto, error) {
	if err := s.checkUserExists(ctx, telegramUserId); err != nil {
		return nil, err
	}
	err := s.updateLastOnline(ctx, telegramUserId)
	if err != nil {
		return nil, err
	}
	if pr.Longitude != nil && pr.Latitude != nil {
		longitude := *pr.Longitude
		latitude := *pr.Latitude
		_, err = s.updateNavigator(ctx, telegramUserId, longitude, latitude)
		if err != nil {
			errorMessage := s.getErrorMessage("GetProfileByTelegramUserId",
				"updateNavigator")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	profileEntity, err := s.profileRepository.FindByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileByTelegramUserId",
			"profileRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	navigatorEntity, _ := s.navigatorRepository.FindByTelegramUserId(ctx, telegramUserId)
	var navigatorResponse *response.NavigatorResponseDto
	if navigatorEntity != nil {
		longitude := navigatorEntity.Location.Longitude
		latitude := navigatorEntity.Location.Latitude
		navigatorMapper := &mapper.NavigatorMapper{}
		navigatorResponse = navigatorMapper.MapToResponse(telegramUserId, longitude, latitude)
	}
	filterMapper := &mapper.FilterMapper{}
	filterEntity, err := s.filterRepository.FindByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileByTelegramUserId",
			"filterRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	telegramEntity, err := s.telegramRepository.FindByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileByTelegramUserId",
			"telegramRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	telegramMapper := &mapper.TelegramMapper{}
	telegramResponse := telegramMapper.MapToResponse(telegramEntity)
	isOnline := s.checkIsOnline(profileEntity.LastOnline)
	imageEntityList, err := s.imageRepository.SelectListByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileByTelegramUserId",
			"imageRepository.SelectListByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileResponse := profileMapper.MapToResponse(
		profileEntity, navigatorResponse, filterResponse, telegramResponse, imageEntityList, isOnline)
	return profileResponse, err
}

func (s *ProfileService) GetProfileDetail(ctx context.Context, viewedTelegramUserId string,
	pr *request.ProfileGetDetailRequestDto) (*response.ProfileDetailResponseDto, error) {
	if err := s.checkUserExists(ctx, pr.TelegramUserId); err != nil {
		return nil, err
	}
	err := s.updateLastOnline(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileDetail", "updateLastOnline")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if pr.Longitude != nil && pr.Latitude != nil {
		longitude := *pr.Longitude
		latitude := *pr.Latitude
		_, err = s.updateNavigator(ctx, pr.TelegramUserId, longitude, latitude)
		if err != nil {
			errorMessage := s.getErrorMessage("GetProfileDetail", "updateNavigator")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	profileMapper := &mapper.ProfileMapper{}
	profileEntity, err := s.profileRepository.FindByTelegramUserId(ctx, viewedTelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileDetail",
			"profileRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	navigatorMapper := &mapper.NavigatorMapper{}
	navigatorEntity, _ := s.navigatorRepository.FindByTelegramUserId(ctx, pr.TelegramUserId)
	navigatorViewedEntity, _ := s.navigatorRepository.FindByTelegramUserId(ctx, viewedTelegramUserId)
	var navigatorResponse *response.NavigatorDetailResponseDto
	if navigatorEntity != nil && navigatorViewedEntity != nil {
		navigatorDistanceResponse, err := s.navigatorRepository.FindDistance(ctx, navigatorEntity, navigatorViewedEntity)
		if err != nil {
			errorMessage := s.getErrorMessage("GetProfileDetail",
				"navigatorRepository.FindDistance")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
		distance := navigatorDistanceResponse.Distance
		if distance < minDistance {
			distance = minDistance
		}
		navigatorResponse = navigatorMapper.MapToDetailResponse(distance)
		if pr.Longitude != nil && pr.Latitude != nil {
			longitude := *pr.Longitude
			latitude := *pr.Latitude
			_, err = s.updateNavigator(ctx, pr.TelegramUserId, longitude, latitude)
			if err != nil {
				errorMessage := s.getErrorMessage("GetProfileDetail",
					"updateNavigator")
				s.logger.Debug(errorMessage, zap.Error(err))
				return nil, err
			}
		}
	}
	blockEntity, err := s.blockRepository.Find(ctx, pr.TelegramUserId, viewedTelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileDetail",
			"blockRepository.Find")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	blockMapper := mapper.BlockMapper{}
	blockResponse := blockMapper.MapToResponse(blockEntity)
	likeEntity, err := s.likeRepository.FindByTelegramUserId(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileDetail",
			"likeRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	likeMapper := &mapper.LikeMapper{}
	likeResponse := likeMapper.MapToResponse(likeEntity)
	telegramEntity, err := s.telegramRepository.FindByTelegramUserId(ctx, viewedTelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileDetail",
			"telegramRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	telegramMapper := &mapper.TelegramMapper{}
	telegramResponse := telegramMapper.MapToResponse(telegramEntity)
	isOnline := s.checkIsOnline(profileEntity.LastOnline)
	imageEntityList, err := s.imageRepository.SelectListByTelegramUserId(ctx, viewedTelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileDetail",
			"imageRepository.SelectListByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileResponse := profileMapper.MapToDetailResponse(
		profileEntity, navigatorResponse, blockResponse, likeResponse, telegramResponse, imageEntityList, isOnline)
	return profileResponse, err
}

func (s *ProfileService) GetProfileShortInfo(ctx context.Context, telegramUserId string,
	pr *request.ProfileGetShortInfoRequestDto) (*response.ProfileShortInfoResponseDto, error) {
	if err := s.checkUserExists(ctx, telegramUserId); err != nil {
		return nil, err
	}
	err := s.updateLastOnline(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileShortInfo", "updateLastOnline")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if pr.Longitude != nil && pr.Latitude != nil {
		longitude := *pr.Longitude
		latitude := *pr.Latitude
		_, err = s.updateNavigator(ctx, telegramUserId, longitude, latitude)
		if err != nil {
			errorMessage := s.getErrorMessage("GetProfileShortInfo", "updateNavigator")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	profileMapper := &mapper.ProfileMapper{}
	profileEntity, err := s.profileRepository.FindByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileShortInfo",
			"profileRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	lastImage, err := s.imageRepository.FindLastByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileShortInfo",
			"imageRepository.FindLastByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileResponse := profileMapper.MapToShortInfoResponse(profileEntity, lastImage.Url)
	return profileResponse, err
}

func (s *ProfileService) GetProfileList(ctx context.Context,
	pr *request.ProfileGetListRequestDto) (*response.ProfileListResponseDto, error) {
	if err := s.checkUserExists(ctx, pr.TelegramUserId); err != nil {
		return nil, err
	}
	err := s.updateLastOnline(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileList", "updateLastOnline")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if pr.Longitude != nil && pr.Latitude != nil {
		longitude := *pr.Longitude
		latitude := *pr.Latitude
		_, err = s.updateNavigator(ctx, pr.TelegramUserId, longitude, latitude)
		if err != nil {
			errorMessage := s.getErrorMessage("GetProfileList", "updateNavigator")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	filterEntity, err := s.filterRepository.FindByTelegramUserId(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileList",
			"filterRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToListRequest(filterEntity)
	var paginationProfileEntityList *response.ProfileListResponseRepositoryDto
	paginationProfileEntityList, err = s.profileRepository.SelectListByTelegramUserId(ctx, profileRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileList",
			"profileRepository.SelectListByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileContentResponse := make([]*response.ProfileListItemResponseDto, 0)
	if len(paginationProfileEntityList.Content) > 0 {
		for _, profileEntity := range paginationProfileEntityList.Content {
			lastImage, err := s.imageRepository.FindLastByTelegramUserId(ctx, profileEntity.TelegramUserId)
			if err != nil {
				errorMessage := s.getErrorMessage("GetProfileList",
					"imageRepository.FindLastByTelegramUserId")
				s.logger.Debug(errorMessage, zap.Error(err))
				return nil, err
			}
			lastOnline := profileEntity.LastOnline
			isOnline := s.checkIsOnline(lastOnline)
			profileItem := response.ProfileListItemResponseDto{
				TelegramUserId: profileEntity.TelegramUserId,
				Distance:       profileEntity.Distance,
				Url:            lastImage.Url,
				IsOnline:       isOnline,
				LastOnline:     profileEntity.LastOnline,
			}
			profileContentResponse = append(profileContentResponse, &profileItem)
		}
	}
	profileListResponse := &response.ProfileListResponseDto{
		PaginationEntity: paginationProfileEntityList.PaginationEntity,
		Content:          profileContentResponse,
	}
	return profileListResponse, err
}

func (s *ProfileService) AddImageList(
	ctx context.Context, unitOfWork *UnitOfWork, telegramUserId string, files []*entity.FileMetadata) error {
	if len(files) > 0 {
		for _, file := range files {
			_, err := s.AddImage(ctx, unitOfWork, telegramUserId, file)
			if err != nil {
				errorMessage := s.getErrorMessage("AddImageList", "AddImage")
				s.logger.Debug(errorMessage, zap.Error(err))
				return err
			}
		}
	}
	return nil
}

func (s *ProfileService) AddImage(ctx context.Context, unitOfWork *UnitOfWork, telegramUserId string,
	file *entity.FileMetadata) (*entity.ImageEntity, error) {
	imageConverted, err := s.uploadImageToFileSystem(ctx, file, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("AddImage", "uploadImageToFileSystem")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return unitOfWork.ImageRepository().Add(ctx, imageConverted)
}

func (s *ProfileService) GetImageByTelegramUserId(ctx context.Context, telegramUserId, fileName string) ([]byte, error) {
	filePath := fmt.Sprintf("static/profiles/%s/images/%s", telegramUserId, fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		errorMessage := s.getErrorMessage("GetImageByTelegramUserId", "IsNotExist")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		errorMessage := s.getErrorMessage("GetImageByTelegramUserId", "ReadFile")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return data, nil
}

func (s *ProfileService) GetImageById(ctx context.Context, imageId uint64) (*entity.ImageEntity, error) {
	return s.imageRepository.FindById(ctx, imageId)
}

func (s *ProfileService) updateImageList(
	ctx context.Context, unitOfWork *UnitOfWork, telegramUserId string, files []*entity.FileMetadata) error {
	return s.AddImageList(ctx, unitOfWork, telegramUserId, files)
}

func (s *ProfileService) deleteImageListByS3(ctx context.Context, unitOfWork *UnitOfWork, telegramUserId string) error {
	imageList, err := s.imageRepository.SelectListAllByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("deleteImageListByS3",
			"imageRepository.SelectListAllByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	if len(imageList) > 0 {
		for _, image := range imageList {
			_, err := s.deleteImageByS3(ctx, unitOfWork, image.Id)
			if err != nil {
				errorMessage := s.getErrorMessage("deleteImageListByS3",
					"deleteImageByS3")
				s.logger.Debug(errorMessage, zap.Error(err))
				return err
			}
		}
	}
	return nil
}

func (s *ProfileService) deleteImageByS3(
	ctx context.Context, unitOfWork *UnitOfWork, id uint64) (*response.ResponseDto, error) {
	image, err := s.imageRepository.FindById(ctx, id)
	if err != nil {
		errorMessage := s.getErrorMessage("deleteImageByS3",
			"imageRepository.FindById")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	// Удаление из файловой системы
	//filePath := image.Url
	//if err := os.Remove(filePath); err != nil {
	//	errorMessage := s.getErrorMessage("DeleteImage", "Remove")
	//	s.logger.Debug(errorMessage, zap.Error(err))
	//	return nil, err
	//}
	// Удаление из S3 хранилища
	pathToS3 := fmt.Sprintf("/profiles/%s/images/%s", image.TelegramUserId, image.Name)
	if err := s.s3.Delete(pathToS3); err != nil {
		errorMessage := s.getErrorMessage("deleteImageByS3", "s.s3.Delete")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageResponse := &response.ResponseDto{
		Success: true,
	}
	return imageResponse, nil
}
func (s *ProfileService) deleteImageByDB(
	ctx context.Context, unitOfWork *UnitOfWork, id uint64) (*response.ResponseDto, error) {
	imageMapper := &mapper.ImageMapper{}
	imageRequest := imageMapper.MapToDeleteRequest(id)
	return unitOfWork.ImageRepository().Delete(ctx, imageRequest)
}

func (s *ProfileService) DeleteImage(
	ctx context.Context, id uint64) (*response.ResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	_, err := s.deleteImageByS3(ctx, unitOfWork, id)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteImage", "deleteImageByS3")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	_, err = s.deleteImageByDB(ctx, unitOfWork, id)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteImage", "deleteImageByDB")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	responseDto := &response.ResponseDto{
		Success: true,
	}
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("DeleteImage", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("DeleteImage", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return responseDto, err
}

func (s *ProfileService) GetFilterByTelegramUserId(
	ctx context.Context, telegramUserId string, fr *request.FilterGetRequestDto) (*response.FilterResponseDto, error) {
	if err := s.checkUserExists(ctx, telegramUserId); err != nil {
		return nil, err
	}
	err := s.updateLastOnline(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetFilterByTelegramUserId",
			"updateLastOnline")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if fr.Longitude != nil && fr.Latitude != nil {
		longitude := *fr.Longitude
		latitude := *fr.Latitude
		_, err = s.updateNavigator(ctx, telegramUserId, longitude, latitude)
		if err != nil {
			errorMessage := s.getErrorMessage("GetFilterByTelegramUserId",
				"updateNavigator")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	filterEntity, err := s.filterRepository.FindByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetFilterByTelegramUserId",
			"filterRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	filterMapper := &mapper.FilterMapper{}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	return filterResponse, nil
}
func (s *ProfileService) UpdateFilter(
	ctx context.Context, fr *request.FilterUpdateRequestDto) (*response.FilterResponseDto, error) {
	tx, err := s.db.Begin()
	if err != nil {
		errorMessage := s.getErrorMessage("GetFilterByTelegramUserId", "Begin")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	err = s.updateLastOnline(ctx, fr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetFilterByTelegramUserId",
			"updateLastOnline")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	filterMapper := &mapper.FilterMapper{}
	filterRequest := filterMapper.MapToUpdateRequest(fr)
	filterEntity, err := s.filterRepository.Update(ctx, filterRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("GetFilterByTelegramUserId",
			"filterRepository.Update")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	tx.Commit()
	return filterResponse, nil
}

func (s *ProfileService) removeStrSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func (s *ProfileService) uploadImageToFileSystem(ctx context.Context, file *entity.FileMetadata,
	telegramUserId string) (*request.ImageAddRequestRepositoryDto, error) {
	directoryPath := fmt.Sprintf("static/profiles/%s/images", telegramUserId)
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		if err := os.MkdirAll(directoryPath, 0755); err != nil {
			errorMessage := s.getErrorMessage("uploadImageToFileSystem", "MkdirAll")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	filenameWithoutSpaces := s.removeStrSpaces(file.Filename)
	filePath := fmt.Sprintf("%s/%s", directoryPath, filenameWithoutSpaces)
	if err := os.WriteFile(filePath, file.Content, 0666); err != nil {
		errorMessage := s.getErrorMessage("uploadImageToFileSystem", "WriteFile")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	newFileName, newFilePath, newFileSize, err := s.convertImage(telegramUserId, directoryPath, filePath, filenameWithoutSpaces)
	if err != nil {
		errorMessage := s.getErrorMessage("uploadImageToFileSystem", "convertImage")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageConverted := &request.ImageAddRequestRepositoryDto{
		TelegramUserId: telegramUserId,
		Name:           newFileName,
		Url:            newFilePath,
		Size:           newFileSize,
		IsBlocked:      false,
		IsPrimary:      false,
		IsPrivate:      false,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
	return imageConverted, nil
}

func (s *ProfileService) convertImage(telegramUserId, directoryPath, filePath, fileName string) (string, string, int64, error) {
	newFileName := s.replaceFileName(fileName)
	newFilePathLocal := fmt.Sprintf("%s/%s", directoryPath, newFileName)
	output, err := os.Create(directoryPath + "/" + newFileName)
	if err != nil {
		errorMessage := s.getErrorMessage("convertImage", "Create")
		s.logger.Debug(errorMessage, zap.Error(err))
		return "", "", 0, err
	}
	defer output.Close()
	buffer, err := bimg.Read(filePath)
	if err != nil {
		errorMessage := s.getErrorMessage("convertImage", "Read")
		s.logger.Debug(errorMessage, zap.Error(err))
		return "", "", 0, err
	}
	newFile, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		errorMessage := s.getErrorMessage("convertImage", "NewImage")
		s.logger.Debug(errorMessage, zap.Error(err))
		return "", "", 0, err
	}
	bimg.Write(newFilePathLocal, newFile)
	newFileInfo, err := os.Stat(newFilePathLocal)
	if err != nil {
		errorMessage := s.getErrorMessage("convertImage", "Stat")
		s.logger.Debug(errorMessage, zap.Error(err))
		return "", "", 0, err
	}
	newFileSize := newFileInfo.Size()
	// S3 storage
	f, err := os.Open(newFilePathLocal)
	if err != nil {
		errorMessage := s.getErrorMessage("convertImage", "os.Open")
		s.logger.Debug(errorMessage, zap.Error(err))
		return "", "", 0, err
	}
	defer f.Close()
	pathToS3 := fmt.Sprintf("/profiles/%s/images/%s", telegramUserId, newFileName)
	result, err := s.s3.Upload(pathToS3, f)
	if err != nil {
		errorMessage := s.getErrorMessage("convertImage", "s.s3.Upload")
		s.logger.Debug(errorMessage, zap.Error(err))
		return "", "", 0, err
	}
	s.logger.Info(fmt.Sprintf("file uploaded to, %s", aws.StringValue(&result.Location)))
	newFilePath := fmt.Sprintf("%s%s", s.config.S3BucketPublicDomain, pathToS3)
	//
	if err := s.deleteFile(filePath); err != nil {
		errorMessage := s.getErrorMessage("convertImage", "deleteFile")
		s.logger.Debug(errorMessage, zap.Error(err))
		return "", "", 0, err
	}
	if err := os.RemoveAll("static/profiles"); err != nil {
		errorMessage := s.getErrorMessage("convertImage", "os.RemoveAll")
		s.logger.Debug(errorMessage, zap.Error(err))
		return "", "", 0, err
	}
	return newFileName, newFilePath, newFileSize, nil
}

func (s *ProfileService) deleteFile(filePath string) error {
	return os.Remove(filePath)
}

func (s *ProfileService) AddBlock(ctx context.Context, pr *request.BlockAddRequestDto) (*entity.BlockEntity, error) {
	tx, err := s.db.Begin()
	if err != nil {
		errorMessage := s.getErrorMessage("AddBlock", "Begin")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	blockMapper := &mapper.BlockMapper{}
	blockRequest := blockMapper.MapToAddRequest(pr)
	prForTwoUser := &request.BlockAddRequestDto{
		TelegramUserId:        pr.BlockedTelegramUserId,
		BlockedTelegramUserId: pr.TelegramUserId,
	}
	blockForTwoUserRequest := blockMapper.MapToAddRequest(prForTwoUser)
	_, err = s.blockRepository.Add(ctx, blockForTwoUserRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddBlock", "blockRepository.Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return s.blockRepository.Add(ctx, blockRequest)
}

func (s *ProfileService) AddLike(
	ctx context.Context, pr *request.LikeAddRequestDto, locale string) (*response.LikeResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	if err := s.checkUserExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("AddLike", "checkUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	telegramProfile, err := s.telegramRepository.FindByTelegramUserId(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("AddLike",
			"telegramRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	lastImageProfile, err := s.imageRepository.FindLastByTelegramUserId(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("AddLike",
			"imageRepository.FindLastByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	likedTelegramProfile, err := s.telegramRepository.FindByTelegramUserId(ctx, pr.LikedTelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("AddLike",
			"telegramRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	hc := &entity.HubContent{
		LikedTelegramUserId: likedTelegramProfile.UserId,
		Message:             s.GetMessageLike(locale),
		Type:                "like",
		UserImageUrl:        lastImageProfile.Url,
		Username:            telegramProfile.UserName,
	}
	hubContentJson, err := json.Marshal(hc)
	if err != nil {
		errorMessage := s.getErrorMessage("AddLike", "json.Marshal")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	err = s.kafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(likedTelegramProfile.UserId),
			Value: hubContentJson,
		},
	)
	if err != nil {
		errorMessage := s.getErrorMessage("AddLike", "kafkaWriter.WriteMessages")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	likeMapper := &mapper.LikeMapper{}
	likeRequest := likeMapper.MapToAddRequest(pr)
	likeAdded, err := unitOfWork.LikeRepository().Add(ctx, likeRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddLike",
			"LikeRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	likeResponse := likeMapper.MapToResponse(likeAdded)
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("AddLike", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("AddLike", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return likeResponse, nil
}

func (s *ProfileService) UpdateLike(
	ctx context.Context, pr *request.LikeUpdateRequestDto) (*response.LikeResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	if err := s.checkUserExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("UpdateLike", "checkUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	likeMapper := &mapper.LikeMapper{}
	likeRequest := likeMapper.MapToUpdateRequest(pr)
	likeUpdated, err := unitOfWork.LikeRepository().Update(ctx, likeRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateLike", "likeRepository.Update")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	likeResponse := likeMapper.MapToResponse(likeUpdated)
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("UpdateLike", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("UpdateLike", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return likeResponse, nil
}

func (s *ProfileService) AddComplaint(
	ctx context.Context, pr *request.ComplaintAddRequestDto) (*entity.ComplaintEntity, error) {
	unitOfWork := s.uwf.CreateUnit()
	complaintMapper := &mapper.ComplaintMapper{}
	complaintRequest := complaintMapper.MapToAddRequest(pr)
	complaintResponse, err := unitOfWork.ComplaintRepository().Add(ctx, complaintRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddComplaint",
			"ComplaintRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("AddComplaint", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("AddComplaint", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return complaintResponse, nil
}

func (s *ProfileService) UpdateCoordinates(
	ctx context.Context, pr *request.NavigatorUpdateRequestDto) (*response.NavigatorResponseDto, error) {
	longitude := pr.Longitude
	latitude := pr.Latitude
	return s.updateNavigator(ctx, pr.TelegramUserId, longitude, latitude)
}

func (s *ProfileService) updateLastOnline(ctx context.Context, telegramUserId string) error {
	tx, err := s.db.Begin()
	if err != nil {
		errorMessage := s.getErrorMessage("updateLastOnline", "Begin")
		s.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	defer tx.Rollback()
	updateLastOnlineMapper := &mapper.ProfileUpdateLastOnlineMapper{}
	updateLastOnlineRequest := updateLastOnlineMapper.MapToAddRequest(telegramUserId)
	err = s.profileRepository.UpdateLastOnline(ctx, updateLastOnlineRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("updateLastOnline",
			"profileRepository.UpdateLastOnline")
		s.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	tx.Commit()
	return nil
}

func (s *ProfileService) updateNavigator(
	ctx context.Context,
	telegramUserId string, longitude float64, latitude float64) (*response.NavigatorResponseDto, error) {
	tx, err := s.db.Begin()
	if err != nil {
		errorMessage := s.getErrorMessage("updateNavigator", "Begin")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	navigatorMapper := &mapper.NavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToUpdateRequest(telegramUserId, longitude, latitude)
	navigatorUpdated, err := s.navigatorRepository.Update(ctx, navigatorRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("updateNavigator",
			"navigatorRepository.Update")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	navigatorResponse := navigatorMapper.MapToResponse(telegramUserId, navigatorUpdated.Location.Longitude,
		navigatorUpdated.Location.Latitude)
	tx.Commit()
	return navigatorResponse, nil
}

func (s *ProfileService) getNowUtc() time.Time {
	return time.Now().UTC()
}

func (s *ProfileService) checkIsOnline(lastOnline time.Time) bool {
	now := s.getNowUtc()
	duration := now.Sub(lastOnline)
	minutes := duration.Minutes()
	return minutes < 5
}

func (s *ProfileService) replaceFileName(filename string) string {
	// Получаем текущее время
	now := time.Now().UTC()
	// Форматируем дату и время
	formattedTime := fmt.Sprintf("IMG_%04d%02d%02d_%02d%02d%02d_%03d", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1000000)
	// Добавляем расширение
	newFilename := formattedTime + filepath.Ext(filename)
	// Заменяем расширение, если необходимо
	return s.replaceExtension(newFilename)
}

func (s *ProfileService) replaceExtension(filename string) string {
	webpExtension := ".webp"
	if filepath.Ext(filename) == webpExtension {
		return filename
	}
	// Удаляем текущее расширение
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	// Добавляем новое расширение .webp
	return filename + webpExtension
}

func (s *ProfileService) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}

func (s *ProfileService) GetMessageLike(locale string) string {
	switch locale {
	case "ru":
		return "Есть симпатия! Начинай общаться"
	case "en":
		return "There is sympathy! Start communicating"
	default:
		return fmt.Sprintf("Unsupported language: %s", locale)
	}
}

func (s *ProfileService) checkUserExists(ctx context.Context, telegramUserId string) error {
	p, err := s.profileRepository.FindByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("checkUserExists",
			"profileRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	isFrozen := p.IsFrozen
	if isFrozen {
		err := errors.Wrap(err, "user has already been frozen")
		return err
	}
	return nil
}
