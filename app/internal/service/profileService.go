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

type ProfileRepository interface {
	AddProfile(ctx context.Context, p *request.ProfileAddRequestRepositoryDto) (*entity.ProfileEntity, error)
	UpdateProfile(ctx context.Context, p *request.ProfileUpdateRequestRepositoryDto) (*entity.ProfileEntity, error)
	DeleteProfile(ctx context.Context, p *request.ProfileDeleteRequestRepositoryDto) (*entity.ProfileEntity, error)
	FindProfileById(ctx context.Context, id uint64) (*entity.ProfileEntity, error)
	FindProfileBySessionId(ctx context.Context, sessionId string) (*entity.ProfileEntity, error)
	AddImage(ctx context.Context, p *request.ProfileImageAddRequestRepositoryDto) (*entity.ProfileImageEntity, error)
	UpdateImage(ctx context.Context,
		p *request.ProfileImageUpdateRequestRepositoryDto) (*entity.ProfileImageEntity, error)
	DeleteImage(ctx context.Context,
		p *request.ProfileImageDeleteRequestRepositoryDto) (*entity.ProfileImageEntity, error)
	FindImageById(ctx context.Context, imageId uint64) (*entity.ProfileImageEntity, error)
	SelectImageListPublicBySessionId(ctx context.Context, sessionId string) ([]*entity.ProfileImageEntity, error)
	SelectImageListBySessionId(ctx context.Context, sessionId string) ([]*entity.ProfileImageEntity, error)
	AddNavigator(ctx context.Context,
		p *request.ProfileNavigatorAddRequestRepositoryDto) (*entity.ProfileNavigatorEntity, error)
	UpdateNavigator(ctx context.Context,
		p *request.ProfileNavigatorUpdateRequestDto) (*entity.ProfileNavigatorEntity, error)
	DeleteNavigator(ctx context.Context,
		p *request.ProfileNavigatorDeleteRequestDto) (*entity.ProfileNavigatorEntity, error)
	FindNavigatorById(ctx context.Context, id uint64) (*entity.ProfileNavigatorEntity, error)
	FindNavigatorBySessionId(ctx context.Context, sessionId string) (*entity.ProfileNavigatorEntity, error)
	AddFilter(ctx context.Context, p *request.ProfileFilterAddRequestRepositoryDto) (*entity.ProfileFilterEntity, error)
	UpdateFilter(ctx context.Context,
		p *request.ProfileFilterUpdateRequestRepositoryDto) (*entity.ProfileFilterEntity, error)
	DeleteFilter(ctx context.Context,
		p *request.ProfileFilterDeleteRequestRepositoryDto) (*entity.ProfileFilterEntity, error)
	FindFilterBySessionId(ctx context.Context, sessionId string) (*entity.ProfileFilterEntity, error)
	AddTelegram(ctx context.Context,
		p *request.ProfileTelegramAddRequestRepositoryDto) (*entity.ProfileTelegramEntity, error)
	UpdateTelegram(ctx context.Context,
		p *request.ProfileTelegramUpdateRequestRepositoryDto) (*entity.ProfileTelegramEntity, error)
	DeleteTelegram(ctx context.Context,
		p *request.ProfileTelegramDeleteRequestRepositoryDto) (*entity.ProfileTelegramEntity, error)
	FindTelegramById(ctx context.Context, id uint64) (*entity.ProfileTelegramEntity, error)
	FindTelegramBySessionId(ctx context.Context, sessionId string) (*entity.ProfileTelegramEntity, error)
	AddBlock(ctx context.Context, p *request.ProfileBlockAddRequestRepositoryDto) (*entity.ProfileBlockEntity, error)
	FindBlockById(ctx context.Context, id uint64) (*entity.ProfileBlockEntity, error)
	AddLike(ctx context.Context, p *request.ProfileLikeAddRequestRepositoryDto) (*entity.ProfileLikeEntity, error)
	FindLikeById(ctx context.Context, id uint64) (*entity.ProfileLikeEntity, error)
	AddComplaint(
		ctx context.Context, p *request.ProfileComplaintAddRequestRepositoryDto) (*entity.ProfileComplaintEntity, error)
	FindComplaintById(ctx context.Context, id uint64) (*entity.ProfileComplaintEntity, error)
	UpdateLastOnline(ctx context.Context, p *request.ProfileUpdateLastOnlineRequestRepositoryDto) error
}

type ProfileService struct {
	logger     logger.Logger
	repository ProfileRepository
}

func NewProfileService(l logger.Logger, r ProfileRepository) *ProfileService {
	return &ProfileService{
		logger:     l,
		repository: r,
	}
}

func (s *ProfileService) AddProfile(
	ctx context.Context, ctf *fiber.Ctx, pr *request.ProfileAddRequestDto) (*response.ProfileAddResponseDto, error) {
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToAddRequest(pr)
	profileCreated, err := s.repository.AddProfile(ctx, profileRequest)
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
	profileEntity, err := s.repository.UpdateProfile(ctx, profileRequest)
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
	filterMapper := &mapper.ProfileFilterMapper{}
	filterRequest := filterMapper.MapToUpdateRequest(pr)
	filterEntity, err := s.repository.UpdateFilter(ctx, filterRequest)
	if err != nil {
		return nil, err
	}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	telegramMapper := &mapper.ProfileTelegramMapper{}
	telegramRequest := telegramMapper.MapToUpdateRequest(pr)
	telegramEntity, err := s.repository.UpdateTelegram(ctx, telegramRequest)
	if err != nil {
		return nil, err
	}
	telegramResponse := telegramMapper.MapToResponse(telegramEntity)
	imageEntityList, err := s.repository.SelectImageListPublicBySessionId(ctx, sessionId)
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
	_, err := s.repository.DeleteProfile(ctx, profileRequest)
	if err != nil {
		return nil, err
	}
	navigatorMapper := &mapper.ProfileNavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToDeleteRequest(sessionId)
	_, err = s.repository.DeleteNavigator(ctx, navigatorRequest)
	if err != nil {
		return nil, err
	}
	err = s.DeleteImageList(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	filterMapper := &mapper.ProfileFilterMapper{}
	filterRequest := filterMapper.MapToDeleteRequest(sessionId)
	_, err = s.repository.DeleteFilter(ctx, filterRequest)
	if err != nil {
		return nil, err
	}
	telegramMapper := &mapper.ProfileTelegramMapper{}
	telegramRequest := telegramMapper.MapToDeleteRequest(sessionId)
	_, err = s.repository.DeleteTelegram(ctx, telegramRequest)
	if err != nil {
		return nil, err
	}
	profileResponse := &response.ResponseDto{
		Success: true,
	}
	return profileResponse, err
}

func (s *ProfileService) GetProfileBySessionId(
	ctx context.Context, sessionId string, pr *request.ProfileGetBySessionIdRequestDto) (*response.ProfileResponseDto, error) {
	if err := s.checkUserExists(ctx, sessionId); err != nil {
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileEntity, err := s.repository.FindProfileBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	err = s.updateLastOnline(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	navigatorResponse, err := s.updateNavigator(ctx, sessionId, pr.Longitude, pr.Latitude)
	if err != nil {
		return nil, err
	}
	filterMapper := &mapper.ProfileFilterMapper{}
	filterEntity, err := s.repository.FindFilterBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	filterResponse := filterMapper.MapToResponse(filterEntity)
	telegramEntity, err := s.repository.FindTelegramBySessionId(ctx, sessionId)
	if err != nil {
		return nil, err
	}
	telegramMapper := &mapper.ProfileTelegramMapper{}
	telegramResponse := telegramMapper.MapToResponse(telegramEntity)
	isOnline := s.checkIsOnline(profileEntity.LastOnline)
	imageEntityList, err := s.repository.SelectImageListPublicBySessionId(ctx, sessionId)
	profileResponse := profileMapper.MapToResponse(
		profileEntity, navigatorResponse, filterResponse, telegramResponse, imageEntityList, isOnline)
	return profileResponse, err
}

func (s *ProfileService) AddImageList(
	ctx context.Context, ctf *fiber.Ctx, sessionId string) error {
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
	file *multipart.FileHeader) (*entity.ProfileImageEntity, error) {
	imageConverted, err := s.uploadImageToFileSystem(ctx, ctf, file, sessionId)
	if err != nil {
		errorMessage := s.getErrorMessage("AddImage", "uploadImageToFileSystem")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return s.repository.AddImage(ctx, imageConverted)
}

func (s *ProfileService) UpdateImageList(
	ctx context.Context, ctf *fiber.Ctx, sessionId string) error {
	return s.AddImageList(ctx, ctf, sessionId)
}

func (s *ProfileService) DeleteImageList(ctx context.Context, sessionId string) error {
	imageList, err := s.repository.SelectImageListBySessionId(ctx, sessionId)
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

func (s *ProfileService) DeleteImage(
	ctx context.Context, image *entity.ProfileImageEntity) (*entity.ProfileImageEntity, error) {
	filePath := image.Url
	if err := os.Remove(filePath); err != nil {
		errorMessage := s.getErrorMessage("DeleteImage", "Remove")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageMapper := &mapper.ProfileImageMapper{}
	imageRequest := imageMapper.MapToDeleteRequest(image.Id)
	return s.repository.DeleteImage(ctx, imageRequest)
}

func (s *ProfileService) uploadImageToFileSystem(ctx context.Context, ctf *fiber.Ctx, file *multipart.FileHeader,
	sessionId string) (*request.ProfileImageAddRequestRepositoryDto, error) {
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
	imageConverted := &request.ProfileImageAddRequestRepositoryDto{
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
	ctx context.Context, pr *request.ProfileAddRequestDto) (*entity.ProfileNavigatorEntity, error) {
	navigatorMapper := &mapper.ProfileNavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToAddRequest(pr)
	return s.repository.AddNavigator(ctx, navigatorRequest)
}

func (s *ProfileService) AddFilter(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*entity.ProfileFilterEntity, error) {
	filterMapper := &mapper.ProfileFilterMapper{}
	filterRequest := filterMapper.MapToAddRequest(pr)
	return s.repository.AddFilter(ctx, filterRequest)
}

func (s *ProfileService) AddTelegram(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*entity.ProfileTelegramEntity, error) {
	telegramMapper := &mapper.ProfileTelegramMapper{}
	telegramRequest := telegramMapper.MapToAddRequest(pr)
	return s.repository.AddTelegram(ctx, telegramRequest)
}

func (s *ProfileService) AddBlock(
	ctx context.Context, pr *request.ProfileBlockRequestDto) (*entity.ProfileBlockEntity, error) {
	blockMapper := &mapper.ProfileBlockMapper{}
	blockRequest := blockMapper.MapToAddRequest(pr)
	return s.repository.AddBlock(ctx, blockRequest)
}

func (s *ProfileService) AddLike(
	ctx context.Context, pr *request.ProfileLikeAddRequestDto) (*response.ProfileLikeResponseDto, error) {
	likeMapper := &mapper.ProfileLikeMapper{}
	likeRequest := likeMapper.MapToAddRequest(pr)
	likeAdded, err := s.repository.AddLike(ctx, likeRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddLike", "AddLike")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	likeResponse := likeMapper.MapToAddResponse(likeAdded)
	return likeResponse, nil
}
func (s *ProfileService) AddComplaint(
	ctx context.Context, pr *request.ProfileComplaintAddRequestDto) (*entity.ProfileComplaintEntity, error) {
	complaintMapper := &mapper.ProfileComplaintMapper{}
	complaintRequest := complaintMapper.MapToAddRequest(pr)
	return s.repository.AddComplaint(ctx, complaintRequest)
}

func (s *ProfileService) updateLastOnline(ctx context.Context, sessionId string) error {
	updateLastOnlineMapper := &mapper.ProfileUpdateLastOnlineMapper{}
	updateLastOnlineRequest := updateLastOnlineMapper.MapToAddRequest(sessionId)
	err := s.repository.UpdateLastOnline(ctx, updateLastOnlineRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("updateLastOnline", "UpdateLastOnline")
		s.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	return nil
}

func (s *ProfileService) updateNavigator(
	ctx context.Context,
	sessionId string, longitude float64, latitude float64) (*response.ProfileNavigatorResponseDto, error) {
	navigatorMapper := &mapper.ProfileNavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToUpdateRequest(sessionId, longitude, latitude)
	navigatorUpdated, err := s.repository.UpdateNavigator(ctx, navigatorRequest)
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
	p, err := s.repository.FindProfileBySessionId(ctx, sessionId)
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
