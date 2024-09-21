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

func (r *ProfileRepository) FindProfileById(
	ctx context.Context, id uint64) (*entity.ProfileEntity, error) {
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

func (r *ProfileRepository) AddImage(
	ctx context.Context, p *request.ProfileImageAddRequestRepositoryDto) (*entity.ProfileImageEntity, error) {
	query := "INSERT INTO profile_images (session_id, name, url, size, is_deleted, is_blocked, is_primary," +
		" is_private, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.Name, &p.Url, &p.Size, &p.IsDeleted, &p.IsBlocked,
		&p.IsPrimary, &p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
	if row == nil {
		errorMessage := r.getErrorMessage("AddImage", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("AddImage", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindImageById(ctx, id)
}

func (r *ProfileRepository) UpdateImage(
	ctx context.Context, p *request.ProfileImageUpdateRequestRepositoryDto) (*entity.ProfileImageEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateImage", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_images SET name=$1, url=$2, size=$3, is_deleted=$4, is_blocked=$5," +
		" is_primary=$6, is_private=$7, updated_at=$8 WHERE id=$9"
	_, err = r.db.ExecContext(ctx, query, &p.Name, &p.Url, &p.Size, &p.IsDeleted, &p.IsBlocked,
		&p.IsPrimary, &p.IsPrivate, &p.UpdatedAt, &p.Id)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateImage", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindImageById(ctx, p.Id)
}

func (r *ProfileRepository) DeleteImage(
	ctx context.Context, p *request.ProfileImageDeleteRequestRepositoryDto) (*entity.ProfileImageEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteImage", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_images SET is_deleted=$1, updated_at=$2 WHERE id=$3"
	_, err = r.db.ExecContext(ctx, query, &p.IsDeleted, &p.UpdatedAt, &p.Id)
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteImage", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindImageById(ctx, p.Id)
}

func (r *ProfileRepository) FindImageById(ctx context.Context, imageId uint64) (*entity.ProfileImageEntity, error) {
	p := &entity.ProfileImageEntity{}
	query := "SELECT id, session_id, name, url, size, is_deleted, is_blocked, is_primary," +
		" is_private, created_at, updated_at" +
		" FROM profile_images" +
		" WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, imageId)
	if row == nil {
		errorMessage := r.getErrorMessage("FindImageById", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.Name, &p.Url, &p.Size, &p.IsDeleted, &p.IsBlocked, &p.IsPrimary,
		&p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindImageById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) SelectImageListPublicBySessionId(
	ctx context.Context, sessionId string) ([]*entity.ProfileImageEntity, error) {
	query := "SELECT id, session_id, name, url, size, is_deleted, is_blocked, is_primary," +
		" is_private, created_at, updated_at" +
		" FROM profile_images" +
		" WHERE session_id=$1 AND is_deleted=false AND is_blocked=false AND is_private=false"
	rows, err := r.db.QueryContext(ctx, query, sessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectImageListPublicBySessionId",
			"QueryContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, err
	}
	defer rows.Close()
	list := make([]*entity.ProfileImageEntity, 0)
	for rows.Next() {
		p := entity.ProfileImageEntity{}
		err := rows.Scan(&p.Id, &p.SessionId, &p.Name, &p.Url, &p.Size, &p.IsDeleted, &p.IsBlocked, &p.IsPrimary,
			&p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectImageListPublicBySessionId",
				"Scan")
			r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
			continue
		}
		list = append(list, &p)
	}
	return list, nil
}

func (r *ProfileRepository) SelectImageListBySessionId(
	ctx context.Context, sessionId string) ([]*entity.ProfileImageEntity, error) {
	query := "SELECT id, session_id, name, url, size, is_deleted, is_blocked, is_primary," +
		" is_private, created_at, updated_at" +
		" FROM profile_images" +
		" WHERE session_id=$1 AND is_deleted=false"
	rows, err := r.db.QueryContext(ctx, query, sessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectImageListBySessionId",
			"QueryContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, err
	}
	defer rows.Close()
	list := make([]*entity.ProfileImageEntity, 0)
	for rows.Next() {
		p := entity.ProfileImageEntity{}
		err := rows.Scan(&p.Id, &p.SessionId, &p.Name, &p.Url, &p.Size, &p.IsDeleted, &p.IsBlocked, &p.IsPrimary,
			&p.IsPrivate, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectImageListBySessionId",
				"Scan")
			r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
			continue
		}
		list = append(list, &p)
	}
	return list, nil
}

func (r *ProfileRepository) AddNavigator(
	ctx context.Context, p *request.ProfileNavigatorAddRequestRepositoryDto) (*entity.ProfileNavigatorEntity, error) {
	query := "INSERT INTO profile_navigators (session_id, location, is_deleted, created_at, updated_at)" +
		" VALUES ($1, ST_SetSRID(ST_MakePoint($2, $3),  4326), $4, $5, $6) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.Location.Longitude, &p.Location.Latitude, &p.IsDeleted,
		&p.CreatedAt, &p.UpdatedAt)
	if row == nil {
		errorMessage := r.getErrorMessage("AddNavigator", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("AddNavigator", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindNavigatorById(ctx, id)
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
	_, err = r.db.ExecContext(ctx, query, &p.Longitude, &p.Latitude, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateNavigator", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindNavigatorBySessionId(ctx, p.SessionId)
}

func (r *ProfileRepository) DeleteNavigator(
	ctx context.Context, p *request.ProfileNavigatorDeleteRequestDto) (*entity.ProfileNavigatorEntity, error) {
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

func (r *ProfileRepository) FindNavigatorById(
	ctx context.Context, id uint64) (*entity.ProfileNavigatorEntity, error) {
	p := &entity.ProfileNavigatorEntity{}
	var longitude sql.NullFloat64
	var latitude sql.NullFloat64
	query := `SELECT id, session_id, ST_X(location) as longitude, ST_Y(location) as latitude, is_deleted, created_at,
                updated_at
			  FROM profile_navigators
			  WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindNavigatorById", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
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

func (r *ProfileRepository) FindNavigatorBySessionId(
	ctx context.Context, sessionId string) (*entity.ProfileNavigatorEntity, error) {
	p := &entity.ProfileNavigatorEntity{}
	var longitude sql.NullFloat64
	var latitude sql.NullFloat64
	query := `SELECT id, session_id, ST_X(location) as longitude, ST_Y(location) as latitude, is_deleted, created_at,
                updated_at
			  FROM profile_navigators
			  WHERE session_id = $1`
	row := r.db.QueryRowContext(ctx, query, sessionId)
	if row == nil {
		errorMessage := r.getErrorMessage("FindNavigatorBySessionId",
			"QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
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

func (r *ProfileRepository) AddFilter(
	ctx context.Context, p *request.ProfileFilterAddRequestRepositoryDto) (*entity.ProfileFilterEntity, error) {
	query := "INSERT INTO profile_filters (session_id, search_gender, looking_for, age_from, age_to, distance, page," +
		" size, is_deleted, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.SearchGender, &p.LookingFor, &p.AgeFrom, &p.AgeTo,
		&p.Distance, &p.Page, &p.Size, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if row == nil {
		errorMessage := r.getErrorMessage("AddFilter", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("AddFilter", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindFilterById(ctx, id)
}

func (r *ProfileRepository) UpdateFilter(
	ctx context.Context, p *request.ProfileFilterUpdateRequestRepositoryDto) (*entity.ProfileFilterEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateFilter", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_filters SET search_gender=$1, looking_for=$2, age_from=$3, age_to=$4, distance=$5," +
		" page=$6, size=$7, updated_at=$8 WHERE session_id=$9"
	_, err = r.db.ExecContext(ctx, query, &p.SearchGender, &p.LookingFor, &p.AgeFrom, &p.AgeTo,
		&p.Distance, &p.Page, &p.Size, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateFilter", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindFilterBySessionId(ctx, p.SessionId)
}

func (r *ProfileRepository) DeleteFilter(
	ctx context.Context, p *request.ProfileFilterDeleteRequestRepositoryDto) (*entity.ProfileFilterEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteFilter", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_filters SET is_deleted=$1, updated_at=$2 WHERE session_id=$3"
	_, err = r.db.ExecContext(ctx, query, &p.IsDeleted, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteFilter", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindFilterBySessionId(ctx, p.SessionId)
}

func (r *ProfileRepository) FindFilterById(
	ctx context.Context, id uint64) (*entity.ProfileFilterEntity, error) {
	p := &entity.ProfileFilterEntity{}
	query := "SELECT id, session_id, search_gender, looking_for, age_from, age_to, distance, page, size, is_deleted," +
		" created_at, updated_at" +
		" FROM profile_filters" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindFilterById", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.SearchGender, &p.LookingFor, &p.AgeFrom, &p.AgeTo, &p.Distance, &p.Page,
		&p.Size, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindFilterById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) FindFilterBySessionId(
	ctx context.Context, sessionId string) (*entity.ProfileFilterEntity, error) {
	p := &entity.ProfileFilterEntity{}
	query := "SELECT id, session_id, search_gender, looking_for, age_from, age_to, distance, page, size, is_deleted," +
		" created_at, updated_at" +
		" FROM profile_filters" +
		" WHERE session_id = $1"
	row := r.db.QueryRowContext(ctx, query, sessionId)
	if row == nil {
		errorMessage := r.getErrorMessage("FindFilterBySessionId", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.SearchGender, &p.LookingFor, &p.AgeFrom, &p.AgeTo, &p.Distance, &p.Page,
		&p.Size, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindFilterBySessionId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) AddTelegram(
	ctx context.Context, p *request.ProfileTelegramAddRequestRepositoryDto) (*entity.ProfileTelegramEntity, error) {
	query := "INSERT INTO profile_telegrams (session_id, user_id, username, first_name, last_name, language_code," +
		" allows_write_to_pm, query_id, chat_id, is_deleted, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.UserId, &p.UserName, &p.FirstName, &p.LastName,
		&p.LanguageCode, &p.AllowsWriteToPm, &p.QueryId, &p.ChatId, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if row == nil {
		errorMessage := r.getErrorMessage("AddTelegram", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("AddTelegram", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindTelegramById(ctx, id)
}

func (r *ProfileRepository) UpdateTelegram(
	ctx context.Context, p *request.ProfileTelegramUpdateRequestRepositoryDto) (*entity.ProfileTelegramEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateTelegram", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_telegrams SET user_id=$1, username=$2, first_name=$3, last_name=$4, language_code=$5," +
		" allows_write_to_pm=$6, query_id=$7, chat_id=$8, updated_at=$9" +
		" WHERE session_id=$10"
	_, err = r.db.ExecContext(ctx, query, &p.UserId, &p.UserName, &p.FirstName, &p.LastName, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryId, &p.ChatId, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateTelegram", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindTelegramBySessionId(ctx, p.SessionId)
}

func (r *ProfileRepository) DeleteTelegram(
	ctx context.Context, p *request.ProfileTelegramDeleteRequestRepositoryDto) (*entity.ProfileTelegramEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteTelegram", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profile_telegrams SET is_deleted=$1, updated_at=$2 WHERE session_id=$3"
	_, err = r.db.ExecContext(ctx, query, &p.IsDeleted, &p.UpdatedAt, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteTelegram", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindTelegramBySessionId(ctx, p.SessionId)
}

func (r *ProfileRepository) FindTelegramById(
	ctx context.Context, id uint64) (*entity.ProfileTelegramEntity, error) {
	p := &entity.ProfileTelegramEntity{}
	query := "SELECT id, session_id, user_id, username, first_name, last_name, language_code, allows_write_to_pm," +
		" query_id, chat_id, is_deleted, created_at, updated_at" +
		" FROM profile_telegrams" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindTelegramById", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.UserId, &p.UserName, &p.FirstName, &p.LastName, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryId, &p.ChatId, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindTelegramById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) FindTelegramBySessionId(
	ctx context.Context, sessionID string) (*entity.ProfileTelegramEntity, error) {
	p := &entity.ProfileTelegramEntity{}
	query := "SELECT id, session_id, user_id, username, first_name, last_name, language_code, allows_write_to_pm," +
		" query_id, chat_id, is_deleted, created_at, updated_at" +
		" FROM profile_telegrams" +
		" WHERE session_id = $1"
	row := r.db.QueryRowContext(ctx, query, sessionID)
	if row == nil {
		errorMessage := r.getErrorMessage("FindTelegramBySessionId",
			"QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.UserId, &p.UserName, &p.FirstName, &p.LastName, &p.LanguageCode,
		&p.AllowsWriteToPm, &p.QueryId, &p.ChatId, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindTelegramBySessionId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) AddBlock(
	ctx context.Context, p *request.ProfileBlockAddRequestRepositoryDto) (*entity.ProfileBlockEntity, error) {
	query := "INSERT INTO profile_blocks (session_id, blocked_user_session_id, is_blocked, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.BlockedUserSessionId, &p.IsBlocked, &p.CreatedAt,
		&p.UpdatedAt)
	if row == nil {
		errorMessage := r.getErrorMessage("AddBlock", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("AddBlock", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindBlockById(ctx, id)
}

func (r *ProfileRepository) FindBlockById(ctx context.Context, id uint64) (*entity.ProfileBlockEntity, error) {
	p := &entity.ProfileBlockEntity{}
	query := "SELECT id, session_id, blocked_user_session_id, is_blocked, created_at, updated_at " +
		" FROM profile_blocks" +
		" WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindBlockById", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.BlockedUserSessionId, &p.IsBlocked, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindBlockById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) AddLike(
	ctx context.Context, p *request.ProfileLikeAddRequestRepositoryDto) (*entity.ProfileLikeEntity, error) {
	query := "INSERT INTO profile_likes (session_id, liked_session_id, is_liked, is_deleted, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.LikedSessionId, &p.IsLiked, &p.IsDeleted, &p.CreatedAt,
		&p.UpdatedAt)
	if row == nil {
		errorMessage := r.getErrorMessage("AddLike", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("AddLike", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindLikeById(ctx, id)
}

func (r *ProfileRepository) FindLikeById(ctx context.Context, id uint64) (*entity.ProfileLikeEntity, error) {
	p := &entity.ProfileLikeEntity{}
	query := "SELECT id, session_id, liked_session_id, is_liked, is_deleted, created_at, updated_at " +
		" FROM profile_likes" +
		" WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindLikeById", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.LikedSessionId, &p.IsLiked, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindLikeById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) AddComplaint(
	ctx context.Context, p *request.ProfileComplaintAddRequestRepositoryDto) (*entity.ProfileComplaintEntity, error) {
	query := "INSERT INTO profile_complaints (session_id, criminal_session_id, reason, is_deleted, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.CriminalSessionId, &p.Reason, &p.IsDeleted, &p.IsDeleted,
		&p.CreatedAt, &p.UpdatedAt)
	if row == nil {
		errorMessage := r.getErrorMessage("AddComplaint", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("AddComplaint", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindComplaintById(ctx, id)
}

func (r *ProfileRepository) FindComplaintById(ctx context.Context, id uint64) (*entity.ProfileComplaintEntity, error) {
	p := &entity.ProfileComplaintEntity{}
	query := "SELECT id, session_id, criminal_session_id, reason, is_deleted, created_at, updated_at " +
		" FROM profile_complaints" +
		" WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindComplaintById", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, ErrNotRowsFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.CriminalSessionId, &p.Reason, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindComplaintById", "Scan")
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
