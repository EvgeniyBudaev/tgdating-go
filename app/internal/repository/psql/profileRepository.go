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

var (
	ErrNotRowsFound = errors.New("no rows found")
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

func (r *ProfileRepository) AddProfile(
	ctx context.Context, p *request.ProfileAddRequestRepositoryDto) (*entity.ProfileEntity, error) {
	birthday := p.Birthday.Format("2006-01-02")
	query := "INSERT INTO profiles (session_id, display_name, birthday, gender, location, description," +
		" height, weight, is_deleted, is_blocked, is_premium, is_show_distance, is_invisible," +
		" created_at, updated_at, last_online)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.DisplayName, &birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, p.IsDeleted, &p.IsBlocked, &p.IsPremium, &p.IsShowDistance,
		&p.IsInvisible, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline)
	if row == nil {
		errorMessage := r.getErrorMessage("AddProfile", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("AddProfile", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindProfileById(ctx, id)
}

func (r *ProfileRepository) UpdateProfile(
	ctx context.Context, p *request.ProfileUpdateRequestRepositoryDto) (*entity.ProfileEntity, error) {
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
		&p.Description, &p.Height, &p.Weight, &p.UpdatedAt, &p.LastOnline, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateProfile", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindProfileBySessionId(ctx, p.SessionId)
}

func (r *ProfileRepository) DeleteProfile(
	ctx context.Context, p *request.ProfileDeleteRequestRepositoryDto) (*entity.ProfileEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteProfile", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profiles SET is_deleted=$1, updated_at=$2, last_online=$3 WHERE session_id=$4"
	_, err = r.db.ExecContext(ctx, query, &p.IsDeleted, &p.UpdatedAt, &p.LastOnline, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteProfile", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindProfileBySessionId(ctx, p.SessionId)
}

func (r *ProfileRepository) FindProfileById(ctx context.Context, id uint64) (*entity.ProfileEntity, error) {
	p := &entity.ProfileEntity{}
	query := "SELECT id, session_id, display_name, birthday, gender, location, description, height, weight," +
		" is_deleted, is_blocked, is_premium, is_show_distance, is_invisible, created_at, updated_at, last_online" +
		" FROM profiles" +
		" WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindProfileById", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.IsDeleted, &p.IsBlocked, &p.IsPremium,
		&p.IsShowDistance, &p.IsInvisible, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline)
	if err != nil {
		errorMessage := r.getErrorMessage("FindProfileById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) FindProfileBySessionId(
	ctx context.Context, sessionId string) (*entity.ProfileEntity, error) {
	p := &entity.ProfileEntity{}
	query := "SELECT id, session_id, display_name, birthday, gender, location, description, height, weight," +
		" is_deleted, is_blocked, is_premium, is_show_distance, is_invisible, created_at, updated_at, last_online" +
		" FROM profiles" +
		" WHERE session_id=$1"
	row := r.db.QueryRowContext(ctx, query, sessionId)
	if row == nil {
		errorMessage := r.getErrorMessage("FindProfileBySessionId", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.IsDeleted, &p.IsBlocked, &p.IsPremium,
		&p.IsShowDistance, &p.IsInvisible, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline)
	if err != nil {
		errorMessage := r.getErrorMessage("FindProfileBySessionId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) UpdateLastOnline(
	ctx context.Context, p *request.ProfileUpdateLastOnlineRequestRepositoryDto) error {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateLastOnline", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	query := "UPDATE profiles SET last_online=$1 WHERE session_id=$2"
	_, err = r.db.ExecContext(ctx, query, &p.LastOnline, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateLastOnline", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	tx.Commit()
	defer tx.Rollback()
	return nil
}

func (r *ProfileRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}
