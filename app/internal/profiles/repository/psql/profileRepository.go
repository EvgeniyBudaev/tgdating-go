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
	errorFilePath = "internal/repository/psql/profileRepository.go"
)

var (
	ErrNotRowsFound = errors.New("profiles not found")
	ErrNotRowFound  = errors.New("profile not found")
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
	query := "INSERT INTO profiles (session_id, display_name, birthday, gender, location, description," +
		" height, weight, is_deleted, is_blocked, is_premium, is_show_distance, is_invisible," +
		" created_at, updated_at, last_online)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.DisplayName, &birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, p.IsDeleted, &p.IsBlocked, &p.IsPremium, &p.IsShowDistance,
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
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "Begin")
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
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindBySessionId(ctx, p.SessionId)
}

func (r *ProfileRepository) Delete(
	ctx context.Context, p *request.ProfileDeleteRequestRepositoryDto) (*entity.ProfileEntity, error) {
	tx, err := r.db.Begin()
	if err != nil {
		errorMessage := r.getErrorMessage("Delete", "Begin")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer tx.Rollback()
	query := "UPDATE profiles SET is_deleted=$1, updated_at=$2, last_online=$3 WHERE session_id=$4"
	_, err = r.db.ExecContext(ctx, query, &p.IsDeleted, &p.UpdatedAt, &p.LastOnline, &p.SessionId)
	if err != nil {
		errorMessage := r.getErrorMessage("Delete", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	tx.Commit()
	return r.FindBySessionId(ctx, p.SessionId)
}

func (r *ProfileRepository) FindById(ctx context.Context, id uint64) (*entity.ProfileEntity, error) {
	p := &entity.ProfileEntity{}
	query := "SELECT id, session_id, display_name, birthday, gender, location, description, height, weight," +
		" is_deleted, is_blocked, is_premium, is_show_distance, is_invisible, created_at, updated_at, last_online" +
		" FROM profiles" +
		" WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindById", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.IsDeleted, &p.IsBlocked, &p.IsPremium,
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
		" is_deleted, is_blocked, is_premium, is_show_distance, is_invisible, created_at, updated_at, last_online" +
		" FROM profiles" +
		" WHERE session_id=$1"
	row := r.db.QueryRowContext(ctx, query, sessionId)
	if row == nil {
		errorMessage := r.getErrorMessage("FindBySessionId", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.IsDeleted, &p.IsBlocked, &p.IsPremium,
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
	query := "SELECT p.id, p.session_id, p.display_name, p.birthday, p.gender, p.location, p.description, p.height," +
		" p.weight, p.is_deleted, p.is_blocked, p.is_premium, p.is_show_distance, p.is_invisible, p.created_at," +
		" p.updated_at, p.last_online," +
		" EXTRACT(YEAR FROM AGE(NOW(), p.birthday)) AS age," +
		" ST_Distance(" +
		" (SELECT location FROM profile_navigators WHERE session_id = p.session_id)::geography," +
		" ST_SetSRID(ST_Force2D(ST_MakePoint(" +
		" (SELECT ST_X(location) FROM profile_navigators WHERE session_id = $1)," +
		" (SELECT ST_Y(location) FROM profile_navigators WHERE session_id = $1)" +
		" )), 4326)::geography) AS distance" +
		" FROM profiles p" +
		" JOIN profile_navigators pn ON p.session_id = pn.session_id" +
		" WHERE p.is_deleted = false AND p.is_blocked = false AND" +
		" (EXTRACT(YEAR FROM AGE(NOW(), p.birthday)) BETWEEN $3 AND $4) AND" +
		" ($2 = 'all' OR gender = $2) AND p.session_id <> $1 AND" +
		" NOT EXISTS (SELECT 1 FROM profile_blocks WHERE session_id = $1 AND" +
		" blocked_user_session_id = p.session_id) AND" +
		" ST_Distance((SELECT location FROM profile_navigators WHERE session_id = p.session_id)::geography," +
		" ST_SetSRID(ST_MakePoint((SELECT ST_X(location) FROM profile_navigators WHERE session_id = $1)," +
		" (SELECT ST_Y(location) FROM profile_navigators" +
		" WHERE session_id = $1)), 4326)::geography) <= $5" +
		" ORDER BY distance ASC, p.last_online DESC" +
		" LIMIT $6 OFFSET $7"
	rows, err := r.db.QueryContext(ctx, query, sessionId, searchGender, ageFrom, ageTo, distance, size, offset)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectListBySessionId",
			"QueryContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, err
	}
	defer rows.Close()
	profileEntityList := make([]*response.ProfileListItemResponseRepositoryDto, 0)
	for rows.Next() {
		p := response.ProfileListItemResponseRepositoryDto{}
		err := rows.Scan(&p.Id, &p.SessionId, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
			&p.Description, &p.Height, &p.Weight, &p.IsDeleted, &p.IsBlocked, &p.IsPremium,
			&p.IsShowDistance, &p.IsInvisible, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline, &p.Age, &p.Distance)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectListBySessionId", "Scan")
			r.logger.Info(errorMessage, zap.Error(ErrNotRowsFound))
			continue
		}
		profileEntityList = append(profileEntityList, &p)
	}
	numberEntities, err := r.getNumberEntities(ctx, sessionId, searchGender, ageFrom, ageTo)
	if err != nil {
		return nil, err
	}
	paginationEntity := entity.GetPagination(page, size, numberEntities)
	paginationProfileEntityList := &response.ProfileListResponseRepositoryDto{
		PaginationEntity: paginationEntity,
		Content:          profileEntityList,
	}
	return paginationProfileEntityList, nil
}

func (r *ProfileRepository) getNumberEntities(
	ctx context.Context, sessionId, searchGender string, ageFrom, ageTo uint64) (uint64, error) {
	query := "SELECT COUNT(*)" +
		" FROM profiles" +
		" WHERE is_deleted=false AND is_blocked=false AND" +
		" (EXTRACT(YEAR FROM AGE(NOW(), birthday)) BETWEEN $3 AND $4) AND" +
		" ($2 = 'all' OR gender = $2) AND session_id <> $1"
	var numberEntities uint64
	err := r.db.QueryRowContext(ctx, query, sessionId, searchGender, ageFrom, ageTo).Scan(&numberEntities)
	if err != nil {
		errorMessage := r.getErrorMessage("getNumberEntities", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return 0, err
	}
	return numberEntities, nil
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
