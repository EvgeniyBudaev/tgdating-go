package psql

import (
	"context"
	"database/sql"
	"fmt"
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
		" created_at, updated_at, last_online) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14," +
		" $15, $16) RETURNING id"
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

func (r *ProfileRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}
