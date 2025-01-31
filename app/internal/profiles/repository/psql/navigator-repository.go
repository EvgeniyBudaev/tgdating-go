package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	//"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"go.uber.org/zap"
)

const (
	errorFilePathNavigator = "internal/repository/psql/navigator-repository.go"
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
	ctx context.Context, p *request.NavigatorAddRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "INSERT INTO dating.profile_navigators (telegram_user_id, country_code, country_name, city, location," +
		" created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, ST_SetSRID(ST_MakePoint($5, $6),  4326), $7, $8) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.CountryCode, &p.CountryName, &p.City,
		&p.Location.Longitude, &p.Location.Latitude, &p.CreatedAt, &p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	navigatorResponse := &response.ResponseDto{
		Success: true,
	}
	return navigatorResponse, nil
}

func (r *NavigatorRepository) Update(
	ctx context.Context, p *request.NavigatorUpdateRequestRepositoryDto) (*entity.NavigatorEntity, error) {
	query := "UPDATE dating.profile_navigators SET country_code = $1, country_name = $2, city = $3," +
		" location=ST_SetSRID(ST_MakePoint($4, $5),  4326), updated_at = $6" +
		" WHERE telegram_user_id = $7"
	_, err := r.db.ExecContext(ctx, query, &p.CountryCode, &p.CountryName, &p.City, &p.Longitude, &p.Latitude,
		&p.UpdatedAt, &p.TelegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindByTelegramUserId(ctx, p.TelegramUserId)
}

func (r *NavigatorRepository) FindById(
	ctx context.Context, id uint64) (*entity.NavigatorEntity, error) {
	p := &entity.NavigatorEntity{}
	var longitude sql.NullFloat64
	var latitude sql.NullFloat64
	query := "SELECT id, telegram_user_id, country_code, country_name, city, ST_X(location) as longitude," +
		" ST_Y(location) as latitude, created_at, updated_at" +
		" FROM dating.profile_navigators" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.CountryCode, &p.CountryName, &p.City, &longitude, &latitude,
		&p.CreatedAt, &p.UpdatedAt)
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

func (r *NavigatorRepository) FindByTelegramUserId(
	ctx context.Context, telegramUserId string) (*entity.NavigatorEntity, error) {
	p := &entity.NavigatorEntity{}
	var longitude sql.NullFloat64
	var latitude sql.NullFloat64
	query := "SELECT id, telegram_user_id, country_code, country_name, city, ST_X(location) as longitude," +
		" ST_Y(location) as latitude, created_at, updated_at" +
		" FROM dating.profile_navigators" +
		" WHERE telegram_user_id = $1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.CountryCode, &p.CountryName, &p.City, &longitude, &latitude,
		&p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		errorMessage := r.getErrorMessage("FindByTelegramUserId", "Scan")
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

func (r *NavigatorRepository) CheckNavigatorExists(
	ctx context.Context, telegramUserId string) (*response.ResponseDto, error) {
	var existingRecordCount int
	query := "SELECT COUNT(*) FROM dating.profile_navigators WHERE telegram_user_id = $1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	if row == nil {
		errorMessage := r.getErrorMessage("CheckNavigatorExists", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	err := row.Scan(&existingRecordCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		errorMessage := r.getErrorMessage("CheckNavigatorExists", "Scan")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	if existingRecordCount == 0 {
		return nil, nil
	}
	navigatorResponse := &response.ResponseDto{
		Success: true,
	}
	return navigatorResponse, nil
}

func (r *NavigatorRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathNavigator)
}
