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
	ctx context.Context, p *request.FilterAddRequestRepositoryDto) (*entity.FilterEntity, error) {
	query := "INSERT INTO dating.profile_filters (session_id, search_gender, looking_for, age_from, age_to, distance," +
		" page, size, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.SearchGender, &p.LookingFor, &p.AgeFrom, &p.AgeTo,
		&p.Distance, &p.Page, &p.Size, &p.CreatedAt, &p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindById(ctx, id)
}

func (r *FilterRepository) Update(
	ctx context.Context, p *request.FilterUpdateRequestRepositoryDto) (*entity.FilterEntity, error) {
	query := "UPDATE dating.profile_filters SET search_gender=$1, looking_for=$2, age_from=$3, age_to=$4," +
		" updated_at=$5" +
		" WHERE session_id=$6"
	_, err := r.db.ExecContext(ctx, query, &p.SearchGender, &p.LookingFor, &p.AgeFrom, &p.AgeTo, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindBySessionId(ctx, p.SessionId)
}

func (r *FilterRepository) FindById(
	ctx context.Context, id uint64) (*entity.FilterEntity, error) {
	p := &entity.FilterEntity{}
	query := "SELECT id, session_id, search_gender, looking_for, age_from, age_to, distance, page, size," +
		" created_at, updated_at" +
		" FROM dating.profile_filters" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&p.Id, &p.SessionId, &p.SearchGender, &p.LookingFor, &p.AgeFrom, &p.AgeTo, &p.Distance, &p.Page,
		&p.Size, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *FilterRepository) FindBySessionId(
	ctx context.Context, sessionId string) (*entity.FilterEntity, error) {
	p := &entity.FilterEntity{}
	query := "SELECT id, session_id, search_gender, looking_for, age_from, age_to, distance, page, size," +
		" created_at, updated_at" +
		" FROM dating.profile_filters" +
		" WHERE session_id = $1"
	row := r.db.QueryRowContext(ctx, query, sessionId)
	err := row.Scan(&p.Id, &p.SessionId, &p.SearchGender, &p.LookingFor, &p.AgeFrom, &p.AgeTo, &p.Distance, &p.Page,
		&p.Size, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindBySessionId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *FilterRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathFilter)
}