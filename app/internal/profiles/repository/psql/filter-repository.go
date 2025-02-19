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
	errorFilePathFilter = "internal/repository/psql/filter-repository.go"
)

type FilterRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewFilterRepository(l logger.Logger, db *sql.DB) *FilterRepository {
	return &FilterRepository{
		logger: l,
		db:     db,
	}
}

func (r *FilterRepository) Add(
	ctx context.Context, p *request.FilterAddRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "INSERT INTO dating.profile_filters (telegram_user_id, search_gender, age_from, age_to," +
		" distance, page, size, is_liked, is_online, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.SearchGender, &p.AgeFrom, &p.AgeTo,
		&p.Distance, &p.Page, &p.Size, &p.IsLiked, &p.IsOnline, &p.CreatedAt, &p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	filterResponse := &response.ResponseDto{
		Success: true,
	}
	return filterResponse, nil
}

func (r *FilterRepository) Update(
	ctx context.Context, p *request.FilterUpdateRequestRepositoryDto) (*entity.FilterEntity, error) {
	query := "UPDATE dating.profile_filters SET search_gender = $1, age_from = $2, age_to = $3," +
		" distance = $4, is_liked = $5, is_online = $6, updated_at = $7" +
		" WHERE telegram_user_id = $8"
	_, err := r.db.ExecContext(ctx, query, &p.SearchGender, &p.AgeFrom, &p.AgeTo,
		&p.Distance, &p.IsLiked, &p.IsOnline, &p.UpdatedAt, &p.TelegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindByTelegramUserId(ctx, p.TelegramUserId)
}

func (r *FilterRepository) FindById(
	ctx context.Context, id uint64) (*entity.FilterEntity, error) {
	p := &entity.FilterEntity{}
	query := "SELECT id, telegram_user_id, search_gender, age_from, age_to, distance, page, size," +
		" is_liked, is_online, created_at, updated_at" +
		" FROM dating.profile_filters" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.SearchGender, &p.AgeFrom, &p.AgeTo, &p.Distance, &p.Page, &p.Size,
		&p.IsLiked, &p.IsOnline, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *FilterRepository) FindByTelegramUserId(
	ctx context.Context, telegramUserId string) (*entity.FilterEntity, error) {
	p := &entity.FilterEntity{}
	query := "SELECT id, telegram_user_id, search_gender, age_from, age_to, distance, page, size," +
		" is_liked, is_online, created_at, updated_at" +
		" FROM dating.profile_filters" +
		" WHERE telegram_user_id = $1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.SearchGender, &p.AgeFrom, &p.AgeTo, &p.Distance, &p.Page, &p.Size,
		&p.IsLiked, &p.IsOnline, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindByTelegramUserId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *FilterRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathFilter)
}
