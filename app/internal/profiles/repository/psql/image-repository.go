package psql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
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
	ctx context.Context, p *request.ImageAddRequestRepositoryDto) (*entity.ImageEntity, error) {
	query := "INSERT INTO dating.profile_images (telegram_user_id, name, url, size, is_blocked, is_primary," +
		" is_private, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.Name, &p.Url, &p.Size, &p.IsBlocked,
		&p.IsPrimary, &p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindById(ctx, id)
}

func (r *ImageRepository) Update(
	ctx context.Context, p *request.ImageUpdateRequestRepositoryDto) (*entity.ImageEntity, error) {
	query := "UPDATE dating.profile_images SET name = $1, url = $2, size = $3, is_blocked = $4," +
		" is_primary = $5, is_private = $6, updated_at = $7 WHERE id = $8"
	_, err := r.db.ExecContext(ctx, query, &p.Name, &p.Url, &p.Size, &p.IsBlocked,
		&p.IsPrimary, &p.IsPrivate, &p.UpdatedAt, &p.Id)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindById(ctx, p.Id)
}

func (r *ImageRepository) Delete(
	ctx context.Context, p *request.ImageDeleteRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "DELETE FROM dating.profile_images WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, &p.Id)
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

func (r *ImageRepository) FindById(ctx context.Context, imageId uint64) (*entity.ImageEntity, error) {
	p := &entity.ImageEntity{}
	query := "SELECT id, telegram_user_id, name, url, size, is_blocked, is_primary," +
		" is_private, created_at, updated_at" +
		" FROM dating.profile_images" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, imageId)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.Name, &p.Url, &p.Size, &p.IsBlocked, &p.IsPrimary,
		&p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ImageRepository) FindLastByTelegramUserId(ctx context.Context, telegramUserId string) (*entity.ImageEntity, error) {
	p := &entity.ImageEntity{}
	query := "SELECT id, telegram_user_id, name, url, size, is_blocked, is_primary," +
		" is_private, created_at, updated_at" +
		" FROM dating.profile_images" +
		" WHERE telegram_user_id = $1 AND is_blocked = false AND is_private = false" +
		" ORDER BY id DESC" +
		" LIMIT 1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.Name, &p.Url, &p.Size, &p.IsBlocked, &p.IsPrimary,
		&p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindLastByTelegramUserId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ImageRepository) SelectListAllByTelegramUserId(
	ctx context.Context, telegramUserId string) ([]*entity.ImageEntity, error) {
	query := "SELECT id, telegram_user_id, name, url, size, is_blocked, is_primary," +
		" is_private, created_at, updated_at" +
		" FROM dating.profile_images" +
		" WHERE telegram_user_id = $1"
	rows, err := r.db.QueryContext(ctx, query, telegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectListAllByTelegramUserId",
			"QueryContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*entity.ImageEntity, 0)
	for rows.Next() {
		p := entity.ImageEntity{}
		err := rows.Scan(&p.Id, &p.TelegramUserId, &p.Name, &p.Url, &p.Size, &p.IsBlocked, &p.IsPrimary,
			&p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectListAllByTelegramUserId", "Scan")
			r.logger.Debug(errorMessage, zap.Error(err))
			continue
		}
		list = append(list, &p)
	}
	return list, nil
}

func (r *ImageRepository) SelectListPublicByTelegramUserId(
	ctx context.Context, telegramUserId string) ([]*entity.ImageEntity, error) {
	query := "SELECT id, telegram_user_id, name, url, size, is_blocked, is_primary," +
		" is_private, created_at, updated_at" +
		" FROM dating.profile_images" +
		" WHERE telegram_user_id = $1 AND is_blocked = false AND is_private = false"
	rows, err := r.db.QueryContext(ctx, query, telegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectListPublicByTelegramUserId",
			"QueryContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*entity.ImageEntity, 0)
	for rows.Next() {
		p := entity.ImageEntity{}
		err := rows.Scan(&p.Id, &p.TelegramUserId, &p.Name, &p.Url, &p.Size, &p.IsBlocked, &p.IsPrimary,
			&p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectListPublicByTelegramUserId",
				"Scan")
			r.logger.Debug(errorMessage, zap.Error(err))
			continue
		}
		list = append(list, &p)
	}
	return list, nil
}

func (r *ImageRepository) SelectListByTelegramUserId(
	ctx context.Context, telegramUserId string) ([]*entity.ImageEntity, error) {
	query := "SELECT id, telegram_user_id, name, url, size, is_blocked, is_primary," +
		" is_private, created_at, updated_at" +
		" FROM dating.profile_images" +
		" WHERE telegram_user_id = $1 AND is_blocked = false"
	rows, err := r.db.QueryContext(ctx, query, telegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectListByTelegramUserId",
			"QueryContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*entity.ImageEntity, 0)
	for rows.Next() {
		p := entity.ImageEntity{}
		err := rows.Scan(&p.Id, &p.TelegramUserId, &p.Name, &p.Url, &p.Size, &p.IsBlocked, &p.IsPrimary,
			&p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectListByTelegramUserId", "Scan")
			r.logger.Debug(errorMessage, zap.Error(err))
			continue
		}
		list = append(list, &p)
	}
	return list, nil
}

func (r *ImageRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathImage)
}
