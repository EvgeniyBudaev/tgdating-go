package psql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"go.uber.org/zap"
)

const (
	errorFilePathTelegram = "internal/repository/psql/telegramRepository.go"
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

func (r *TelegramRepository) Add(
	ctx context.Context, p *request.TelegramAddRequestRepositoryDto) (*entity.TelegramEntity, error) {
	query := "INSERT INTO profile_telegrams (session_id, user_id, username, first_name, last_name, language_code," +
		" allows_write_to_pm, query_id, is_deleted, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.UserId, &p.UserName, &p.FirstName, &p.LastName,
		&p.LanguageCode, &p.AllowsWriteToPm, &p.QueryId, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindById(ctx, id)
}

func (r *TelegramRepository) Update(
	ctx context.Context, p *request.TelegramUpdateRequestRepositoryDto) (*entity.TelegramEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateTelegram", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_telegrams SET user_id=$1, username=$2, first_name=$3, last_name=$4, language_code=$5," +
		" allows_write_to_pm=$6, query_id=$7, updated_at=$8" +
		" WHERE session_id=$9"
	_, err = r.db.ExecContext(ctx, query, &p.UserId, &p.UserName, &p.FirstName, &p.LastName, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryId, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindBySessionId(ctx, p.SessionId)
}

func (r *TelegramRepository) Delete(
	ctx context.Context, p *request.TelegramDeleteRequestRepositoryDto) (*entity.TelegramEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("Delete", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_telegrams SET is_deleted=$1, updated_at=$2 WHERE session_id=$3"
	_, err = r.db.ExecContext(ctx, query, &p.IsDeleted, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("Delete", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindBySessionId(ctx, p.SessionId)
}

func (r *TelegramRepository) FindById(ctx context.Context, id uint64) (*entity.TelegramEntity, error) {
	p := &entity.TelegramEntity{}
	query := "SELECT id, session_id, user_id, username, first_name, last_name, language_code, allows_write_to_pm," +
		" query_id, is_deleted, created_at, updated_at" +
		" FROM profile_telegrams" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&p.Id, &p.SessionId, &p.UserId, &p.UserName, &p.FirstName, &p.LastName, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryId, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *TelegramRepository) FindBySessionId(
	ctx context.Context, sessionID string) (*entity.TelegramEntity, error) {
	p := &entity.TelegramEntity{}
	query := "SELECT id, session_id, user_id, username, first_name, last_name, language_code, allows_write_to_pm," +
		" query_id, is_deleted, created_at, updated_at" +
		" FROM profile_telegrams" +
		" WHERE session_id = $1"
	row := r.db.QueryRowContext(ctx, query, sessionID)
	err := row.Scan(&p.Id, &p.SessionId, &p.UserId, &p.UserName, &p.FirstName, &p.LastName, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryId, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindBySessionId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *TelegramRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathTelegram)
}
