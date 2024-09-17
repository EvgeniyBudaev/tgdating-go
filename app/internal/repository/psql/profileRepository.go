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
	errorFilePath = "internal/repository/psql/profileRepository.go"
)

type ProfileRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewProfileRepository(l logger.Logger, db *sql.DB) *ProfileRepository {
	return &ProfileRepository{
		logger: l,
		db:     db,
	}
}

func (r *ProfileRepository) AddProfile(ctx context.Context, p *entity.ProfileEntity) (*entity.ProfileEntity, error) {
	birthday := p.Birthday.Format("2006-01-02")
	query := "INSERT INTO profiles (session_id, display_name, birthday, gender, location, description," +
		" height, weight, is_deleted, is_blocked, is_premium, is_show_distance, is_invisible," +
		" created_at, updated_at, last_online)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)" +
		" RETURNING id"
	err := r.db.QueryRowContext(ctx, query, &p.SessionID, &p.DisplayName, &birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, p.IsDeleted, &p.IsBlocked, &p.IsPremium, &p.IsShowDistance,
		&p.IsInvisible, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline).Scan(&p.ID)
	if err != nil {
		errorMessage := r.getErrorMessage("AddProfile", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) UpdateProfile(ctx context.Context, p *entity.ProfileEntity) (*entity.ProfileEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateProfile", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profiles SET display_name=$1, birthday=$2, gender=$3, location=$4," +
		" description=$5, height=$6, weight=$7, updated_at=$8, last_online=$9" +
		" WHERE session_id=$10"
	_, err = r.db.ExecContext(ctx, query, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.UpdatedAt, &p.LastOnline, &p.SessionID)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateProfile", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	profileResponse, err := r.FindProfileBySessionID(ctx, p.SessionID)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateProfile", "FindProfileBySessionID")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return profileResponse, nil
}

func (r *ProfileRepository) FindProfileBySessionID(
	ctx context.Context, sessionID string) (*entity.ProfileEntity, error) {
	p := &entity.ProfileEntity{}
	query := "SELECT id, session_id, display_name, birthday, gender, location, description, height, weight," +
		" is_deleted, is_blocked, is_premium, is_show_distance, is_invisible, created_at, updated_at, last_online" +
		" FROM profiles" +
		" WHERE session_id=$1"
	row := r.db.QueryRowContext(ctx, query, sessionID)
	if row == nil {
		err := errors.New("no rows found")
		errorMessage := r.getErrorMessage("FindProfileBySessionID", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	err := row.Scan(&p.ID, &p.SessionID, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.IsDeleted, &p.IsBlocked, &p.IsPremium,
		&p.IsShowDistance, &p.IsInvisible, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline)
	if err != nil {
		errorMessage := r.getErrorMessage("FindProfileBySessionID", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) AddImage(
	ctx context.Context, p *entity.ProfileImageEntity) (*entity.ProfileImageEntity, error) {
	query := "INSERT INTO profile_images (session_id, name, url, size, is_deleted, is_blocked, is_primary," +
		"  is_private, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	fmt.Println("Repository p: ", p)
	err := r.db.QueryRowContext(ctx, query, &p.SessionID, &p.Name, &p.Url, &p.Size, &p.IsDeleted, &p.IsBlocked,
		&p.IsPrimary, &p.IsPrivate, &p.CreatedAt, &p.UpdatedAt).Scan(&p.ID)
	if err != nil {
		errorMessage := r.getErrorMessage("AddImage", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) AddNavigator(
	ctx context.Context, p *entity.ProfileNavigatorEntity) (*entity.ProfileNavigatorEntity, error) {
	query := "INSERT INTO profile_navigators (session_id, location, is_deleted, created_at, updated_at)" +
		" VALUES ($1, ST_SetSRID(ST_MakePoint($2, $3),  4326), $4, $5, $6) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, &p.SessionID, &p.Location.Longitude, &p.Location.Latitude, &p.IsDeleted,
		&p.CreatedAt, &p.UpdatedAt).Scan(&p.ID)
	if err != nil {
		errorMessage := r.getErrorMessage("AddNavigator", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) UpdateNavigator(
	ctx context.Context, p *request.ProfileNavigatorUpdateRequestDto) (*entity.ProfileNavigatorEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateNavigator", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_navigators SET location=ST_SetSRID(ST_MakePoint($1, $2),  4326), updated_at=$3" +
		" WHERE session_id=$4"
	_, err = r.db.ExecContext(ctx, query, &p.Longitude, &p.Latitude, &p.UpdatedAt, &p.SessionID)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateNavigator", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	navigatorResponse, err := r.FindNavigatorBySessionID(ctx, p.SessionID)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateNavigator",
			"FindNavigatorBySessionID")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return navigatorResponse, nil
}

func (r *ProfileRepository) FindNavigatorBySessionID(
	ctx context.Context, sessionID string) (*entity.ProfileNavigatorEntity, error) {
	p := &entity.ProfileNavigatorEntity{}
	var longitude sql.NullFloat64
	var latitude sql.NullFloat64
	query := `SELECT id, session_id, ST_X(location) as longitude, ST_Y(location) as latitude, is_deleted, created_at,
                updated_at
			  FROM profile_navigators
			  WHERE session_id = $1`
	row := r.db.QueryRowContext(ctx, query, sessionID)
	if row == nil {
		err := errors.New("no rows found")
		errorMessage := r.getErrorMessage("FindNavigatorBySessionID",
			"QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	err := row.Scan(&p.ID, &p.SessionID, &longitude, &latitude, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindNavigatorBySessionID", "Scan")
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

func (r *ProfileRepository) AddFilter(
	ctx context.Context, p *entity.ProfileFilterEntity) (*entity.ProfileFilterEntity, error) {
	query := "INSERT INTO profile_filters (session_id, search_gender, looking_for, age_from, age_to, distance, page," +
		" size, is_deleted, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, &p.SessionID, &p.SearchGender, &p.LookingFor, &p.AgeFrom, &p.AgeTo,
		&p.Distance, &p.Page, &p.Size, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt).Scan(&p.ID)
	if err != nil {
		errorMessage := r.getErrorMessage("AddFilter", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) AddTelegram(
	ctx context.Context, p *entity.ProfileTelegramEntity) (*entity.ProfileTelegramEntity, error) {
	query := "INSERT INTO profile_telegrams (session_id, user_id, username, first_name, last_name, language_code," +
		" allows_write_to_pm, query_id, chat_id, is_deleted, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, &p.SessionID, &p.UserID, &p.UserName, &p.Firstname, &p.Lastname,
		&p.LanguageCode, &p.AllowsWriteToPm, &p.QueryID, &p.ChatID, &p.IsDeleted, &p.CreatedAt,
		&p.UpdatedAt).Scan(&p.ID)
	if err != nil {
		errorMessage := r.getErrorMessage("AddTelegram", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) UpdateTelegram(
	ctx context.Context, p *entity.ProfileTelegramEntity) (*entity.ProfileTelegramEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateTelegram", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_telegrams SET username=$1, first_name=$2, last_name=$3, language_code=$4," +
		" allows_write_to_pm=$5, updated_at=$6" +
		" WHERE session_id=$7"
	_, err = r.db.ExecContext(ctx, query, &p.UserName, &p.Firstname, &p.Lastname, &p.LanguageCode, &p.AllowsWriteToPm,
		&p.UpdatedAt, &p.SessionID)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateTelegram", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	telegramResponse, err := r.FindTelegramBySessionID(ctx, p.SessionID)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateTelegram", "FindTelegramBySessionID")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return telegramResponse, nil
}

func (r *ProfileRepository) FindTelegramBySessionID(
	ctx context.Context, sessionID string) (*entity.ProfileTelegramEntity, error) {
	p := &entity.ProfileTelegramEntity{}
	query := "SELECT id, session_id, user_id, username, first_name, last_name, language_code, allows_write_to_pm," +
		" query_id, chat_id, is_deleted, created_at, updated_at" +
		" FROM profile_telegrams" +
		" WHERE session_id = $1"
	row := r.db.QueryRowContext(ctx, query, sessionID)
	if row == nil {
		err := errors.New("no rows found")
		errorMessage := r.getErrorMessage("FindTelegramBySessionID",
			"QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	err := row.Scan(&p.ID, &p.SessionID, &p.UserID, &p.UserName, &p.Firstname, &p.Lastname, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryID, &p.ChatID, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindTelegramBySessionID", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}
