package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"go.uber.org/zap"
	"time"
)

const (
	errorFilePathStatus = "internal/repository/psql/status-repository.go"
)

type StatusRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewStatusRepository(l logger.Logger, db *sql.DB) *StatusRepository {
	return &StatusRepository{
		logger: l,
		db:     db,
	}
}

func (r *StatusRepository) Add(
	ctx context.Context, p *request.StatusAddRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "INSERT INTO dating.profile_statuses (telegram_user_id, is_blocked, is_frozen, is_hidden_age," +
		" is_hidden_distance, is_invisible, is_left_hand, is_online, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.IsBlocked, &p.IsFrozen, &p.IsHiddenAge,
		&p.IsHiddenDistance, &p.IsInvisible, &p.IsLeftHand, &p.IsOnline, &p.CreatedAt, &p.UpdatedAt)
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

func (r *StatusRepository) Update(
	ctx context.Context, p *request.StatusUpdateRequestRepositoryDto) (*entity.StatusEntity, error) {
	query := "UPDATE dating.profile_statuses SET is_blocked = $1, is_frozen = $2, is_hidden_age = $3," +
		" is_hidden_distance = $4, is_invisible = $5, is_left_hand = $6, is_online = $7, updated_at = $8" +
		" WHERE telegram_user_id = $10"
	_, err := r.db.ExecContext(ctx, query, &p.IsBlocked, &p.IsFrozen, &p.IsHiddenAge, &p.IsHiddenDistance,
		&p.IsInvisible, &p.IsLeftHand, &p.IsOnline, &p.UpdatedAt, &p.TelegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindByTelegramUserId(ctx, p.TelegramUserId)
}

func (r *StatusRepository) Block(
	ctx context.Context, telegramUserId string) (*entity.StatusEntity, error) {
	query := "UPDATE dating.profile_statuses SET is_blocked = $1, updated_at = $2 WHERE telegram_user_id = $3"
	updatedAt := time.Now().UTC()
	_, err := r.db.ExecContext(ctx, query, true, updatedAt, telegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("Block", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindByTelegramUserId(ctx, telegramUserId)
}

func (r *StatusRepository) Freeze(
	ctx context.Context, telegramUserId string) (*entity.StatusEntity, error) {
	query := "UPDATE dating.profile_statuses SET is_frozen = $1, updated_at = $2 WHERE telegram_user_id = $3"
	updatedAt := time.Now().UTC()
	_, err := r.db.ExecContext(ctx, query, true, updatedAt, telegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("Freeze", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindByTelegramUserId(ctx, telegramUserId)
}

func (r *StatusRepository) Restore(
	ctx context.Context, telegramUserId string) (*entity.StatusEntity, error) {
	query := "UPDATE dating.profile_statuses SET is_frozen = $1, updated_at = $2 WHERE telegram_user_id = $3"
	updatedAt := time.Now().UTC()
	_, err := r.db.ExecContext(ctx, query, false, updatedAt, telegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("Restore", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindByTelegramUserId(ctx, telegramUserId)
}

func (r *StatusRepository) UpdateSettings(
	ctx context.Context, p *request.StatusUpdateSettingsRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "UPDATE dating.profile_statuses SET is_hidden_age = $1, updated_at = $2 WHERE telegram_user_id = $3"
	updatedAt := time.Now().UTC()
	_, err := r.db.ExecContext(ctx, query, p.IsHiddenAge, updatedAt, p.TelegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateSettings", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	statusResponse := &response.ResponseDto{
		Success: true,
	}
	return statusResponse, nil
}

func (r *StatusRepository) FindByTelegramUserId(
	ctx context.Context, telegramUserId string) (*entity.StatusEntity, error) {
	p := &entity.StatusEntity{}
	query := "SELECT id, telegram_user_id, is_blocked, is_frozen, is_hidden_age, is_hidden_distance," +
		" is_invisible, is_left_hand, is_online, created_at, updated_at" +
		" FROM dating.profile_statuses" +
		" WHERE telegram_user_id = $1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.IsBlocked, &p.IsFrozen, &p.IsHiddenAge, &p.IsHiddenDistance,
		&p.IsInvisible, &p.IsLeftHand, &p.IsOnline, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindByTelegramUserId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *StatusRepository) CheckProfileExists(
	ctx context.Context, telegramUserId string) (*response.CheckExistsDto, error) {
	p := &response.CheckExistsDto{}
	query := "SELECT telegram_user_id, is_frozen" +
		" FROM dating.profile_statuses" +
		" WHERE telegram_user_id = $1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	if row == nil {
		errorMessage := r.getErrorMessage("CheckProfileExists", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	err := row.Scan(&p.TelegramUserId, &p.IsFrozen)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		errorMessage := r.getErrorMessage("CheckProfileExists", "Scan")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	if err != nil {
		errorMessage := r.getErrorMessage("CheckProfileExists", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *StatusRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathStatus)
}
