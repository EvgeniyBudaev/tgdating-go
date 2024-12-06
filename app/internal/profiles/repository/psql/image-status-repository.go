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
	errorFilePathImageStatus = "internal/repository/psql/image-status-repository.go"
)

type ImageStatusRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewImageStatusRepository(l logger.Logger, db *sql.DB) *ImageStatusRepository {
	return &ImageStatusRepository{
		logger: l,
		db:     db,
	}
}

func (r *ImageStatusRepository) Add(
	ctx context.Context, p *request.ImageStatusAddRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "INSERT INTO dating.profile_image_statuses (image_id, is_blocked, is_primary, is_private," +
		" created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.ImageId, &p.IsBlocked, &p.IsPrimary, &p.IsPrivate, &p.CreatedAt,
		&p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	imageStatusResponse := &response.ResponseDto{
		Success: true,
	}
	return imageStatusResponse, nil
}

func (r *ImageStatusRepository) FindById(ctx context.Context, id uint64) (*entity.ImageStatusEntity, error) {
	p := &entity.ImageStatusEntity{}
	query := "SELECT id, is_blocked, is_primary, is_private, created_at, updated_at" +
		" FROM dating.profile_image_statuses" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&p.Id, &p.IsBlocked, &p.IsPrimary, &p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ImageStatusRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathImageStatus)
}
