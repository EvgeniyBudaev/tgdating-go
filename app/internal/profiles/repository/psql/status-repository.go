package psql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
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
	ctx context.Context, p *request.StatusAddRequestRepositoryDto) (*entity.StatusEntity, error) {
	query := "INSERT INTO dating.profile_statuses (telegram_user_id, is_frozen, is_blocked, is_premium," +
		" is_show_distance, is_invisible)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.IsFrozen, &p.IsBlocked, &p.IsPremium,
		&p.IsShowDistance, &p.IsInvisible, &p.CreatedAt, &p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindById(ctx, id)
}

func (r *StatusRepository) Update(
	ctx context.Context, p *request.StatusUpdateRequestRepositoryDto) (*entity.StatusEntity, error) {
	query := "UPDATE dating.profile_statuses SET is_frozen = $1, is_blocked = $2, is_premium = $3," +
		" is_show_distance = $4, is_invisible = $5, updated_at = $6" +
		" WHERE telegram_user_id = $7"
	_, err := r.db.ExecContext(ctx, query, &p.IsFrozen, &p.IsBlocked, &p.IsPremium, &p.IsShowDistance, &p.IsInvisible,
		&p.UpdatedAt, &p.TelegramUserId)
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
	_, err := r.db.ExecContext(ctx, query, true, time.Now().UTC(), telegramUserId)
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
	_, err := r.db.ExecContext(ctx, query, true, time.Now().UTC(), telegramUserId)
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
	_, err := r.db.ExecContext(ctx, query, false, time.Now().UTC(), telegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("Restore", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindByTelegramUserId(ctx, telegramUserId)
}

func (r *StatusRepository) FindById(ctx context.Context, id uint64) (*entity.StatusEntity, error) {
	p := &entity.StatusEntity{}
	query := "SELECT id, telegram_user_id, is_frozen, is_blocked, is_premium, is_show_distance, is_invisible," +
		" created_at, updated_at" +
		" FROM dating.profile_statuses" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.IsFrozen, &p.IsBlocked, &p.IsPremium, &p.IsShowDistance,
		&p.IsInvisible, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *StatusRepository) FindByTelegramUserId(
	ctx context.Context, telegramUserId string) (*entity.StatusEntity, error) {
	p := &entity.StatusEntity{}
	query := "SELECT id, telegram_user_id, is_frozen, is_blocked, is_premium, is_show_distance, is_invisible," +
		" created_at, updated_at" +
		" FROM dating.profile_statuses" +
		" WHERE telegram_user_id = $1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.IsFrozen, &p.IsBlocked, &p.IsPremium, &p.IsShowDistance,
		&p.IsInvisible, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindByTelegramUserId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *StatusRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathStatus)
}
