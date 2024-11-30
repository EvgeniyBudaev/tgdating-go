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
	errorFilePath          = "internal/repository/psql/profile-repository.go"
	ErrNotRowsFoundMessage = "profiles not found"
	ErrNotRowFoundMessage  = "profile not found"
)

var (
	ErrNotRowsFound = errors.New(ErrNotRowsFoundMessage)
	ErrNotRowFound  = errors.New(ErrNotRowFoundMessage)
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

func (r *ProfileRepository) Add(
	ctx context.Context, p *request.ProfileAddRequestRepositoryDto) (*entity.ProfileEntity, error) {
	birthday := p.Birthday.Format("2006-01-02")
	query := "INSERT INTO dating.profiles (session_id, display_name, birthday, gender, location, description," +
		" height, weight, is_frozen, is_blocked, is_premium, is_show_distance, is_invisible," +
		" created_at, updated_at, last_online)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.DisplayName, &birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, p.IsFrozen, &p.IsBlocked, &p.IsPremium, &p.IsShowDistance,
		&p.IsInvisible, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline)
	if row == nil {
		errorMessage := r.getErrorMessage("Add", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindById(ctx, id)
}

func (r *ProfileRepository) Update(
	ctx context.Context, p *request.ProfileUpdateRequestRepositoryDto) (*entity.ProfileEntity, error) {
	query := "UPDATE dating.profiles SET display_name=$1, birthday=$2, gender=$3, location=$4," +
		" description=$5, height=$6, weight=$7, updated_at=$8, last_online=$9" +
		" WHERE session_id=$10"
	_, err := r.db.ExecContext(ctx, query, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.UpdatedAt, &p.LastOnline, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindBySessionId(ctx, p.SessionId)
}

func (r *ProfileRepository) Freeze(
	ctx context.Context, p *request.ProfileFreezeRequestRepositoryDto) (*entity.ProfileEntity, error) {
	query := "UPDATE dating.profiles SET is_frozen=$1, updated_at=$2, last_online=$3 WHERE session_id=$4"
	_, err := r.db.ExecContext(ctx, query, &p.IsFrozen, &p.UpdatedAt, &p.LastOnline, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("Freeze", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindBySessionId(ctx, p.SessionId)
}

func (r *ProfileRepository) Restore(
	ctx context.Context, p *request.ProfileRestoreRequestRepositoryDto) (*entity.ProfileEntity, error) {
	query := "UPDATE dating.profiles SET is_frozen=$1, updated_at=$2, last_online=$3 WHERE session_id=$4"
	_, err := r.db.ExecContext(ctx, query, &p.IsFrozen, &p.UpdatedAt, &p.LastOnline, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("Restore", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindBySessionId(ctx, p.SessionId)
}

func (r *ProfileRepository) Delete(
	ctx context.Context, p *request.ProfileDeleteRequestDto) (*response.ResponseDto, error) {
	query := "DELETE FROM dating.profiles WHERE session_id=$1"
	_, err := r.db.ExecContext(ctx, query, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("Delete", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	profileResponse := &response.ResponseDto{
		Success: true,
	}
	return profileResponse, nil
}

func (r *ProfileRepository) FindById(ctx context.Context, id uint64) (*entity.ProfileEntity, error) {
	p := &entity.ProfileEntity{}
	query := "SELECT id, session_id, display_name, birthday, gender, location, description, height, weight," +
		" is_frozen, is_blocked, is_premium, is_show_distance, is_invisible, created_at, updated_at, last_online" +
		" FROM dating.profiles" +
		" WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindById", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.IsFrozen, &p.IsBlocked, &p.IsPremium,
		&p.IsShowDistance, &p.IsInvisible, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	if err != nil {
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) FindBySessionId(
	ctx context.Context, sessionId string) (*entity.ProfileEntity, error) {
	p := &entity.ProfileEntity{}
	query := "SELECT id, session_id, display_name, birthday, gender, location, description, height, weight," +
		" is_frozen, is_blocked, is_premium, is_show_distance, is_invisible, created_at, updated_at, last_online" +
		" FROM dating.profiles" +
		" WHERE session_id=$1"
	row := r.db.QueryRowContext(ctx, query, sessionId)
	if row == nil {
		errorMessage := r.getErrorMessage("FindBySessionId", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.IsFrozen, &p.IsBlocked, &p.IsPremium,
		&p.IsShowDistance, &p.IsInvisible, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		errorMessage := r.getErrorMessage("FindBySessionId", "Scan")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	if err != nil {
		errorMessage := r.getErrorMessage("FindBySessionId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) SelectListBySessionId(ctx context.Context,
	pr *request.ProfileGetListRequestRepositoryDto) (*response.ProfileListResponseRepositoryDto, error) {
	sessionId := pr.SessionId
	searchGender := pr.SearchGender
	ageFrom := pr.AgeFrom
	ageTo := pr.AgeTo
	page := pr.Page
	size := pr.Size
	offset := (page - 1) * size
	distance := pr.Distance * 1000
	query := "WITH filtered_profiles AS (" +
		" SELECT p.id, p.session_id, p.display_name, p.birthday, p.gender, p.location, p.description, p.height," +
		" p.weight, p.is_frozen, p.is_blocked, p.is_premium, p.is_show_distance, p.is_invisible, p.created_at," +
		" p.updated_at, p.last_online," +
		" EXTRACT(YEAR FROM AGE(NOW(), p.birthday)) AS age," +
		" COALESCE(" +
		" ST_Distance(" +
		" (SELECT location FROM dating.profile_navigators WHERE session_id = p.session_id)::geography," +
		" ST_SetSRID(ST_Force2D(ST_MakePoint(" +
		"(SELECT ST_X(location) FROM dating.profile_navigators WHERE session_id = $1)," +
		" (SELECT ST_Y(location) FROM dating.profile_navigators WHERE session_id = $1)" +
		" )), 4326)::geography)," +
		" NULL::numeric" +
		" ) AS distance" +
		" FROM dating.profiles p" +
		" LEFT JOIN dating.profile_navigators pn ON p.session_id = pn.session_id" +
		" WHERE p.is_frozen = false AND p.is_blocked = false AND" +
		" (EXTRACT(YEAR FROM AGE(NOW(), p.birthday)) BETWEEN $3 AND $4) AND" +
		" ($2 = 'all' OR gender = $2) AND p.session_id <> $1 AND" +
		" NOT EXISTS (SELECT 1 FROM dating.profile_blocks" +
		" WHERE session_id = $1 AND blocked_user_session_id = p.session_id)" +
		" )" +
		" SELECT *" +
		" FROM filtered_profiles" +
		" WHERE distance IS NULL OR (distance < $5 AND distance IS NOT NULL)" +
		" ORDER BY CASE WHEN distance IS NULL THEN 1 ELSE 0 END, distance ASC, last_online DESC" +
		" LIMIT $6 OFFSET $7"
	rows, err := r.db.QueryContext(ctx, query, sessionId, searchGender, ageFrom, ageTo, distance, size, offset)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectListBySessionId",
			"QueryContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, err
	}
	defer rows.Close()
	profileContent := make([]*response.ProfileListItemResponseRepositoryDto, 0)
	for rows.Next() {
		p := response.ProfileListItemResponseRepositoryDto{}
		err := rows.Scan(&p.Id, &p.SessionId, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
			&p.Description, &p.Height, &p.Weight, &p.IsFrozen, &p.IsBlocked, &p.IsPremium,
			&p.IsShowDistance, &p.IsInvisible, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline, &p.Age, &p.Distance)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectListBySessionId", "Scan")
			r.logger.Info(errorMessage, zap.Error(ErrNotRowsFound))
			continue
		}
		profileContent = append(profileContent, &p)
	}
	totalEntities, err := r.getTotalEntities(ctx, sessionId, searchGender, ageFrom, ageTo)
	if err != nil {
		return nil, err
	}
	paginationEntity := entity.GetPagination(page, size, totalEntities)
	paginationProfileEntityList := &response.ProfileListResponseRepositoryDto{
		PaginationEntity: paginationEntity,
		Content:          profileContent,
	}
	return paginationProfileEntityList, nil
}

func (r *ProfileRepository) getTotalEntities(
	ctx context.Context, sessionId, searchGender string, ageFrom, ageTo uint64) (uint64, error) {
	query := "SELECT COUNT(*)" +
		" FROM dating.profiles" +
		" WHERE is_frozen=false AND is_blocked=false AND" +
		" (EXTRACT(YEAR FROM AGE(NOW(), birthday)) BETWEEN $3 AND $4) AND" +
		" ($2 = 'all' OR gender = $2) AND session_id <> $1"
	var totalEntities uint64
	err := r.db.QueryRowContext(ctx, query, sessionId, searchGender, ageFrom, ageTo).Scan(&totalEntities)
	if err != nil {
		errorMessage := r.getErrorMessage("getTotalEntities", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return 0, err
	}
	return totalEntities, nil
}

func (r *ProfileRepository) UpdateLastOnline(
	ctx context.Context, p *request.ProfileUpdateLastOnlineRequestRepositoryDto) error {
	query := "UPDATE dating.profiles SET last_online=$1 WHERE session_id=$2"
	_, err := r.db.ExecContext(ctx, query, &p.LastOnline, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("UpdateLastOnline", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return err
	}
	return nil
}

func (r *ProfileRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePath)
}
