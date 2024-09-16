package service

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/h2non/bimg"
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
	AddProfile(ctx context.Context, p *entity.ProfileEntity) (*entity.ProfileEntity, error)
	AddImage(ctx context.Context, p *entity.ProfileImageEntity) (*entity.ProfileImageEntity, error)
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
	ctx context.Context, pr *request.ProfileAddRequestDto) (*response.ProfileAddResponseDto, error) {
	profile := &entity.ProfileEntity{
		SessionID:      pr.SessionID,
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
	newProfile, err := s.repository.AddProfile(ctx, profile)
	if err != nil {
		errorMessage := s.getErrorMessage("AddProfile", "AddProfile")
		s.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	responseProfile := &response.ProfileAddResponseDto{
		SessionID: newProfile.SessionID,
	}
	return responseProfile, err
}

func (s *ProfileService) AddImage(ctx context.Context, ctf *fiber.Ctx, sessionId string) error {
	form, err := ctf.MultipartForm()
	if err != nil {
		errorMessage := s.getErrorMessage("AddProfile", "MultipartForm")
		s.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	files := form.File["image"]
	for _, file := range files {
		imageConverted, err := s.uploadImageToFileSystem(ctx, ctf, file, sessionId)
		if err != nil {
			errorMessage := s.getErrorMessage("AddImage", "uploadImageToFileSystem")
			s.logger.Debug(errorMessage, zap.Error(err))
			return err
		}
		_, err = s.repository.AddImage(ctx, imageConverted)
		if err != nil {
			errorMessage := s.getErrorMessage("AddImage", "AddImage")
			s.logger.Debug(errorMessage, zap.Error(err))
			return err
		}
	}
	return nil
}

func (s *ProfileService) uploadImageToFileSystem(ctx context.Context, ctf *fiber.Ctx, file *multipart.FileHeader,
	sessionId string) (*entity.ProfileImageEntity, error) {
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
	imageConverted := &entity.ProfileImageEntity{
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
