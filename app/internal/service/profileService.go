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
	FindProfileByID(ctx context.Context, id uint64) (*entity.ProfileEntity, error)
	FindProfileBySessionID(ctx context.Context, sessionID string) (*entity.ProfileEntity, error)
	AddImage(ctx context.Context, p *request.ProfileImageAddRequestRepositoryDto) (*entity.ProfileImageEntity, error)
	UpdateImage(ctx context.Context,
		p *request.ProfileImageUpdateRequestRepositoryDto) (*entity.ProfileImageEntity, error)
	DeleteImage(ctx context.Context,
		p *request.ProfileImageDeleteRequestRepositoryDto) (*entity.ProfileImageEntity, error)
	FindImageById(ctx context.Context, imageID uint64) (*entity.ProfileImageEntity, error)
	SelectImageListPublicBySessionID(ctx context.Context, sessionID string) ([]*entity.ProfileImageEntity, error)
	SelectImageListBySessionID(ctx context.Context, sessionID string) ([]*entity.ProfileImageEntity, error)
	AddNavigator(ctx context.Context,
		p *request.ProfileNavigatorAddRequestRepositoryDto) (*entity.ProfileNavigatorEntity, error)
	UpdateNavigator(ctx context.Context,
		p *request.ProfileNavigatorUpdateRequestDto) (*entity.ProfileNavigatorEntity, error)
	DeleteNavigator(ctx context.Context,
		p *request.ProfileNavigatorDeleteRequestDto) (*entity.ProfileNavigatorEntity, error)
	FindNavigatorByID(ctx context.Context, id uint64) (*entity.ProfileNavigatorEntity, error)
	FindNavigatorBySessionID(ctx context.Context, sessionID string) (*entity.ProfileNavigatorEntity, error)
	AddFilter(ctx context.Context, p *request.ProfileFilterAddRequestRepositoryDto) (*entity.ProfileFilterEntity, error)
	UpdateFilter(ctx context.Context,
		p *request.ProfileFilterUpdateRequestRepositoryDto) (*entity.ProfileFilterEntity, error)
	DeleteFilter(ctx context.Context,
		p *request.ProfileFilterDeleteRequestRepositoryDto) (*entity.ProfileFilterEntity, error)
	FindFilterBySessionID(ctx context.Context, sessionID string) (*entity.ProfileFilterEntity, error)
	AddTelegram(ctx context.Context,
		p *request.ProfileTelegramAddRequestRepositoryDto) (*entity.ProfileTelegramEntity, error)
	UpdateTelegram(ctx context.Context,
		p *request.ProfileTelegramUpdateRequestRepositoryDto) (*entity.ProfileTelegramEntity, error)
	DeleteTelegram(ctx context.Context,
		p *request.ProfileTelegramDeleteRequestRepositoryDto) (*entity.ProfileTelegramEntity, error)
	FindTelegramByID(ctx context.Context, id uint64) (*entity.ProfileTelegramEntity, error)
	FindTelegramBySessionID(ctx context.Context, sessionID string) (*entity.ProfileTelegramEntity, error)
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
		errorMessage := s.getErrorMessage("AddProfile", "AddProfile")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	_, err = s.AddNavigator(ctx, pr)
	if err != nil {
		errorMessage := s.getErrorMessage("AddProfile", "AddNavigator")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	responseProfile := profileMapper.MapToAddResponse(profileCreated)
	if err := s.AddImageList(ctx, ctf, profileCreated.SessionID); err != nil {
		errorMessage := s.getErrorMessage("AddProfile", "AddImageList")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	_, err = s.AddFilter(ctx, pr)
	if err != nil {
		errorMessage := s.getErrorMessage("AddProfile", "AddFilter")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	_, err = s.AddTelegram(ctx, pr)
	if err != nil {
		errorMessage := s.getErrorMessage("AddProfile", "AddTelegram")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return responseProfile, err
}

func (s *ProfileService) UpdateProfile(ctx context.Context, ctf *fiber.Ctx,
	pr *request.ProfileUpdateRequestDto) (*response.ProfileUpdateResponseDto, error) {
	isDeleted, err := s.checkUserExists(ctx, pr.SessionID)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile", "checkUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if isDeleted {
		err := errors.Wrap(err, "user has already been deleted")
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToUpdateRequest(pr)
	profileUpdated, err := s.repository.UpdateProfile(ctx, profileRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile", "UpdateProfile")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if err := s.UpdateImageList(ctx, ctf, profileUpdated.SessionID); err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile", "UpdateImageList")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	navigatorMapper := &mapper.ProfileNavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToUpdateRequest(profileUpdated, pr)
	navigatorUpdated, err := s.repository.UpdateNavigator(ctx, navigatorRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile", "UpdateNavigator")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	navigatorResponse := navigatorMapper.MapToResponse(profileUpdated, navigatorUpdated)
	filterMapper := &mapper.ProfileFilterMapper{}
	filterRequest := filterMapper.MapToUpdateRequest(pr)
	filterUpdated, err := s.repository.UpdateFilter(ctx, filterRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile", "UpdateFilter")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	filterResponse := filterMapper.MapToResponse(filterUpdated)
	telegramMapper := &mapper.ProfileTelegramMapper{}
	telegramRequest := telegramMapper.MapToUpdateRequest(pr)
	telegramUpdated, err := s.repository.UpdateTelegram(ctx, telegramRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("UpdateProfile", "UpdateTelegram")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	telegramResponse := telegramMapper.MapToResponse(telegramUpdated)
	isOnline := s.checkIsOnline(profileUpdated.LastOnline)
	profileResponse := profileMapper.MapToResponse(profileUpdated, navigatorResponse, filterResponse, telegramResponse,
		isOnline)
	return profileResponse, nil
}

func (s *ProfileService) DeleteProfile(
	ctx context.Context, pr *request.ProfileDeleteRequestDto) (*response.ResponseDto, error) {
	sessionID := pr.SessionID
	isDeleted, err := s.checkUserExists(ctx, sessionID)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile", "checkUserExists")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if isDeleted {
		err := errors.Wrap(err, "user has already been deleted")
		return nil, err
	}
	profileMapper := &mapper.ProfileMapper{}
	profileRequest := profileMapper.MapToDeleteRequest(sessionID)
	_, err = s.repository.DeleteProfile(ctx, profileRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile", "DeleteProfile")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	navigatorMapper := &mapper.ProfileNavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToDeleteRequest(sessionID)
	_, err = s.repository.DeleteNavigator(ctx, navigatorRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile", "DeleteNavigator")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	err = s.DeleteImageList(ctx, sessionID)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile", "DeleteImageList")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	filterMapper := &mapper.ProfileFilterMapper{}
	filterRequest := filterMapper.MapToDeleteRequest(sessionID)
	_, err = s.repository.DeleteFilter(ctx, filterRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile", "DeleteFilter")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	telegramMapper := &mapper.ProfileTelegramMapper{}
	telegramRequest := telegramMapper.MapToDeleteRequest(sessionID)
	_, err = s.repository.DeleteTelegram(ctx, telegramRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteProfile", "DeleteTelegram")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileResponse := &response.ResponseDto{
		Success: true,
	}
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
				errorMessage := s.getErrorMessage("AddImageList", "AddImage")
				s.logger.Debug(errorMessage, zap.Error(err))
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
	imageResponse, err := s.repository.AddImage(ctx, imageConverted)
	if err != nil {
		errorMessage := s.getErrorMessage("AddImage", "AddImage")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return imageResponse, err
}

func (s *ProfileService) UpdateImageList(
	ctx context.Context, ctf *fiber.Ctx, sessionId string) error {
	return s.AddImageList(ctx, ctf, sessionId)
}

func (s *ProfileService) DeleteImageList(ctx context.Context, sessionId string) error {
	imageList, err := s.repository.SelectImageListBySessionID(ctx, sessionId)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteImageList",
			"SelectImageListBySessionID")
		s.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	if len(imageList) > 0 {
		for _, image := range imageList {
			_, err := s.DeleteImage(ctx, image)
			if err != nil {
				errorMessage := s.getErrorMessage("DeleteImageList", "DeleteImage")
				s.logger.Debug(errorMessage, zap.Error(err))
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
	imageRequest := imageMapper.MapToDeleteRequest(image.ID)
	imageDeleted, err := s.repository.DeleteImage(ctx, imageRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("DeleteImage", "DeleteImage")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return imageDeleted, nil
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
		SessionID: sessionId,
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
	if err := os.Remove(filePath); err != nil {
		errorMessage := s.getErrorMessage("deleteFile", "Remove")
		s.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	return nil
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

func (s *ProfileService) AddNavigator(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*entity.ProfileNavigatorEntity, error) {
	navigatorMapper := &mapper.ProfileNavigatorMapper{}
	navigatorRequest := navigatorMapper.MapToAddRequest(pr)
	NavigatorResponse, err := s.repository.AddNavigator(ctx, navigatorRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddNavigator", "AddNavigator")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return NavigatorResponse, nil
}

func (s *ProfileService) AddFilter(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*entity.ProfileFilterEntity, error) {
	filterMapper := &mapper.ProfileFilterMapper{}
	filterRequest := filterMapper.MapToAddRequest(pr)
	filterResponse, err := s.repository.AddFilter(ctx, filterRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddFilter", "AddFilter")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return filterResponse, nil
}

func (s *ProfileService) AddTelegram(
	ctx context.Context, pr *request.ProfileAddRequestDto) (*entity.ProfileTelegramEntity, error) {
	telegramMapper := &mapper.ProfileTelegramMapper{}
	telegramRequest := telegramMapper.MapToAddRequest(pr)
	telegramResponse, err := s.repository.AddTelegram(ctx, telegramRequest)
	if err != nil {
		errorMessage := s.getErrorMessage("AddTelegram", "AddTelegram")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return telegramResponse, nil
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

func (s *ProfileService) checkUserExists(ctx context.Context, sessionID string) (bool, error) {
	p, err := s.repository.FindProfileBySessionID(ctx, sessionID)
	if err != nil {
		errorMessage := s.getErrorMessage("checkUserExists", "FindProfileBySessionID")
		s.logger.Debug(errorMessage, zap.Error(err))
		return false, err
	}
	isDeleted := p.IsDeleted
	if isDeleted {
		return true, err
	}
	return false, nil
}
