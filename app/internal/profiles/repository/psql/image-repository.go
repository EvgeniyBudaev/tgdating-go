package psql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"go.uber.org/zap"
)

const (
	errorFilePathImage = "internal/repository/psql/image-repository.go"
)

type ImageRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewImageRepository(l logger.Logger, db *sql.DB) *ImageRepository {
	return &ImageRepository{
		logger: l,
		db:     db,
	}
}

func (r *ImageRepository) Add(
	ctx context.Context, p *request.ImageAddRequestRepositoryDto) (uint64, error) {
	query := "INSERT INTO dating.profile_images (telegram_user_id, name, url, size, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.Name, &p.Url, &p.Size, &p.CreatedAt, &p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return 0, err
	}
	return id, nil
}

func (r *ImageRepository) Delete(
	ctx context.Context, id uint64) (*response.ResponseDto, error) {
	query := "DELETE FROM dating.profile_images WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		errorMessage := r.getErrorMessage("Delete", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageResponse := &response.ResponseDto{
		Success: true,
	}
	return imageResponse, nil
}

func (r *ImageRepository) FindById(ctx context.Context, imageId uint64) (*response.ImageResponseRepositoryDto, error) {
	p := &response.ImageResponseRepositoryDto{}
	query := "SELECT pi.id, pi.telegram_user_id, pi.name, pi.url, pi.size, pis.is_blocked, pis.is_primary," +
		" pis.is_private, pi.created_at, pi.updated_at" +
		" FROM dating.profile_images pi" +
		" JOIN dating.profile_image_statuses pis ON pi.id = pis.id" +
		" WHERE pi.id = $1"
	row := r.db.QueryRowContext(ctx, query, imageId)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.Name, &p.Url, &p.Size, &p.IsBlocked, &p.IsPrimary, &p.IsPrivate,
		&p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ImageRepository) FindLastByTelegramUserId(
	ctx context.Context, telegramUserId string) (*response.ImageResponseDto, error) {
	p := &response.ImageResponseRepositoryDto{}
	s := &response.ImageStatusResponseDto{}
	query := "SELECT pi.id, pi.telegram_user_id, pi.name, pi.url, pi.size, pis.is_blocked, pis.is_primary," +
		" pis.is_private, pi.created_at, pi.updated_at" +
		" FROM dating.profile_images pi" +
		" JOIN dating.profile_image_statuses pis ON pi.id = pis.id" +
		" WHERE pi.telegram_user_id = $1 AND pis.is_blocked = false AND pis.is_private = false" +
		" ORDER BY pi.id DESC" +
		" LIMIT 1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.Name, &p.Url, &p.Size, &s.IsBlocked, &s.IsPrimary, &s.IsPrivate,
		&p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindLastByTelegramUserId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	result := &response.ImageResponseDto{
		Id:   p.Id,
		Name: p.Name,
		Url:  p.Url,
	}
	return result, nil
}

func (r *ImageRepository) SelectListAllByTelegramUserId(
	ctx context.Context, telegramUserId string) ([]*response.ImageResponseDto, error) {
	query := "SELECT pi.id, pi.telegram_user_id, pi.name, pi.url, pi.size, pis.is_blocked, pis.is_primary," +
		" pis.is_private, pi.created_at, pi.updated_at" +
		" FROM dating.profile_images pi" +
		" JOIN dating.profile_image_statuses pis ON pi.id = pis.id" +
		" WHERE pi.telegram_user_id = $1"
	rows, err := r.db.QueryContext(ctx, query, telegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectListAllByTelegramUserId",
			"QueryContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*response.ImageResponseDto, 0)
	for rows.Next() {
		p := &response.ImageResponseDto{}
		pr := &response.ImageResponseRepositoryDto{}
		err := rows.Scan(&p.Id, &pr.TelegramUserId, &p.Name, &p.Url, &pr.Size, &pr.IsBlocked, &pr.IsPrimary,
			&pr.IsPrivate, &pr.CreatedAt, &pr.UpdatedAt)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectListAllByTelegramUserId",
				"Scan")
			r.logger.Debug(errorMessage, zap.Error(err))
			continue
		}
		result := &response.ImageResponseDto{
			Id:   p.Id,
			Name: p.Name,
			Url:  p.Url,
		}
		list = append(list, result)
	}
	return list, nil
}

func (r *ImageRepository) SelectListPublicByTelegramUserId(
	ctx context.Context, telegramUserId string) ([]*response.ImageResponseDto, error) {
	query := "SELECT pi.id, pi.telegram_user_id, pi.name, pi.url, pi.size, pis.is_blocked, pis.is_primary," +
		" pis.is_private, pi.created_at, pi.updated_at" +
		" FROM dating.profile_images pi" +
		" JOIN dating.profile_image_statuses pis ON pi.id = pis.id" +
		" WHERE pi.telegram_user_id = $1 AND pis.is_blocked = false AND pis.is_private = false"
	rows, err := r.db.QueryContext(ctx, query, telegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectListPublicByTelegramUserId",
			"QueryContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*response.ImageResponseDto, 0)
	for rows.Next() {
		p := &response.ImageResponseDto{}
		pr := &response.ImageResponseRepositoryDto{}
		err := rows.Scan(&p.Id, &pr.TelegramUserId, &p.Name, &p.Url, &pr.Size, &pr.IsBlocked, &pr.IsPrimary,
			&pr.IsPrivate, &pr.CreatedAt, &pr.UpdatedAt)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectListPublicByTelegramUserId",
				"Scan")
			r.logger.Debug(errorMessage, zap.Error(err))
			continue
		}
		result := &response.ImageResponseDto{
			Id:   p.Id,
			Name: p.Name,
			Url:  p.Url,
		}
		list = append(list, result)
	}
	return list, nil
}

func (r *ImageRepository) SelectListByTelegramUserId(
	ctx context.Context, telegramUserId string) ([]*response.ImageResponseDto, error) {
	query := "SELECT pi.id, pi.telegram_user_id, pi.name, pi.url, pi.size, pis.is_blocked, pis.is_primary," +
		" pis.is_private, pi.created_at, pi.updated_at" +
		" FROM dating.profile_images pi" +
		" JOIN dating.profile_image_statuses pis ON pi.id = pis.id" +
		" WHERE pi.telegram_user_id = $1 AND pis.is_blocked = false"
	rows, err := r.db.QueryContext(ctx, query, telegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectListByTelegramUserId",
			"QueryContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*response.ImageResponseDto, 0)
	for rows.Next() {
		p := &response.ImageResponseDto{}
		pr := &response.ImageResponseRepositoryDto{}
		err := rows.Scan(&p.Id, &pr.TelegramUserId, &p.Name, &p.Url, &pr.Size, &pr.IsBlocked, &pr.IsPrimary,
			&pr.IsPrivate, &pr.CreatedAt, &pr.UpdatedAt)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectListByTelegramUserId", "Scan")
			r.logger.Debug(errorMessage, zap.Error(err))
			continue
		}
		result := &response.ImageResponseDto{
			Id:   p.Id,
			Name: p.Name,
			Url:  p.Url,
		}
		list = append(list, result)
	}
	return list, nil
}

func (r *ImageRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathImage)
}
