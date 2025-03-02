package service

import (
	"context"
	"database/sql"
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
	//"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	errorFilePath          = "internal/profiles/service/profile-service.go"
	maxCountUserComplaints = 1
)

type ProfileService struct {
	logger                logger.Logger
	db                    *sql.DB
	config                *config.Config
	hub                   *entity.Hub
	s3                    *config.S3
	uwf                   *UnitOfWorkFactory
	profileRepository     ProfileRepository
	navigatorRepository   NavigatorRepository
	filterRepository      FilterRepository
	telegramRepository    TelegramRepository
	imageRepository       ImageRepository
	imageStatusRepository ImageStatusRepository
	likeRepository        LikeRepository
	blockRepository       BlockRepository
	complaintRepository   ComplaintRepository
	statusRepository      StatusRepository
	paymentRepository     PaymentRepository
	settingsRepository    SettingsRepository
}

func NewProfileService(
	l logger.Logger,
	db *sql.DB,
	cfg *config.Config,
	h *entity.Hub,
	s3 *config.S3,
	uwf *UnitOfWorkFactory,
	pr ProfileRepository,
	nr NavigatorRepository,
	fr FilterRepository,
	tr TelegramRepository,
	ir ImageRepository,
	isr ImageStatusRepository,
	lr LikeRepository,
	br BlockRepository,
	cr ComplaintRepository,
	sr StatusRepository,
	pa PaymentRepository,
	str SettingsRepository) *ProfileService {
	return &ProfileService{
		logger:                l,
		db:                    db,
		config:                cfg,
		hub:                   h,
		s3:                    s3,
		uwf:                   uwf,
		profileRepository:     pr,
		navigatorRepository:   nr,
		telegramRepository:    tr,
		filterRepository:      fr,
		imageRepository:       ir,
		imageStatusRepository: isr,
		likeRepository:        lr,
		blockRepository:       br,
		complaintRepository:   cr,
		statusRepository:      sr,
		paymentRepository:     pa,
		settingsRepository:    str,
	}
}

func (s *ProfileService) AddProfile(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*response.ResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToAddRequest(pr)
	profileResponse, err := unitOfWork.ProfileRepository().Add(ctx, profileRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddProfile", "ProfileRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	statusMapper := &mapper.StatusMapper{}
	statusRequest := statusMapper.MapToAddRequest(pr)
	_, err = unitOfWork.StatusRepository().Add(ctx, statusRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddProfile",
			"StatusRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if pr.Longitude != nil && pr.Latitude != nil {
		longitude := *pr.Longitude
		latitude := *pr.Latitude
		navigatorMapper := &mapper.NavigatorMapper{}
		navigatorRequest := navigatorMapper.MapToAddRequest(pr.TelegramUserId, pr.CountryCode, pr.CountryName, pr.City,
			longitude, latitude)
		_, err = unitOfWork.NavigatorRepository().Add(ctx, navigatorRequest)
		if err != nil {
			errorMessage := s.getErrorMessage("AddProfile",
				"NavigatorRepository().Add")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	err = s.AddImageList(ctx, unitOfWork, pr.TelegramUserId, pr.Files)
	if err != nil {
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
	settingsMapper := &mapper.SettingsMapper{}
	settingsRequest := settingsMapper.MapToAddRequest(pr)
	_, err = unitOfWork.SettingsRepository().Add(ctx, settingsRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddProfile",
			"SettingsRepository().Add")
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
	if err := s.CheckProfileExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile", "CheckUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if err := s.updateImageList(ctx, unitOfWork, pr.TelegramUserId, pr.Files); err != nil {
		return nil, err
	}
	if pr.Longitude != nil && pr.Latitude != nil {
		longitude := *pr.Longitude
		latitude := *pr.Latitude
		navigatorMapper := &mapper.NavigatorMapper{}
		navigatorRequest := navigatorMapper.MapToUpdateRequest(pr.TelegramUserId, pr.CountryCode, pr.CountryName,
			pr.City, longitude, latitude)
		_, err := unitOfWork.NavigatorRepository().Update(ctx, navigatorRequest)
		if err != nil {
			errorMessage := s.getErrorMessage("UpdateProfile",
				"NavigatorRepository().Update")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	filterMapper := &mapper.FilterMapper{}
	filterRequest := filterMapper.MapProfileToUpdateRequest(pr)
	_, err := unitOfWork.FilterRepository().Update(ctx, filterRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile",
			"FilterRepository().Update")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	telegramMapper := &mapper.TelegramMapper{}
	telegramRequest := telegramMapper.MapToUpdateRequest(pr)
	_, err = unitOfWork.TelegramRepository().Update(ctx, telegramRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile",
			"telegramRepository.Update")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	settingsMapper := &mapper.SettingsMapper{}
	settingsRequest := settingsMapper.MapToUpdateRequest(pr)
	_, err = unitOfWork.SettingsRepository().Update(ctx, settingsRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile",
			"SettingsRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToUpdateRequest(pr)
	_, err = unitOfWork.ProfileRepository().Update(ctx, profileRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile",
			"ProfileRepository().Update")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileEntity, err := unitOfWork.ProfileRepository().GetProfile(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile",
			"ProfileRepository().Update")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageEntityList, err := unitOfWork.ImageRepository().SelectListByTelegramUserId(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile",
			"ImageRepository().SelectListByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	checkPremium, err := s.CheckPremium(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile", "s.CheckPremium")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileResponse := profileMapper.MapToResponse(profileEntity, imageEntityList, checkPremium.IsPremium)
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
	if err := s.CheckProfileExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("FreezeProfile", "CheckUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	updateLastOnlineMapper := &mapper.ProfileUpdateLastOnlineMapper{}
	updateLastOnlineRequest := updateLastOnlineMapper.MapToAddRequest(pr.TelegramUserId)
	err := unitOfWork.ProfileRepository().UpdateLastOnline(ctx, updateLastOnlineRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("FreezeProfile",
			"ProfileRepository().UpdateLastOnline")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	_, err = unitOfWork.StatusRepository().Freeze(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("FreezeProfile",
			"StatusRepository().Freeze")
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
	if err := s.CheckProfileExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("RestoreProfile", "CheckUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	_, err := unitOfWork.StatusRepository().Restore(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("RestoreProfile",
			"StatusRepository().Restore")
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
	if err := s.CheckProfileExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile", "CheckUserExists")
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
	_, err = unitOfWork.BlockRepository().DeleteRelatedProfiles(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile",
			"BlockRepository().DeleteRelatedProfiles")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	_, err = unitOfWork.ComplaintRepository().DeleteRelatedProfiles(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile",
			"ComplaintRepository().DeleteRelatedProfiles")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	_, err = unitOfWork.LikeRepository().DeleteRelatedProfiles(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile",
			"LikeRepository().DeleteRelatedProfiles")
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

func (s *ProfileService) GetProfile(ctx context.Context, telegramUserId string,
	pr *request.ProfileGetRequestDto) (*response.ProfileResponseDto, error) {
	if err := s.CheckProfileExists(ctx, telegramUserId); err != nil {
		return nil, err
	}
	if pr.Longitude != nil && pr.Latitude != nil {
		longitude := *pr.Longitude
		latitude := *pr.Latitude
		_, err := s.updateNavigator(ctx, telegramUserId, pr.CountryCode, pr.CountryName, pr.City, longitude, latitude)
		if err != nil {
			errorMessage := s.getErrorMessage("GetProfile", "updateNavigator")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	profileEntity, err := s.profileRepository.GetProfile(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfile",
			"profileRepository.GetProfile")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageList, err := s.imageRepository.SelectListByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfile",
			"imageRepository.SelectListByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	checkPremium, err := s.CheckPremium(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfile", "s.CheckPremium")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileResponse := profileMapper.MapToResponse(profileEntity, imageList, checkPremium.IsPremium)
	return profileResponse, err
}

func (s *ProfileService) GetProfileDetail(ctx context.Context, viewedTelegramUserId string,
	pr *request.ProfileGetDetailRequestDto) (*response.ProfileDetailResponseDto, error) {
	telegramUserId := pr.TelegramUserId
	if err := s.CheckProfileExists(ctx, telegramUserId); err != nil {
		return nil, err
	}
	if pr.Longitude != nil && pr.Latitude != nil {
		longitude := *pr.Longitude
		latitude := *pr.Latitude
		_, err := s.updateNavigator(ctx, telegramUserId, pr.CountryCode, pr.CountryName, pr.City, longitude, latitude)
		if err != nil {
			errorMessage := s.getErrorMessage("GetProfileDetail", "updateNavigator")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	profileDetail, err := s.profileRepository.GetDetail(ctx, telegramUserId, viewedTelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileDetail",
			"profileRepository.GetDetail")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageEntityList, err := s.imageRepository.SelectListByTelegramUserId(ctx, viewedTelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileDetail",
			"imageRepository.SelectListByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	checkPremium, err := s.CheckPremium(ctx, viewedTelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileDetail", "s.CheckPremium")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileResponse := profileMapper.MapToDetailResponse(profileDetail, imageEntityList, checkPremium.IsPremium)
	return profileResponse, err
}

func (s *ProfileService) GetProfileShortInfo(
	ctx context.Context, telegramUserId string) (*response.ProfileShortInfoResponseDto, error) {
	profileShortInfo, err := s.profileRepository.GetShortInfo(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileShortInfo",
			"profileRepository.GetShortInfo")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	checkPremium, err := s.CheckPremium(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileShortInfo", "s.CheckPremium")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileResponse := profileMapper.MapToShortInfoResponse(profileShortInfo, checkPremium)
	return profileResponse, err
}

func (s *ProfileService) GetProfileList(ctx context.Context,
	pr *request.ProfileGetListRequestDto) (*response.ProfileListResponseDto, error) {
	if err := s.CheckProfileExists(ctx, pr.TelegramUserId); err != nil {
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
		_, err = s.updateNavigator(ctx, pr.TelegramUserId, pr.CountryCode, pr.CountryName, pr.City, longitude, latitude)
		if err != nil {
			errorMessage := s.getErrorMessage("GetProfileList", "updateNavigator")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}

	var paginationProfileEntityList *response.ProfileListResponseRepositoryDto
	filterEntity, err := s.filterRepository.FindByTelegramUserId(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileList",
			"filterRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToListRequest(pr, filterEntity)
	paginationProfileEntityList, err = s.profileRepository.SelectList(ctx, profileRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("GetProfileList",
			"profileRepository.SelectList")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileListResponse := &response.ProfileListResponseDto{
		PaginationEntity: paginationProfileEntityList.PaginationEntity,
		Content:          paginationProfileEntityList.Content,
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
	file *entity.FileMetadata) (*response.ResponseDto, error) {
	imageConverted, err := s.uploadImageToFileSystem(ctx, file, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("AddImage", "uploadImageToFileSystem")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageId, err := unitOfWork.ImageRepository().Add(ctx, imageConverted)
	if err != nil {
		errorMessage := s.getErrorMessage("AddImage", "ImageRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageStatusMapper := &mapper.ImageStatusMapper{}
	imageStatusRequest := imageStatusMapper.MapToAddRequest(imageId)
	return unitOfWork.ImageStatusRepository().Add(ctx, imageStatusRequest)
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

func (s *ProfileService) GetImageLastByTelegramUserId(
	ctx context.Context, telegramUserId string) (*response.ImageResponseDto, error) {
	return s.imageRepository.FindLastByTelegramUserId(ctx, telegramUserId)
}

func (s *ProfileService) GetImageById(ctx context.Context, imageId uint64) (*response.ImageResponseDto, error) {
	image, err := s.imageRepository.FindById(ctx, imageId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetImageById", "imageRepository.FindById")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageMapper := &mapper.ImageMapper{}
	imageResponse := imageMapper.MapToResponse(image)
	return imageResponse, nil
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
	return unitOfWork.ImageRepository().Delete(ctx, id)
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

func (s *ProfileService) GetFilter(ctx context.Context, telegramUserId string) (*response.FilterResponseDto, error) {
	filterEntity, err := s.filterRepository.FindByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetFilter",
			"filterRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	filterMapper := &mapper.FilterMapper{}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	return filterResponse, nil
}

func (s *ProfileService) UpdateFilter(
	ctx context.Context, req *request.FilterUpdateRequestDto) (*response.FilterResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	filterMapper := &mapper.FilterMapper{}
	filterRequest := filterMapper.MapToUpdateRequest(req)
	filterEntity, err := unitOfWork.FilterRepository().Update(ctx, filterRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateFilter", "filterRepository.Update")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("UpdateFilter", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("UpdateFilter", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return filterResponse, err
}

func (s *ProfileService) GetTelegram(
	ctx context.Context, telegramUserId string) (*response.TelegramResponseDto, error) {
	telegramEntity, err := s.telegramRepository.FindByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetTelegram",
			"telegramRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	telegramMapper := &mapper.TelegramMapper{}
	return telegramMapper.MapToResponse(telegramEntity), nil
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

func (s *ProfileService) AddBlock(ctx context.Context, pr *request.BlockAddRequestDto) (*response.ResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	if err := s.CheckProfileExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("AddBlock", "CheckUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	blockExists, err := unitOfWork.BlockRepository().FindBlock(ctx, pr.TelegramUserId, pr.BlockedTelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("AddBlock", "BlockRepository().FindBlock")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	blockMapper := &mapper.BlockMapper{}
	if blockExists != nil {
		blockRequest := blockMapper.MapToUpdateRequest(pr, &pr.TelegramUserId)
		_, err = unitOfWork.BlockRepository().Update(ctx, blockRequest)
		if err != nil {
			errorMessage := s.getErrorMessage("AddBlock", "BlockRepository().Update")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
		prForViewedUser := &request.BlockAddRequestDto{
			TelegramUserId:        pr.BlockedTelegramUserId,
			BlockedTelegramUserId: pr.TelegramUserId,
		}
		blockForViewedUserRequest := blockMapper.MapToUpdateRequest(prForViewedUser, nil)
		_, err = unitOfWork.BlockRepository().Update(ctx, blockForViewedUserRequest)
		if err != nil {
			errorMessage := s.getErrorMessage("AddBlock", "BlockRepository().Update")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	if blockExists == nil {
		blockRequest := blockMapper.MapToAddRequest(pr, &pr.TelegramUserId)
		_, err = unitOfWork.BlockRepository().Add(ctx, blockRequest)
		if err != nil {
			errorMessage := s.getErrorMessage("AddBlock", "BlockRepository().Add")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
		prForViewedUser := &request.BlockAddRequestDto{
			TelegramUserId:        pr.BlockedTelegramUserId,
			BlockedTelegramUserId: pr.TelegramUserId,
		}
		blockForViewedUserRequest := blockMapper.MapToAddRequest(prForViewedUser, nil)
		_, err := unitOfWork.BlockRepository().Add(ctx, blockForViewedUserRequest)
		if err != nil {
			errorMessage := s.getErrorMessage("AddBlock", "BlockRepository().Add")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("AddBlock", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("AddBlock", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	blockResponse := &response.ResponseDto{
		Success: true,
	}
	return blockResponse, nil
}

func (s *ProfileService) GetBlockedList(ctx context.Context,
	telegramUserId string) (*response.BlockedListResponseDto, error) {
	return s.blockRepository.GetBlockedList(ctx, telegramUserId)
}

func (s *ProfileService) Unblock(ctx context.Context, p *request.UnblockRequestDto) (*response.ResponseDto, error) {
	telegramUserId := p.TelegramUserId
	blockedTelegramUserId := p.BlockedTelegramUserId
	unitOfWork := s.uwf.CreateUnit()
	if err := s.CheckProfileExists(ctx, telegramUserId); err != nil {
		errorMessage := s.getErrorMessage("Unblock", "CheckUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	_, err := unitOfWork.BlockRepository().Unblock(ctx, p)
	if err != nil {
		errorMessage := s.getErrorMessage("Unblock", "BlockRepository().Unblock")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	revert := &request.UnblockRequestDto{
		TelegramUserId:        blockedTelegramUserId,
		BlockedTelegramUserId: telegramUserId,
	}
	_, err = unitOfWork.BlockRepository().Unblock(ctx, revert)
	if err != nil {
		errorMessage := s.getErrorMessage("Unblock", "BlockRepository().Unblock")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("Unblock", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("Unblock", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	blockResponse := &response.ResponseDto{
		Success: true,
	}
	return blockResponse, nil
}

func (s *ProfileService) AddLike(
	ctx context.Context, pr *request.LikeAddRequestDto, locale string) (*response.ResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	if err := s.CheckProfileExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("AddLike", "CheckUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	// For channel
	//telegramProfile, err := s.telegramRepository.FindByTelegramUserId(ctx, pr.TelegramUserId)
	//if err != nil {
	//	errorMessage := s.getErrorMessage("AddLike",
	//		"telegramRepository.FindByTelegramUserId")
	//	s.logger.Debug(errorMessage, zap.Error(err))
	//	return nil, err
	//}
	//statusProfile, err := s.statusRepository.FindByTelegramUserId(ctx, pr.TelegramUserId)
	//if err != nil {
	//	errorMessage := s.getErrorMessage("AddLike",
	//		"statusRepository.FindByTelegramUserId")
	//	s.logger.Debug(errorMessage, zap.Error(err))
	//	return nil, err
	//}
	//lastImageProfile, err := s.imageRepository.FindLastByTelegramUserId(ctx, pr.TelegramUserId)
	//if err != nil {
	//	errorMessage := s.getErrorMessage("AddLike",
	//		"imageRepository.FindLastByTelegramUserId")
	//	s.logger.Debug(errorMessage, zap.Error(err))
	//	return nil, err
	//}
	//likedTelegramProfile, err := s.telegramRepository.FindByTelegramUserId(ctx, pr.LikedTelegramUserId)
	//if err != nil {
	//	errorMessage := s.getErrorMessage("AddLike",
	//		"telegramRepository.FindByTelegramUserId")
	//	s.logger.Debug(errorMessage, zap.Error(err))
	//	return nil, err
	//}
	//hc := &entity.HubContent{
	//	LikedTelegramUserId: likedTelegramProfile.UserId,
	//	Message:             s.GetMessageLike(locale),
	//	Type:                "like",
	//	UserImageUrl:        lastImageProfile.Url,
	//	Username:            telegramProfile.UserName,
	//}
	//if !statusProfile.IsBlocked {
	// For channel
	//go func() {
	//	s.hub.Broadcast <- hc
	//}()
	likeMapper := &mapper.LikeMapper{}
	likeRequest := likeMapper.MapToAddRequest(pr)
	likeResponse, err := unitOfWork.LikeRepository().Add(ctx, likeRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddLike",
			"LikeRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
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
	ctx context.Context, pr *request.LikeUpdateRequestDto) (*response.ResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	if err := s.CheckProfileExists(ctx, pr.TelegramUserId); err != nil {
		errorMessage := s.getErrorMessage("UpdateLike", "checkUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	likeMapper := &mapper.LikeMapper{}
	likeRequest := likeMapper.MapToUpdateRequest(pr)
	likeResponse, err := unitOfWork.LikeRepository().Update(ctx, likeRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateLike", "likeRepository.Update")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
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

func (s *ProfileService) GetLastLike(
	ctx context.Context, telegramUserId string) (*entity.LikeEntity, error) {
	return s.likeRepository.FindLastLike(ctx, telegramUserId)
}

func (s *ProfileService) AddComplaint(
	ctx context.Context, pr *request.ComplaintAddRequestDto) (*response.ResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	blockMapper := &mapper.BlockMapper{}
	br := &request.BlockAddRequestDto{
		TelegramUserId:        pr.TelegramUserId,
		BlockedTelegramUserId: pr.CriminalTelegramUserId,
	}
	blockRequest := blockMapper.MapToAddRequest(br, &br.TelegramUserId)
	_, err := unitOfWork.BlockRepository().Add(ctx, blockRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddComplaint", "BlockRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	prForViewedUser := &request.BlockAddRequestDto{
		TelegramUserId:        pr.CriminalTelegramUserId,
		BlockedTelegramUserId: pr.TelegramUserId,
	}
	blockForViewedUserRequest := blockMapper.MapToAddRequest(prForViewedUser, nil)
	_, err = unitOfWork.BlockRepository().Add(ctx, blockForViewedUserRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddComplaint", "BlockRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	complaintMapper := &mapper.ComplaintMapper{}
	complaintRequest := complaintMapper.MapToAddRequest(pr)
	complaintResponse, err := unitOfWork.ComplaintRepository().Add(ctx, complaintRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddComplaint",
			"ComplaintRepository().Add")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	countUserComplaints, err := unitOfWork.ComplaintRepository().GetCountUserComplaintsByCurrentMonth(
		ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("AddComplaint",
			"complaintRepository.GetCountUserComplaintsByToday")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if countUserComplaints >= maxCountUserComplaints {
		_, err := unitOfWork.StatusRepository().Block(ctx, pr.CriminalTelegramUserId)
		if err != nil {
			errorMessage := s.getErrorMessage("AddComplaint",
				"statusRepository().Block")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
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

func (s *ProfileService) GetStatusByTelegramUserId(
	ctx context.Context, telegramUserId string) (*response.StatusResponseDto, error) {
	checkPremium, err := s.CheckPremium(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetStatusByTelegramUserId",
			"s.CheckPremium")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	statusEntity, err := s.statusRepository.FindByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("GetStatusByTelegramUserId",
			"s.statusRepository.FindByTelegramUserId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	statusMapper := &mapper.StatusMapper{}
	statusResponse := statusMapper.MapToResponse(statusEntity, checkPremium.IsPremium)
	return statusResponse, nil
}

func (s *ProfileService) AddPayment(
	ctx context.Context, pr *request.PaymentAddRequestDto) (*response.ResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	paymentLast, err := unitOfWork.PaymentRepository().FindLastByTelegramUserId(ctx, pr.TelegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("AddPayment",
			"unitOfWork.PaymentRepository().Add()")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	paymentMapper := &mapper.PaymentMapper{}
	paymentRequest := paymentMapper.MapToAddRequest(pr, paymentLast)
	paymentResponse, err := unitOfWork.PaymentRepository().Add(ctx, paymentRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddPayment",
			"unitOfWork.PaymentRepository().Add()")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("AddPayment", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("AddPayment", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return paymentResponse, nil
}

func (s *ProfileService) GetPaymentLastByTelegramUserId(
	ctx context.Context, telegramUserId string) (*entity.PaymentEntity, error) {
	return s.paymentRepository.FindLastByTelegramUserId(ctx, telegramUserId)
}

func (s *ProfileService) CheckPremium(
	ctx context.Context, telegramUserId string) (*response.PremiumResponseDto, error) {
	paymentEntity, err := s.paymentRepository.FindLastByTelegramUserId(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("CheckPremium",
			"paymentRepository.FindLastByTelegramUserId()")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	paymentMapper := &mapper.PaymentMapper{}
	paymentResponse := paymentMapper.MapToCheckPremium(paymentEntity)
	return paymentResponse, nil
}

func (s *ProfileService) UpdateSettings(
	ctx context.Context, pr *request.ProfileUpdateSettingsRequestDto) (*response.ResponseDto, error) {
	unitOfWork := s.uwf.CreateUnit()
	statusMapper := &mapper.StatusMapper{}
	statusRequest := statusMapper.MapToUpdateSettingsRequest(pr)
	statusResponse, err := unitOfWork.StatusRepository().UpdateSettings(ctx, statusRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateSettings",
			"unitOfWork.StatusRepository().Add()")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer func() {
		if err != nil {
			if err := unitOfWork.Rollback(ctx); err != nil {
				errorMessage := s.getErrorMessage("UpdateSettings", "Rollback")
				s.logger.Debug(errorMessage, zap.Error(err))
			}
		}
	}()
	if err = unitOfWork.Commit(ctx); err != nil {
		errorMessage := s.getErrorMessage("UpdateSettings", "Commit")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return statusResponse, nil
}

func (s *ProfileService) UpdateCoordinates(
	ctx context.Context, pr *request.NavigatorUpdateRequestDto) (*response.ResponseDto, error) {
	longitude := pr.Longitude
	latitude := pr.Latitude
	return s.updateNavigator(ctx, pr.TelegramUserId, pr.CountryCode, pr.CountryName, pr.City, longitude, latitude)
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

func (s *ProfileService) updateNavigator(ctx context.Context, telegramUserId string, countryCode, countryName,
	city *string, longitude float64, latitude float64) (*response.ResponseDto, error) {
	tx, err := s.db.Begin()
	if err != nil {
		errorMessage := s.getErrorMessage("updateNavigator", "Begin")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	navigatorMapper := &mapper.NavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToUpdateRequest(telegramUserId, countryCode, countryName, city, longitude,
		latitude)
	navigatorExists, err := s.checkNavigatorExists(ctx, telegramUserId)
	if navigatorExists != nil {
		if countryCode != nil && countryName != nil && city != nil {
			_, err := s.navigatorRepository.Update(ctx, navigatorRequest)
			if err != nil {
				errorMessage := s.getErrorMessage("updateNavigator",
					"navigatorRepository.Update")
				s.logger.Debug(errorMessage, zap.Error(err))
				return nil, err
			}
		} else {
			_, err := s.navigatorRepository.UpdateCoordinates(ctx, navigatorRequest)
			if err != nil {
				errorMessage := s.getErrorMessage("updateNavigator",
					"navigatorRepository.UpdateCoordinates")
				s.logger.Debug(errorMessage, zap.Error(err))
				return nil, err
			}
		}
	} else {
		navigatorRequest := navigatorMapper.MapToAddRequest(telegramUserId, countryCode, countryName, city, longitude,
			latitude)
		_, err = s.navigatorRepository.Add(ctx, navigatorRequest)
		if err != nil {
			errorMessage := s.getErrorMessage("updateNavigator",
				"navigatorRepository.Add")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	tx.Commit()
	navigatorResponse := &response.ResponseDto{
		Success: true,
	}
	return navigatorResponse, nil
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

func (s *ProfileService) CheckProfileExists(ctx context.Context, telegramUserId string) error {
	p, err := s.statusRepository.CheckProfileExists(ctx, telegramUserId)
	if err != nil {
		errorMessage := s.getErrorMessage("CheckProfileExists",
			"statusRepository.CheckProfileExists")
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

func (s *ProfileService) checkNavigatorExists(
	ctx context.Context, telegramUserId string) (*response.ResponseDto, error) {
	return s.navigatorRepository.CheckNavigatorExists(ctx, telegramUserId)
}

func (s *ProfileService) GetMessageLike(locale string) string {
	switch locale {
	case "ru":
		return "Есть симпатия! Начинай общаться"
	case "en":
		return "There is sympathy! Start communicating"
	case "ar":
		return "هناك تعاطف! ابدأ التواصل"
	case "be":
		return "Ёсць сімпатыя! Пачынай мець зносіны"
	case "ca":
		return "Hi ha simpatia! Comença a comunicar-te"
	case "cs":
		return "Jsou tam sympatie! Začněte komunikovat"
	case "de":
		return "Es gibt Mitgefühl! Beginnen Sie mit der Kommunikation"
	case "es":
		return "¡Hay simpatía! Empezar a comunicar"
	case "fi":
		return "Sympatiaa on! Aloita kommunikointi"
	case "fr":
		return "Il y a de la sympathie ! Commencez à communiquer"
	case "he":
		return "יש סימפטיה! תתחיל לתקשר"
	case "hi":
		return "सहानुभूति है! संवाद करना शुरू करें"
	case "hr":
		return "Postoji simpatija! Počnite komunicirati"
	case "hu":
		return "Van együttérzés! Kezdj el kommunikálni"
	case "id":
		return "Ada simpati! Mulailah berkomunikasi"
	case "it":
		return "C'è simpatia! Inizia a comunicare"
	case "ja":
		return "共感があるよ！通信を開始する"
	case "kk":
		return "Жанашырлық бар! Қарым-қатынасты бастаңызe"
	case "ko":
		return "동정심이 있습니다! 소통을 시작해 보세요"
	case "nl":
		return "Er is sympathie! Begin met communiceren"
	case "no":
		return "Det er sympati! Begynn å kommunisere"
	case "pt":
		return "Existe simpatia! Comece a se comunicar"
	case "sv":
		return "Det finns sympati! Börja kommunicera"
	case "uk":
		return "Є симпатія! Починай спілкуватися"
	case "zh":
		return "有同情心！开始沟通"
	default:
		return "There is sympathy! Start communicating"
	}
}
