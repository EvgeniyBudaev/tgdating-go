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

func (r *NavigatorRepository) Add(
	ctx context.Context, p *request.NavigatorAddRequestRepositoryDto) (*entity.NavigatorEntity, error) {
	query := "INSERT INTO dating.profile_navigators (session_id, location, created_at, updated_at)" +
		" VALUES ($1, ST_SetSRID(ST_MakePoint($2, $3),  4326), $4, $5) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.Location.Longitude, &p.Location.Latitude,
		&p.CreatedAt, &p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindById(ctx, id)
}

func (r *NavigatorRepository) Update(
	ctx context.Context, p *request.NavigatorUpdateRequestRepositoryDto) (*entity.NavigatorEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE dating.profile_navigators SET location=ST_SetSRID(ST_MakePoint($1, $2),  4326), updated_at=$3" +
		" WHERE session_id=$4"
	_, err = r.db.ExecContext(ctx, query, &p.Longitude, &p.Latitude, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindBySessionId(ctx, p.SessionId)
}

func (r *NavigatorRepository) FindById(
	ctx context.Context, id uint64) (*entity.NavigatorEntity, error) {
	p := &entity.NavigatorEntity{}
	var longitude sql.NullFloat64
	var latitude sql.NullFloat64
	query := "SELECT id, session_id, ST_X(location) as longitude, ST_Y(location) as latitude, created_at," +
		" updated_at" +
		" FROM dating.profile_navigators" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&p.Id, &p.SessionId, &longitude, &latitude, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindById", "Scan")
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

func (r *NavigatorRepository) FindBySessionId(
	ctx context.Context, sessionId string) (*entity.NavigatorEntity, error) {
	p := &entity.NavigatorEntity{}
	var longitude sql.NullFloat64
	var latitude sql.NullFloat64
	query := "SELECT id, session_id, ST_X(location) as longitude, ST_Y(location) as latitude, created_at," +
		" updated_at" +
		" FROM dating.profile_navigators" +
		" WHERE session_id = $1"
	row := r.db.QueryRowContext(ctx, query, sessionId)
	err := row.Scan(&p.Id, &p.SessionId, &longitude, &latitude, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		errorMessage := r.getErrorMessage("FindBySessionId", "Scan")
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
	query := "SELECT id, session_id, ST_X(location) as longitude, ST_Y(location) as latitude, created_at," +
		" updated_at," +
		" ST_DistanceSphere(ST_SetSRID(ST_MakePoint($4, $5),  4326)," +
		" ST_SetSRID(ST_MakePoint($2, $3),  4326)) as distance" +
		" FROM dating.profile_navigators" +
		" WHERE session_id = $1"
	row := r.db.QueryRowContext(ctx, query, sessionId, longitudeSession, latitudeSession, longitudeViewed,
		latitudeViewed)
	err := row.Scan(&p.Id, &p.SessionId, &longitude, &latitude, &p.CreatedAt, &p.UpdatedAt, &p.Distance)
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
