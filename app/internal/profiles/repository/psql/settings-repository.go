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
	errorFilePathSettings = "internal/repository/psql/settings-repository.go"
)

type SettingsRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewSettingsRepository(l logger.Logger, db *sql.DB) *SettingsRepository {
	return &SettingsRepository{
		logger: l,
		db:     db,
	}
}

func (r *SettingsRepository) Add(
	ctx context.Context, p *request.SettingsAddRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "INSERT INTO dating.profile_settings (telegram_user_id, measurement, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.Measurement, &p.CreatedAt, &p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	statusResponse := &response.ResponseDto{
		Success: true,
	}
	return statusResponse, nil
}

func (r *SettingsRepository) Update(
	ctx context.Context, p *request.SettingsUpdateRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "UPDATE dating.profile_settings SET measurement = $1, updated_at = $2" +
		" WHERE telegram_user_id = $3"
	_, err := r.db.ExecContext(ctx, query, &p.Measurement, &p.UpdatedAt, &p.TelegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	statusResponse := &response.ResponseDto{
		Success: true,
	}
	return statusResponse, nil
}

func (r *SettingsRepository) FindByTelegramUserId(
	ctx context.Context, telegramUserId string) (*entity.SettingsEntity, error) {
	p := &entity.SettingsEntity{}
	query := "SELECT id, telegram_user_id, measurement, created_at, updated_at" +
		" FROM dating.profile_settings" +
		" WHERE telegram_user_id = $1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.Measurement, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindByTelegramUserId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *SettingsRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathSettings)
}
