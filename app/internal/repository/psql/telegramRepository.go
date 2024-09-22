package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"go.uber.org/zap"
)

const (
	errorFilePathTelegram = "internal/repository/psql/telegramRepository.go"
)

var (
	ErrNotRowsFoundTelegram = errors.New("no rows found")
)

type TelegramRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewTelegramRepository(l logger.Logger, db *sql.DB) *TelegramRepository {
	return &TelegramRepository{
		logger: l,
		db:     db,
	}
}

func (r *TelegramRepository) AddTelegram(
	ctx context.Context, p *request.TelegramAddRequestRepositoryDto) (*entity.TelegramEntity, error) {
	query := "INSERT INTO profile_telegrams (session_id, user_id, username, first_name, last_name, language_code," +
		" allows_write_to_pm, query_id, chat_id, is_deleted, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.UserId, &p.UserName, &p.FirstName, &p.LastName,
		&p.LanguageCode, &p.AllowsWriteToPm, &p.QueryId, &p.ChatId, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if row == nil {
		errorMessage := r.getErrorMessage("AddTelegram", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFoundTelegram))
		return nil, ErrNotRowsFoundTelegram
	}
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("AddTelegram", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindTelegramById(ctx, id)
}

func (r *TelegramRepository) UpdateTelegram(
	ctx context.Context, p *request.TelegramUpdateRequestRepositoryDto) (*entity.TelegramEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateTelegram", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_telegrams SET user_id=$1, username=$2, first_name=$3, last_name=$4, language_code=$5," +
		" allows_write_to_pm=$6, query_id=$7, chat_id=$8, updated_at=$9" +
		" WHERE session_id=$10"
	_, err = r.db.ExecContext(ctx, query, &p.UserId, &p.UserName, &p.FirstName, &p.LastName, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryId, &p.ChatId, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateTelegram", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindTelegramBySessionId(ctx, p.SessionId)
}

func (r *TelegramRepository) DeleteTelegram(
	ctx context.Context, p *request.TelegramDeleteRequestRepositoryDto) (*entity.TelegramEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteTelegram", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_telegrams SET is_deleted=$1, updated_at=$2 WHERE session_id=$3"
	_, err = r.db.ExecContext(ctx, query, &p.IsDeleted, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteTelegram", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindTelegramBySessionId(ctx, p.SessionId)
}

func (r *TelegramRepository) FindTelegramById(ctx context.Context, id uint64) (*entity.TelegramEntity, error) {
	p := &entity.TelegramEntity{}
	query := "SELECT id, session_id, user_id, username, first_name, last_name, language_code, allows_write_to_pm," +
		" query_id, chat_id, is_deleted, created_at, updated_at" +
		" FROM profile_telegrams" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindTelegramById", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFoundTelegram))
		return nil, ErrNotRowsFoundTelegram
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.UserId, &p.UserName, &p.FirstName, &p.LastName, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryId, &p.ChatId, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindTelegramById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *TelegramRepository) FindTelegramBySessionId(
	ctx context.Context, sessionID string) (*entity.TelegramEntity, error) {
	p := &entity.TelegramEntity{}
	query := "SELECT id, session_id, user_id, username, first_name, last_name, language_code, allows_write_to_pm," +
		" query_id, chat_id, is_deleted, created_at, updated_at" +
		" FROM profile_telegrams" +
		" WHERE session_id = $1"
	row := r.db.QueryRowContext(ctx, query, sessionID)
	if row == nil {
		errorMessage := r.getErrorMessage("FindTelegramBySessionId",
			"QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFoundTelegram))
		return nil, ErrNotRowsFoundTelegram
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.UserId, &p.UserName, &p.FirstName, &p.LastName, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryId, &p.ChatId, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindTelegramBySessionId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *TelegramRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathTelegram)
}
