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
	errorFilePathTelegram = "internal/repository/psql/telegram-repository.go"
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
	query := "INSERT INTO dating.profile_telegrams (user_id, username, first_name, last_name," +
		" language_code, allows_write_to_pm, query_id, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.UserId, &p.UserName, &p.FirstName, &p.LastName,
		&p.LanguageCode, &p.AllowsWriteToPm, &p.QueryId, &p.CreatedAt, &p.UpdatedAt)
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
	query := "UPDATE dating.profile_telegrams SET username=$1, first_name=$2, last_name=$3," +
		" language_code=$4, allows_write_to_pm=$5, query_id=$6, updated_at=$7" +
		" WHERE user_id=$8"
	_, err := r.db.ExecContext(ctx, query, &p.UserName, &p.FirstName, &p.LastName, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryId, &p.UpdatedAt, &p.UserId)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindByTelegramUserId(ctx, p.UserId)
}

func (r *TelegramRepository) FindById(ctx context.Context, id uint64) (*entity.TelegramEntity, error) {
	p := &entity.TelegramEntity{}
	query := "SELECT id, user_id, username, first_name, last_name, language_code, allows_write_to_pm," +
		" query_id, created_at, updated_at" +
		" FROM dating.profile_telegrams" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&p.Id, &p.UserId, &p.UserName, &p.FirstName, &p.LastName, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryId, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *TelegramRepository) FindByTelegramUserId(
	ctx context.Context, telegramUserId string) (*entity.TelegramEntity, error) {
	p := &entity.TelegramEntity{}
	query := "SELECT id, user_id, username, first_name, last_name, language_code, allows_write_to_pm," +
		" query_id, created_at, updated_at" +
		" FROM dating.profile_telegrams" +
		" WHERE user_id = $1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	err := row.Scan(&p.Id, &p.UserId, &p.UserName, &p.FirstName, &p.LastName, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryId, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindByTelegramUserId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *TelegramRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathTelegram)
}
