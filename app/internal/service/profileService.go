package service

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/service/mapper"
	"github.com/gofiber/fiber/v2"
	"github.com/h2non/bimg"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"image/jpeg"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	errorFilePath = "internal/service/profileService.go"
)

type ProfileService struct {
	logger              logger.Logger
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
	pr ProfileRepository,
	nr NavigatorRepository,
	fr FilterRepository,
	tr TelegramRepository,
	ir ImageRepository, lr LikeRepository, br BlockRepository, cr ComplaintRepository) *ProfileService {
	return &ProfileService{
		logger:              l,
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
	ctx context.Context, ctf *fiber.Ctx, pr *request.ProfileAddRequestDto) (*response.ProfileAddResponseDto, error) {
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToAddRequest(pr)
	profileCreated, err := s.profileRepository.AddProfile(ctx, profileRequest)
	if err != nil {
		return nil, err
	}
	_, err = s.AddNavigator(ctx, pr)
	if err != nil {
		return nil, err
	}
	responseProfile := profileMapper.MapToAddResponse(profileCreated)
	if err := s.AddImageList(ctx, ctf, profileCreated.SessionId); err != nil {
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
	return responseProfile, err
}

func (s *ProfileService) UpdateProfile(ctx context.Context, ctf *fiber.Ctx,
	pr *request.ProfileUpdateRequestDto) (*response.ProfileResponseDto, error) {
	sessionId := pr.SessionId
	if err := s.checkUserExists(ctx, sessionId); err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToUpdateRequest(pr)
	profileEntity, err := s.profileRepository.UpdateProfile(ctx, profileRequest)
	if err != nil {
		return nil, err
	}
	if err := s.UpdateImageList(ctx, ctf, profileEntity.SessionId); err != nil {
		return nil, err
	}
	navigatorResponse, err := s.updateNavigator(ctx, sessionId, pr.Longitude, pr.Latitude)
	if err != nil {
		return nil, err
	}
	filterMapper := &mapper.FilterMapper{}
	filterRequest := filterMapper.MapToUpdateRequest(pr)
	filterEntity, err := s.filterRepository.UpdateFilter(ctx, filterRequest)
	if err != nil {
		return nil, err
	}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	telegramMapper := &mapper.TelegramMapper{}
	telegramRequest := telegramMapper.MapToUpdateRequest(pr)
	telegramEntity, err := s.telegramRepository.UpdateTelegram(ctx, telegramRequest)
	if err != nil {
		return nil, err
	}
	telegramResponse := telegramMapper.MapToResponse(telegramEntity)
	imageEntityList, err := s.imageRepository.SelectImageListPublicBySessionId(ctx, sessionId)
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
	_, err := s.profileRepository.DeleteProfile(ctx, profileRequest)
	if err != nil {
		return nil, err
	}
	navigatorMapper := &mapper.NavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToDeleteRequest(sessionId)
	_, err = s.navigatorRepository.DeleteNavigator(ctx, navigatorRequest)
	if err != nil {
		return nil, err
	}
	err = s.DeleteImageList(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	filterMapper := &mapper.FilterMapper{}
	filterRequest := filterMapper.MapToDeleteRequest(sessionId)
	_, err = s.filterRepository.DeleteFilter(ctx, filterRequest)
	if err != nil {
		return nil, err
	}
	telegramMapper := &mapper.TelegramMapper{}
	telegramRequest := telegramMapper.MapToDeleteRequest(sessionId)
	_, err = s.telegramRepository.DeleteTelegram(ctx, telegramRequest)
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
	profileMapper := &mapper.ProfileMapper{}
	profileEntity, err := s.profileRepository.FindProfileBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	err = s.updateLastOnline(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	if pr.Longitude != 0 && pr.Latitude != 0 {
		_, err = s.updateNavigator(ctx, sessionId, pr.Longitude, pr.Latitude)
		if err != nil {
			return nil, err
		}
	}
	navigatorMapper := &mapper.NavigatorMapper{}
	navigatorEntity, err := s.navigatorRepository.FindNavigatorBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	longitude := navigatorEntity.Location.Longitude
	latitude := navigatorEntity.Location.Latitude
	navigatorResponse := navigatorMapper.MapToResponse(sessionId, longitude, latitude)
	filterMapper := &mapper.FilterMapper{}
	filterEntity, err := s.filterRepository.FindFilterBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	telegramEntity, err := s.telegramRepository.FindTelegramBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	telegramMapper := &mapper.TelegramMapper{}
	telegramResponse := telegramMapper.MapToResponse(telegramEntity)
	isOnline := s.checkIsOnline(profileEntity.LastOnline)
	imageEntityList, err := s.imageRepository.SelectImageListPublicBySessionId(ctx, sessionId)
	profileResponse := profileMapper.MapToResponse(
		profileEntity, navigatorResponse, filterResponse, telegramResponse, imageEntityList, isOnline)
	return profileResponse, err
}

func (s *ProfileService) AddImageList(ctx context.Context, ctf *fiber.Ctx, sessionId string) error {
	form, err := ctf.MultipartForm()
	if err != nil {
		errorMessage := s.getErrorMessage("AddImageList", "MultipartForm")
		s.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	files := form.File["image"]
	if len(files) > 0 {
		for _, file := range files {
			_, err := s.AddImage(ctx, ctf, sessionId, file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *ProfileService) AddImage(ctx context.Context, ctf *fiber.Ctx, sessionId string,
	file *multipart.FileHeader) (*entity.ImageEntity, error) {
	imageConverted, err := s.uploadImageToFileSystem(ctx, ctf, file, sessionId)
	if err != nil {
		errorMessage := s.getErrorMessage("AddImage", "uploadImageToFileSystem")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return s.imageRepository.AddImage(ctx, imageConverted)
}

func (s *ProfileService) UpdateImageList(ctx context.Context, ctf *fiber.Ctx, sessionId string) error {
	return s.AddImageList(ctx, ctf, sessionId)
}

func (s *ProfileService) DeleteImageList(ctx context.Context, sessionId string) error {
	imageList, err := s.imageRepository.SelectImageListBySessionId(ctx, sessionId)
	if err != nil {
		return err
	}
	if len(imageList) > 0 {
		for _, image := range imageList {
			_, err := s.DeleteImage(ctx, image)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *ProfileService) DeleteImage(ctx context.Context, image *entity.ImageEntity) (*entity.ImageEntity, error) {
	filePath := image.Url
	if err := os.Remove(filePath); err != nil {
		errorMessage := s.getErrorMessage("DeleteImage", "Remove")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageMapper := &mapper.ImageMapper{}
	imageRequest := imageMapper.MapToDeleteRequest(image.Id)
	return s.imageRepository.DeleteImage(ctx, imageRequest)
}

func (s *ProfileService) uploadImageToFileSystem(ctx context.Context, ctf *fiber.Ctx, file *multipart.FileHeader,
	sessionId string) (*request.ImageAddRequestRepositoryDto, error) {
	directoryPath := fmt.Sprintf("static/profiles/%s/images", sessionId)
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		if err := os.MkdirAll(directoryPath, 0755); err != nil {
			errorMessage := s.getErrorMessage("uploadImageToFileSystem", "MkdirAll")
			s.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
	}
	filePath := fmt.Sprintf("%s/%s", directoryPath, file.Filename)
	if err := ctf.SaveFile(file, filePath); err != nil {
		errorMessage := s.getErrorMessage("uploadImageToFileSystem", "SaveFile")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	newFileName, newFilePath, newFileSize, err := s.convertImage(directoryPath, filePath, file.Filename)
	if err != nil {
		errorMessage := s.getErrorMessage("uploadImageToFileSystem", "convertImage")
		s.logger.Debug(errorMessage, zap.Error(err))
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

func (s *ProfileService) convertImage(directoryPath, filePath, fileName string) (string, string, int64, error) {
	fileImage, err := os.Open(filePath)
	if err != nil {
		errorMessage := s.getErrorMessage("convertImage", "Open")
		s.logger.Debug(errorMessage, zap.Error(err))
		return "", "", 0, err
	}
	// The Decode function is used to read images from a file or other source and convert them into an image.
	// Image structure
	_, err = jpeg.Decode(fileImage)
	if err != nil {
		errorMessage := s.getErrorMessage("convertImage", "Decode")
		s.logger.Debug(errorMessage, zap.Error(err))
		return "", "", 0, err
	}
	newFileName := s.replaceExtension(fileName)
	newFilePath := fmt.Sprintf("%s/%s", directoryPath, newFileName)
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
	bimg.Write(newFilePath, newFile)
	if err := s.deleteFile(filePath); err != nil {
		errorMessage := s.getErrorMessage("convertImage", "deleteFile")
		s.logger.Debug(errorMessage, zap.Error(err))
		return "", "", 0, err
	}
	newFileInfo, err := os.Stat(newFilePath)
	if err != nil {
		errorMessage := s.getErrorMessage("convertImage", "Stat")
		s.logger.Debug(errorMessage, zap.Error(err))
		return "", "", 0, err
	}
	newFileSize := newFileInfo.Size()
	return newFileName, newFilePath, newFileSize, nil
}

func (s *ProfileService) deleteFile(filePath string) error {
	return os.Remove(filePath)
}

func (s *ProfileService) AddNavigator(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*entity.NavigatorEntity, error) {
	navigatorMapper := &mapper.NavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToAddRequest(pr)
	return s.navigatorRepository.AddNavigator(ctx, navigatorRequest)
}

func (s *ProfileService) AddFilter(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*entity.FilterEntity, error) {
	filterMapper := &mapper.FilterMapper{}
	filterRequest := filterMapper.MapToAddRequest(pr)
	return s.filterRepository.AddFilter(ctx, filterRequest)
}

func (s *ProfileService) AddTelegram(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*entity.TelegramEntity, error) {
	telegramMapper := &mapper.TelegramMapper{}
	telegramRequest := telegramMapper.MapToAddRequest(pr)
	return s.telegramRepository.AddTelegram(ctx, telegramRequest)
}

func (s *ProfileService) AddBlock(ctx context.Context, pr *request.BlockRequestDto) (*entity.BlockEntity, error) {
	blockMapper := &mapper.BlockMapper{}
	blockRequest := blockMapper.MapToAddRequest(pr)
	return s.blockRepository.AddBlock(ctx, blockRequest)
}

func (s *ProfileService) AddLike(
	ctx context.Context, pr *request.LikeAddRequestDto) (*response.LikeResponseDto, error) {
	likeMapper := &mapper.LikeMapper{}
	likeRequest := likeMapper.MapToAddRequest(pr)
	likeAdded, err := s.likeRepository.AddLike(ctx, likeRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddLike", "AddLike")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	likeResponse := likeMapper.MapToAddResponse(likeAdded)
	return likeResponse, nil
}
func (s *ProfileService) AddComplaint(
	ctx context.Context, pr *request.ComplaintAddRequestDto) (*entity.ComplaintEntity, error) {
	complaintMapper := &mapper.ComplaintMapper{}
	complaintRequest := complaintMapper.MapToAddRequest(pr)
	return s.complaintRepository.AddComplaint(ctx, complaintRequest)
}

func (s *ProfileService) updateLastOnline(ctx context.Context, sessionId string) error {
	updateLastOnlineMapper := &mapper.ProfileUpdateLastOnlineMapper{}
	updateLastOnlineRequest := updateLastOnlineMapper.MapToAddRequest(sessionId)
	err := s.profileRepository.UpdateLastOnline(ctx, updateLastOnlineRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("updateLastOnline", "UpdateLastOnline")
		s.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	return nil
}

func (s *ProfileService) updateNavigator(
	ctx context.Context,
	sessionId string, longitude float64, latitude float64) (*response.NavigatorResponseDto, error) {
	navigatorMapper := &mapper.NavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToUpdateRequest(sessionId, longitude, latitude)
	navigatorUpdated, err := s.navigatorRepository.UpdateNavigator(ctx, navigatorRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("updateNavigator", "UpdateNavigator")
		s.logger.Debug(errorMessage, zap.Error(err))
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

func (s *ProfileService) replaceExtension(filename string) string {
	// Удаляем текущее расширение
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	// Добавляем новое расширение .webp
	return filename + ".webp"
}

func (s *ProfileService) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}

func (s *ProfileService) checkUserExists(ctx context.Context, sessionId string) error {
	p, err := s.profileRepository.FindProfileBySessionId(ctx, sessionId)
	if err != nil {
		errorMessage := s.getErrorMessage("checkUserExists", "FindProfileBySessionId")
		s.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	isDeleted := p.IsDeleted
	if isDeleted {
		err := errors.Wrap(err, "user has already been deleted")
		return err
	}
	return nil
}
