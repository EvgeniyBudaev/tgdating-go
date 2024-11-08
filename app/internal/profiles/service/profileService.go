package service

import (
	"context"
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
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	errorFilePath = "internal/profiles/service/profileService.go"
	minDistance   = 100
)

type ProfileService struct {
	logger              logger.Logger
	config              *config.Config
	hub                 *entity.Hub
	s3                  *config.S3
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
	cfg *config.Config,
	h *entity.Hub,
	s3 *config.S3,
	pr ProfileRepository,
	nr NavigatorRepository,
	fr FilterRepository,
	tr TelegramRepository,
	ir ImageRepository, lr LikeRepository, br BlockRepository, cr ComplaintRepository) *ProfileService {
	return &ProfileService{
		logger:              l,
		config:              cfg,
		hub:                 h,
		s3:                  s3,
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
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToAddRequest(pr)
	profileCreated, err := s.profileRepository.Add(ctx, profileRequest)
	if err != nil {
		return nil, err
	}
	_, err = s.AddNavigator(ctx, pr)
	if err != nil {
		return nil, err
	}
	profileResponse := profileMapper.MapToAddResponse(profileCreated)
	if err := s.AddImageList(ctx, profileCreated.SessionId, pr.Files); err != nil {
		return nil, err
	}
	_, err = s.AddFilter(ctx, pr)
	if err != nil {
		return nil, err
	}
	_, err = s.AddTelegram(ctx, pr)
	if err != nil {
		return nil, err
	}
	return profileResponse, err
}

func (s *ProfileService) UpdateProfile(
	ctx context.Context, pr *request.ProfileUpdateRequestDto) (*response.ProfileResponseDto, error) {
	sessionId := pr.SessionId
	if err := s.checkUserExists(ctx, sessionId); err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToUpdateRequest(pr)
	profileEntity, err := s.profileRepository.Update(ctx, profileRequest)
	if err != nil {
		return nil, err
	}
	if err := s.UpdateImageList(ctx, profileEntity.SessionId, pr.Files); err != nil {
		return nil, err
	}
	var navigatorResponse *response.NavigatorResponseDto
	if pr.Longitude != 0 && pr.Latitude != 0 {
		navigatorResponse, err = s.updateNavigator(ctx, sessionId, pr.Longitude, pr.Latitude)
		if err != nil {
			return nil, err
		}
	}
	filterMapper := &mapper.FilterMapper{}
	filterRequest := filterMapper.MapProfileToUpdateRequest(pr)
	filterEntity, err := s.filterRepository.Update(ctx, filterRequest)
	if err != nil {
		return nil, err
	}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	telegramMapper := &mapper.TelegramMapper{}
	telegramRequest := telegramMapper.MapToUpdateRequest(pr)
	telegramEntity, err := s.telegramRepository.Update(ctx, telegramRequest)
	if err != nil {
		return nil, err
	}
	telegramResponse := telegramMapper.MapToResponse(telegramEntity)
	imageEntityList, err := s.imageRepository.SelectListBySessionId(ctx, sessionId)
	isOnline := s.checkIsOnline(profileEntity.LastOnline)
	profileResponse := profileMapper.MapToResponse(profileEntity, navigatorResponse, filterResponse, telegramResponse,
		imageEntityList, isOnline)
	return profileResponse, nil
}

func (s *ProfileService) DeleteProfile(
	ctx context.Context, pr *request.ProfileDeleteRequestDto) (*response.ResponseDto, error) {
	sessionId := pr.SessionId
	if err := s.checkUserExists(ctx, sessionId); err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToDeleteRequest(sessionId)
	_, err := s.profileRepository.Delete(ctx, profileRequest)
	if err != nil {
		return nil, err
	}
	navigatorMapper := &mapper.NavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToDeleteRequest(sessionId)
	_, err = s.navigatorRepository.Delete(ctx, navigatorRequest)
	if err != nil {
		return nil, err
	}
	err = s.DeleteImageList(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	filterMapper := &mapper.FilterMapper{}
	filterRequest := filterMapper.MapToDeleteRequest(sessionId)
	_, err = s.filterRepository.Delete(ctx, filterRequest)
	if err != nil {
		return nil, err
	}
	telegramMapper := &mapper.TelegramMapper{}
	telegramRequest := telegramMapper.MapToDeleteRequest(sessionId)
	_, err = s.telegramRepository.Delete(ctx, telegramRequest)
	if err != nil {
		return nil, err
	}
	profileResponse := &response.ResponseDto{
		Success: true,
	}
	return profileResponse, err
}

func (s *ProfileService) GetProfileBySessionId(ctx context.Context, sessionId string,
	pr *request.ProfileGetBySessionIdRequestDto) (*response.ProfileResponseDto, error) {
	if err := s.checkUserExists(ctx, sessionId); err != nil {
		return nil, err
	}
	err := s.updateLastOnline(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	if pr.Longitude != 0 && pr.Latitude != 0 {
		_, err = s.updateNavigator(ctx, sessionId, pr.Longitude, pr.Latitude)
		if err != nil {
			return nil, err
		}
	}
	profileMapper := &mapper.ProfileMapper{}
	profileEntity, err := s.profileRepository.FindBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	navigatorMapper := &mapper.NavigatorMapper{}
	navigatorEntity, err := s.navigatorRepository.FindBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	longitude := navigatorEntity.Location.Longitude
	latitude := navigatorEntity.Location.Latitude
	navigatorResponse := navigatorMapper.MapToResponse(sessionId, longitude, latitude)
	filterMapper := &mapper.FilterMapper{}
	filterEntity, err := s.filterRepository.FindBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	telegramEntity, err := s.telegramRepository.FindBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	telegramMapper := &mapper.TelegramMapper{}
	telegramResponse := telegramMapper.MapToResponse(telegramEntity)
	isOnline := s.checkIsOnline(profileEntity.LastOnline)
	imageEntityList, err := s.imageRepository.SelectListBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	profileResponse := profileMapper.MapToResponse(
		profileEntity, navigatorResponse, filterResponse, telegramResponse, imageEntityList, isOnline)
	return profileResponse, err
}

func (s *ProfileService) GetProfileDetail(ctx context.Context, viewedSessionId string,
	pr *request.ProfileGetDetailRequestDto) (*response.ProfileDetailResponseDto, error) {
	sessionId := pr.SessionId
	if err := s.checkUserExists(ctx, sessionId); err != nil {
		return nil, err
	}
	err := s.updateLastOnline(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	if pr.Longitude != 0 && pr.Latitude != 0 {
		_, err = s.updateNavigator(ctx, sessionId, pr.Longitude, pr.Latitude)
		if err != nil {
			return nil, err
		}
	}
	profileMapper := &mapper.ProfileMapper{}
	profileEntity, err := s.profileRepository.FindBySessionId(ctx, viewedSessionId)
	if err != nil {
		return nil, err
	}
	navigatorMapper := &mapper.NavigatorMapper{}
	navigatorEntity, err := s.navigatorRepository.FindBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	navigatorViewedEntity, err := s.navigatorRepository.FindBySessionId(ctx, viewedSessionId)
	if err != nil {
		return nil, err
	}
	navigatorDistanceResponse, err := s.navigatorRepository.FindDistance(ctx, navigatorEntity, navigatorViewedEntity)
	if err != nil {
		return nil, err
	}
	distance := navigatorDistanceResponse.Distance
	if distance < minDistance {
		distance = minDistance
	}
	navigatorResponse := navigatorMapper.MapToDetailResponse(distance)
	if pr.Longitude != 0 && pr.Latitude != 0 {
		_, err = s.updateNavigator(ctx, sessionId, pr.Longitude, pr.Latitude)
		if err != nil {
			return nil, err
		}
	}
	blockEntity, err := s.blockRepository.Find(ctx, sessionId, viewedSessionId)
	if err != nil {
		return nil, err
	}
	blockMapper := mapper.BlockMapper{}
	blockResponse := blockMapper.MapToResponse(blockEntity)
	likeEntity, err := s.likeRepository.FindBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	likeMapper := &mapper.LikeMapper{}
	likeResponse := likeMapper.MapToResponse(likeEntity)
	telegramEntity, err := s.telegramRepository.FindBySessionId(ctx, viewedSessionId)
	if err != nil {
		return nil, err
	}
	telegramMapper := &mapper.TelegramMapper{}
	telegramResponse := telegramMapper.MapToResponse(telegramEntity)
	isOnline := s.checkIsOnline(profileEntity.LastOnline)
	imageEntityList, err := s.imageRepository.SelectListBySessionId(ctx, viewedSessionId)
	if err != nil {
		return nil, err
	}
	profileResponse := profileMapper.MapToDetailResponse(
		profileEntity, navigatorResponse, blockResponse, likeResponse, telegramResponse, imageEntityList, isOnline)
	return profileResponse, err
}

func (s *ProfileService) GetProfileShortInfo(ctx context.Context, sessionId string,
	pr *request.ProfileGetShortInfoRequestDto) (*response.ProfileShortInfoResponseDto, error) {
	if err := s.checkUserExists(ctx, sessionId); err != nil {
		return nil, err
	}
	err := s.updateLastOnline(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	if pr.Longitude != 0 && pr.Latitude != 0 {
		_, err = s.updateNavigator(ctx, sessionId, pr.Longitude, pr.Latitude)
		if err != nil {
			return nil, err
		}
	}
	profileMapper := &mapper.ProfileMapper{}
	profileEntity, err := s.profileRepository.FindBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	lastImage, err := s.imageRepository.FindLastBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	profileResponse := profileMapper.MapToShortInfoResponse(profileEntity, lastImage.Url)
	return profileResponse, err
}

func (s *ProfileService) GetProfileList(ctx context.Context,
	pr *request.ProfileGetListRequestDto) (*response.ProfileListResponseDto, error) {
	sessionId := pr.SessionId
	if err := s.checkUserExists(ctx, sessionId); err != nil {
		return nil, err
	}
	err := s.updateLastOnline(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	if pr.Longitude != 0 && pr.Latitude != 0 {
		_, err = s.updateNavigator(ctx, sessionId, pr.Longitude, pr.Latitude)
		if err != nil {
			return nil, err
		}
	}
	filterEntity, err := s.filterRepository.FindBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToListRequest(filterEntity)
	paginationProfileEntityList, err := s.profileRepository.SelectListBySessionId(ctx, profileRequest)
	if err != nil {
		return nil, err
	}
	profileContentResponse := make([]*response.ProfileListItemResponseDto, 0)
	if len(paginationProfileEntityList.Content) > 0 {
		for _, profileEntity := range paginationProfileEntityList.Content {
			lastImage, err := s.imageRepository.FindLastBySessionId(ctx, profileEntity.SessionId)
			if err != nil {
				return nil, err
			}
			url := lastImage.Url
			lastOnline := profileEntity.LastOnline
			isOnline := s.checkIsOnline(lastOnline)
			distance := profileEntity.Distance
			profileItem := response.ProfileListItemResponseDto{
				SessionId:  profileEntity.SessionId,
				Distance:   distance,
				Url:        url,
				IsOnline:   isOnline,
				LastOnline: lastOnline,
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

func (s *ProfileService) AddImageList(ctx context.Context, sessionId string, files []*entity.FileMetadata) error {
	if len(files) > 0 {
		for _, file := range files {
			_, err := s.AddImage(ctx, sessionId, file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *ProfileService) AddImage(ctx context.Context, sessionId string,
	file *entity.FileMetadata) (*entity.ImageEntity, error) {
	imageConverted, err := s.uploadImageToFileSystem(ctx, file, sessionId)
	if err != nil {
		errorMessage := s.getErrorMessage("AddImage", "uploadImageToFileSystem")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return s.imageRepository.Add(ctx, imageConverted)
}

func (s *ProfileService) UpdateImageList(ctx context.Context, sessionId string, files []*entity.FileMetadata) error {
	return s.AddImageList(ctx, sessionId, files)
}

func (s *ProfileService) DeleteImageList(ctx context.Context, sessionId string) error {
	imageList, err := s.imageRepository.SelectListBySessionId(ctx, sessionId)
	if err != nil {
		return err
	}
	if len(imageList) > 0 {
		for _, image := range imageList {
			_, err := s.deleteImageById(ctx, image.Id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *ProfileService) GetImageBySessionId(ctx context.Context, sessionId, fileName string) ([]byte, error) {
	filePath := fmt.Sprintf("static/profiles/%s/images/%s", sessionId, fileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		errorMessage := s.getErrorMessage("GetImageBySessionId", "IsNotExist")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		errorMessage := s.getErrorMessage("GetImageBySessionId", "ReadFile")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return data, nil
}

func (s *ProfileService) GetImageById(ctx context.Context, imageId uint64) (*entity.ImageEntity, error) {
	return s.imageRepository.FindById(ctx, imageId)
}

func (s *ProfileService) deleteImageById(ctx context.Context, id uint64) (*entity.ImageEntity, error) {
	image, err := s.imageRepository.FindById(ctx, id)
	if err != nil {
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
	pathToS3 := fmt.Sprintf("/profiles/%s/images/%s", image.SessionId, image.Name)
	if err := s.s3.Delete(pathToS3); err != nil {
		errorMessage := s.getErrorMessage("deleteImageById", "s.s3.Delete")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageMapper := &mapper.ImageMapper{}
	imageRequest := imageMapper.MapToDeleteRequest(image.Id)
	return s.imageRepository.Delete(ctx, imageRequest)
}

func (s *ProfileService) DeleteImage(ctx context.Context, id uint64) (*response.ResponseDto, error) {
	_, err := s.deleteImageById(ctx, id)
	if err != nil {
		return nil, err
	}
	responseDto := &response.ResponseDto{
		Success: true,
	}
	return responseDto, err
}

func (s *ProfileService) GetFilterBySessionId(
	ctx context.Context, sessionId string, fr *request.FilterGetRequestDto) (*response.FilterResponseDto, error) {
	if err := s.checkUserExists(ctx, sessionId); err != nil {
		return nil, err
	}
	err := s.updateLastOnline(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	if fr.Longitude != 0 && fr.Latitude != 0 {
		_, err = s.updateNavigator(ctx, sessionId, fr.Longitude, fr.Latitude)
		if err != nil {
			return nil, err
		}
	}
	filterEntity, err := s.filterRepository.FindBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	filterMapper := &mapper.FilterMapper{}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	return filterResponse, nil
}
func (s *ProfileService) UpdateFilter(
	ctx context.Context, fr *request.FilterUpdateRequestDto) (*response.FilterResponseDto, error) {
	sessionId := fr.SessionId
	err := s.updateLastOnline(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	filterMapper := &mapper.FilterMapper{}
	filterRequest := filterMapper.MapToUpdateRequest(fr)
	filterEntity, err := s.filterRepository.Update(ctx, filterRequest)
	if err != nil {
		return nil, err
	}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	return filterResponse, nil
}

func (s *ProfileService) removeStrSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func (s *ProfileService) uploadImageToFileSystem(ctx context.Context, file *entity.FileMetadata,
	sessionId string) (*request.ImageAddRequestRepositoryDto, error) {
	directoryPath := fmt.Sprintf("static/profiles/%s/images", sessionId)
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
	newFileName, newFilePath, newFileSize, err := s.convertImage(sessionId, directoryPath, filePath, filenameWithoutSpaces)
	if err != nil {
		return nil, err
	}
	imageConverted := &request.ImageAddRequestRepositoryDto{
		SessionId: sessionId,
		Name:      newFileName,
		Url:       newFilePath,
		Size:      newFileSize,
		IsDeleted: false,
		IsBlocked: false,
		IsPrimary: false,
		IsPrivate: false,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	return imageConverted, nil
}

func (s *ProfileService) convertImage(sessionId, directoryPath, filePath, fileName string) (string, string, int64, error) {
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
	pathToS3 := fmt.Sprintf("/profiles/%s/images/%s", sessionId, newFileName)
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

func (s *ProfileService) AddNavigator(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*entity.NavigatorEntity, error) {
	navigatorMapper := &mapper.NavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToAddRequest(pr)
	return s.navigatorRepository.Add(ctx, navigatorRequest)
}

func (s *ProfileService) AddFilter(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*entity.FilterEntity, error) {
	filterMapper := &mapper.FilterMapper{}
	filterRequest := filterMapper.MapToAddRequest(pr)
	return s.filterRepository.Add(ctx, filterRequest)
}

func (s *ProfileService) AddTelegram(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*entity.TelegramEntity, error) {
	telegramMapper := &mapper.TelegramMapper{}
	telegramRequest := telegramMapper.MapToAddRequest(pr)
	return s.telegramRepository.Add(ctx, telegramRequest)
}

func (s *ProfileService) AddBlock(ctx context.Context, pr *request.BlockRequestDto) (*entity.BlockEntity, error) {
	blockMapper := &mapper.BlockMapper{}
	blockRequest := blockMapper.MapToAddRequest(pr)
	prForTwoUser := &request.BlockRequestDto{
		SessionId:            pr.BlockedUserSessionId,
		BlockedUserSessionId: pr.SessionId,
	}
	blockForTwoUserRequest := blockMapper.MapToAddRequest(prForTwoUser)
	_, err := s.blockRepository.Add(ctx, blockForTwoUserRequest)
	if err != nil {
		return nil, err
	}
	return s.blockRepository.Add(ctx, blockRequest)
}

func (s *ProfileService) AddLike(
	ctx context.Context, pr *request.LikeAddRequestDto, locale string) (*response.LikeResponseDto, error) {
	sessionId := pr.SessionId
	if err := s.checkUserExists(ctx, sessionId); err != nil {
		return nil, err
	}
	telegramProfile, err := s.telegramRepository.FindBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	lastImageProfile, err := s.imageRepository.FindLastBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	likedTelegramProfile, err := s.telegramRepository.FindBySessionId(ctx, pr.LikedSessionId)
	if err != nil {
		return nil, err
	}
	go func() {
		s.hub.Broadcast <- &entity.HubContent{
			Type:         "like",
			Message:      s.GetMessageLike(locale),
			UserId:       likedTelegramProfile.UserId,
			UserImageUrl: lastImageProfile.Url,
			Username:     telegramProfile.UserName,
		}
	}()
	likeMapper := &mapper.LikeMapper{}
	likeRequest := likeMapper.MapToAddRequest(pr)
	likeAdded, err := s.likeRepository.Add(ctx, likeRequest)
	if err != nil {
		return nil, err
	}
	likeResponse := likeMapper.MapToResponse(likeAdded)
	return likeResponse, nil
}

func (s *ProfileService) UpdateLike(
	ctx context.Context, pr *request.LikeUpdateRequestDto) (*response.LikeResponseDto, error) {
	sessionId := pr.SessionId
	if err := s.checkUserExists(ctx, sessionId); err != nil {
		return nil, err
	}
	likeMapper := &mapper.LikeMapper{}
	likeRequest := likeMapper.MapToUpdateRequest(pr)
	likeUpdated, err := s.likeRepository.Update(ctx, likeRequest)
	if err != nil {
		return nil, err
	}
	likeResponse := likeMapper.MapToResponse(likeUpdated)
	return likeResponse, nil
}

func (s *ProfileService) AddComplaint(
	ctx context.Context, pr *request.ComplaintAddRequestDto) (*entity.ComplaintEntity, error) {
	complaintMapper := &mapper.ComplaintMapper{}
	complaintRequest := complaintMapper.MapToAddRequest(pr)
	return s.complaintRepository.Add(ctx, complaintRequest)
}

func (s *ProfileService) UpdateCoordinates(
	ctx context.Context, pr *request.NavigatorUpdateRequestDto) (*response.NavigatorResponseDto, error) {
	sessionId := pr.SessionId
	longitude := pr.Longitude
	latitude := pr.Latitude
	return s.updateNavigator(ctx, sessionId, longitude, latitude)
}

func (s *ProfileService) updateLastOnline(ctx context.Context, sessionId string) error {
	updateLastOnlineMapper := &mapper.ProfileUpdateLastOnlineMapper{}
	updateLastOnlineRequest := updateLastOnlineMapper.MapToAddRequest(sessionId)
	err := s.profileRepository.UpdateLastOnline(ctx, updateLastOnlineRequest)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProfileService) updateNavigator(
	ctx context.Context,
	sessionId string, longitude float64, latitude float64) (*response.NavigatorResponseDto, error) {
	fmt.Println("updateNavigator latitude", latitude)
	navigatorMapper := &mapper.NavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToUpdateRequest(sessionId, longitude, latitude)
	navigatorUpdated, err := s.navigatorRepository.Update(ctx, navigatorRequest)
	if err != nil {
		return nil, err
	}
	navigatorResponse := navigatorMapper.MapToResponse(sessionId, navigatorUpdated.Location.Longitude,
		navigatorUpdated.Location.Latitude)
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

func (s *ProfileService) checkUserExists(ctx context.Context, sessionId string) error {
	p, err := s.profileRepository.FindBySessionId(ctx, sessionId)
	if err != nil {
		return err
	}
	isDeleted := p.IsDeleted
	if isDeleted {
		err := errors.Wrap(err, "user has already been deleted")
		return err
	}
	return nil
}
