package psql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"go.uber.org/zap"
)

const (
	errorFilePathImage = "internal/repository/psql/imageRepository.go"
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

func (r *ImageRepository) AddImage(
	ctx context.Context, p *request.ImageAddRequestRepositoryDto) (*entity.ImageEntity, error) {
	query := "INSERT INTO profile_images (session_id, name, url, size, is_deleted, is_blocked, is_primary," +
		" is_private, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.Name, &p.Url, &p.Size, &p.IsDeleted, &p.IsBlocked,
		&p.IsPrimary, &p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("AddImage", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindImageById(ctx, id)
}

func (r *ImageRepository) UpdateImage(
	ctx context.Context, p *request.ImageUpdateRequestRepositoryDto) (*entity.ImageEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateImage", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_images SET name=$1, url=$2, size=$3, is_deleted=$4, is_blocked=$5," +
		" is_primary=$6, is_private=$7, updated_at=$8 WHERE id=$9"
	_, err = r.db.ExecContext(ctx, query, &p.Name, &p.Url, &p.Size, &p.IsDeleted, &p.IsBlocked,
		&p.IsPrimary, &p.IsPrivate, &p.UpdatedAt, &p.Id)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateImage", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindImageById(ctx, p.Id)
}

func (r *ImageRepository) DeleteImage(
	ctx context.Context, p *request.ImageDeleteRequestRepositoryDto) (*entity.ImageEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteImage", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_images SET is_deleted=$1, updated_at=$2 WHERE id=$3"
	_, err = r.db.ExecContext(ctx, query, &p.IsDeleted, &p.UpdatedAt, &p.Id)
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteImage", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindImageById(ctx, p.Id)
}

func (r *ImageRepository) FindImageById(ctx context.Context, imageId uint64) (*entity.ImageEntity, error) {
	p := &entity.ImageEntity{}
	query := "SELECT id, session_id, name, url, size, is_deleted, is_blocked, is_primary," +
		" is_private, created_at, updated_at" +
		" FROM profile_images" +
		" WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, imageId)
	err := row.Scan(&p.Id, &p.SessionId, &p.Name, &p.Url, &p.Size, &p.IsDeleted, &p.IsBlocked, &p.IsPrimary,
		&p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindImageById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ImageRepository) FindLastImageBySessionId(ctx context.Context, sessionId string) (*entity.ImageEntity, error) {
	p := &entity.ImageEntity{}
	query := "SELECT id, session_id, name, url, size, is_deleted, is_blocked, is_primary," +
		" is_private, created_at, updated_at" +
		" FROM profile_images" +
		" WHERE session_id = $1" +
		" ORDER BY id DESC" +
		" LIMIT 1"
	row := r.db.QueryRowContext(ctx, query, sessionId)
	err := row.Scan(&p.Id, &p.SessionId, &p.Name, &p.Url, &p.Size, &p.IsDeleted, &p.IsBlocked, &p.IsPrimary,
		&p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindLastImageBySessionId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ImageRepository) SelectImageListPublicBySessionId(
	ctx context.Context, sessionId string) ([]*entity.ImageEntity, error) {
	query := "SELECT id, session_id, name, url, size, is_deleted, is_blocked, is_primary," +
		" is_private, created_at, updated_at" +
		" FROM profile_images" +
		" WHERE session_id=$1 AND is_deleted=false AND is_blocked=false AND is_private=false"
	rows, err := r.db.QueryContext(ctx, query, sessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectImageListPublicBySessionId",
			"QueryContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*entity.ImageEntity, 0)
	for rows.Next() {
		p := entity.ImageEntity{}
		err := rows.Scan(&p.Id, &p.SessionId, &p.Name, &p.Url, &p.Size, &p.IsDeleted, &p.IsBlocked, &p.IsPrimary,
			&p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectImageListPublicBySessionId",
				"Scan")
			r.logger.Debug(errorMessage, zap.Error(err))
			continue
		}
		list = append(list, &p)
	}
	return list, nil
}

func (r *ImageRepository) SelectImageListBySessionId(
	ctx context.Context, sessionId string) ([]*entity.ImageEntity, error) {
	query := "SELECT id, session_id, name, url, size, is_deleted, is_blocked, is_primary," +
		" is_private, created_at, updated_at" +
		" FROM profile_images" +
		" WHERE session_id=$1 AND is_deleted=false AND is_blocked=false"
	rows, err := r.db.QueryContext(ctx, query, sessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectImageListBySessionId",
			"QueryContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	list := make([]*entity.ImageEntity, 0)
	for rows.Next() {
		p := entity.ImageEntity{}
		err := rows.Scan(&p.Id, &p.SessionId, &p.Name, &p.Url, &p.Size, &p.IsDeleted, &p.IsBlocked, &p.IsPrimary,
			&p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectImageListBySessionId", "Scan")
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
