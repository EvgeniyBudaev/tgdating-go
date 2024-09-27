package psql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"go.uber.org/zap"
)

const (
	errorFilePathNavigator = "internal/repository/psql/navigatorRepository.go"
)

type NavigatorRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewNavigatorRepository(l logger.Logger, db *sql.DB) *NavigatorRepository {
	return &NavigatorRepository{
		logger: l,
		db:     db,
	}
}

func (r *NavigatorRepository) AddNavigator(
	ctx context.Context, p *request.NavigatorAddRequestRepositoryDto) (*entity.NavigatorEntity, error) {
	query := "INSERT INTO profile_navigators (session_id, location, is_deleted, created_at, updated_at)" +
		" VALUES ($1, ST_SetSRID(ST_MakePoint($2, $3),  4326), $4, $5, $6) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.Location.Longitude, &p.Location.Latitude, &p.IsDeleted,
		&p.CreatedAt, &p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("AddNavigator", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindNavigatorById(ctx, id)
}

func (r *NavigatorRepository) UpdateNavigator(
	ctx context.Context, p *request.NavigatorUpdateRequestRepositoryDto) (*entity.NavigatorEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateNavigator", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_navigators SET location=ST_SetSRID(ST_MakePoint($1, $2),  4326), updated_at=$3" +
		" WHERE session_id=$4"
	_, err = r.db.ExecContext(ctx, query, &p.Longitude, &p.Latitude, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateNavigator", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindNavigatorBySessionId(ctx, p.SessionId)
}

func (r *NavigatorRepository) DeleteNavigator(
	ctx context.Context, p *request.NavigatorDeleteRequestRepositoryDto) (*entity.NavigatorEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteNavigator", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_navigators SET is_deleted=$1, updated_at=$2 WHERE session_id=$3"
	_, err = r.db.ExecContext(ctx, query, &p.IsDeleted, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteNavigator", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindNavigatorBySessionId(ctx, p.SessionId)
}

func (r *NavigatorRepository) FindNavigatorById(
	ctx context.Context, id uint64) (*entity.NavigatorEntity, error) {
	p := &entity.NavigatorEntity{}
	var longitude sql.NullFloat64
	var latitude sql.NullFloat64
	query := "SELECT id, session_id, ST_X(location) as longitude, ST_Y(location) as latitude, is_deleted, created_at," +
		" updated_at" +
		" FROM profile_navigators" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&p.Id, &p.SessionId, &longitude, &latitude, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindNavigatorById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if !longitude.Valid && !latitude.Valid {
		return nil, err
	}
	p.Location = &entity.PointEntity{
		Latitude:  latitude.Float64,
		Longitude: longitude.Float64,
	}
	return p, nil
}

func (r *NavigatorRepository) FindNavigatorBySessionId(
	ctx context.Context, sessionId string) (*entity.NavigatorEntity, error) {
	p := &entity.NavigatorEntity{}
	var longitude sql.NullFloat64
	var latitude sql.NullFloat64
	query := "SELECT id, session_id, ST_X(location) as longitude, ST_Y(location) as latitude, is_deleted, created_at," +
		" updated_at" +
		" FROM profile_navigators" +
		" WHERE session_id = $1"
	row := r.db.QueryRowContext(ctx, query, sessionId)
	err := row.Scan(&p.Id, &p.SessionId, &longitude, &latitude, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindNavigatorBySessionId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if !longitude.Valid && !latitude.Valid {
		return nil, err
	}
	p.Location = &entity.PointEntity{
		Latitude:  latitude.Float64,
		Longitude: longitude.Float64,
	}
	return p, nil
}

func (r *NavigatorRepository) FindDistance(ctx context.Context, pe *entity.NavigatorEntity,
	pve *entity.NavigatorEntity) (*response.NavigatorDistanceResponseRepositoryDto, error) {
	sessionId := pe.SessionId
	longitudeSession := pe.Location.Longitude
	latitudeSession := pe.Location.Latitude
	longitudeViewed := pve.Location.Longitude
	latitudeViewed := pve.Location.Latitude
	p := &response.NavigatorDistanceResponseRepositoryDto{}
	var longitude sql.NullFloat64
	var latitude sql.NullFloat64
	query := "SELECT id, session_id, ST_X(location) as longitude, ST_Y(location) as latitude, is_deleted, created_at," +
		" updated_at," +
		" ST_DistanceSphere(ST_SetSRID(ST_MakePoint($4, $5),  4326)," +
		" ST_SetSRID(ST_MakePoint($2, $3),  4326)) as distance" +
		" FROM profile_navigators" +
		" WHERE session_id = $1"
	row := r.db.QueryRowContext(ctx, query, sessionId, longitudeSession, latitudeSession, longitudeViewed,
		latitudeViewed)
	err := row.Scan(&p.Id, &p.SessionId, &longitude, &latitude, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt, &p.Distance)
	if err != nil {
		errorMessage := r.getErrorMessage("FindDistance", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	if !longitude.Valid && !latitude.Valid {
		return nil, err
	}
	p.Location = &entity.PointEntity{
		Latitude:  latitude.Float64,
		Longitude: longitude.Float64,
	}
	return p, nil
}

func (r *NavigatorRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathNavigator)
}
