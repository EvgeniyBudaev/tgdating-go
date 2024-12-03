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
	query := "INSERT INTO dating.profiles (telegram_user_id, display_name, birthday, gender, location, description," +
		" height, weight, created_at, updated_at, last_online)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.DisplayName, &birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline)
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
	query := "UPDATE dating.profiles SET display_name = $1, birthday = $2, gender = $3, location = $4," +
		" description = $5, height = $6, weight = $7, updated_at = $8, last_online = $9" +
		" WHERE telegram_user_id = $10"
	_, err := r.db.ExecContext(ctx, query, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.UpdatedAt, &p.LastOnline, &p.TelegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindByTelegramUserId(ctx, p.TelegramUserId)
}

func (r *ProfileRepository) Delete(
	ctx context.Context, p *request.ProfileDeleteRequestDto) (*response.ResponseDto, error) {
	query := "DELETE FROM dating.profiles WHERE telegram_user_id = $1"
	_, err := r.db.ExecContext(ctx, query, &p.TelegramUserId)
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
	query := "SELECT id, telegram_user_id, display_name, birthday, gender, location, description, height, weight," +
		" created_at, updated_at, last_online" +
		" FROM dating.profiles" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindById", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline)
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

func (r *ProfileRepository) FindByTelegramUserId(
	ctx context.Context, telegramUserId string) (*entity.ProfileEntity, error) {
	p := &entity.ProfileEntity{}
	query := "SELECT id, telegram_user_id, display_name, birthday, gender, location, description, height, weight," +
		" created_at, updated_at, last_online" +
		" FROM dating.profiles" +
		" WHERE telegram_user_id = $1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	if row == nil {
		errorMessage := r.getErrorMessage("FindByTelegramUserId", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		errorMessage := r.getErrorMessage("FindByTelegramUserId", "Scan")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	if err != nil {
		errorMessage := r.getErrorMessage("FindByTelegramUserId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) SelectListByTelegramUserId(ctx context.Context,
	pr *request.ProfileGetListRequestRepositoryDto) (*response.ProfileListResponseRepositoryDto, error) {
	telegramUserId := pr.TelegramUserId
	searchGender := pr.SearchGender
	ageFrom := pr.AgeFrom
	ageTo := pr.AgeTo
	page := pr.Page
	size := pr.Size
	offset := (page - 1) * size
	distance := pr.Distance * 1000
	query := "WITH filtered_profiles AS (" +
		" SELECT p.id, p.telegram_user_id, p.display_name, p.birthday, p.gender, p.location, p.description, p.height," +
		" p.weight, ps.is_frozen, ps.is_blocked, ps.is_premium, ps.is_show_distance, ps.is_invisible, p.created_at," +
		" p.updated_at, p.last_online," +
		" EXTRACT(YEAR FROM AGE(NOW(), p.birthday)) AS age," +
		" COALESCE(" +
		" ST_Distance(" +
		" (SELECT location FROM dating.profile_navigators WHERE telegram_user_id = p.telegram_user_id)::geography," +
		" ST_SetSRID(ST_Force2D(ST_MakePoint(" +
		"(SELECT ST_X(location) FROM dating.profile_navigators WHERE telegram_user_id = $1)," +
		" (SELECT ST_Y(location) FROM dating.profile_navigators WHERE telegram_user_id = $1)" +
		" )), 4326)::geography)," +
		" NULL::numeric" +
		" ) AS distance" +
		" FROM dating.profiles p" +
		" JOIN dating.profile_statuses ps ON p.telegram_user_id = ps.telegram_user_id" +
		" LEFT JOIN dating.profile_navigators pn ON p.telegram_user_id = pn.telegram_user_id" +
		" WHERE ps.is_frozen = false AND ps.is_blocked = false AND" +
		" (EXTRACT(YEAR FROM AGE(NOW(), p.birthday)) BETWEEN $3 AND $4) AND" +
		" ($2 = 'all' OR p.gender = $2) AND p.telegram_user_id <> $1 AND" +
		" NOT EXISTS (SELECT 1 FROM dating.profile_blocks" +
		" WHERE telegram_user_id = $1 AND blocked_telegram_user_id = p.telegram_user_id)" +
		" )" +
		" SELECT *" +
		" FROM filtered_profiles" +
		" WHERE distance IS NULL OR (distance < $5 AND distance IS NOT NULL)" +
		" ORDER BY CASE WHEN distance IS NULL THEN 1 ELSE 0 END, distance ASC, p.last_online DESC" +
		" LIMIT $6 OFFSET $7"
	rows, err := r.db.QueryContext(ctx, query, telegramUserId, searchGender, ageFrom, ageTo, distance, size, offset)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectListByTelegramUserId",
			"QueryContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, err
	}
	defer rows.Close()
	profileContent := make([]*response.ProfileListItemResponseRepositoryDto, 0)
	for rows.Next() {
		p := response.ProfileListItemResponseRepositoryDto{}
		err := rows.Scan(&p.Id, &p.TelegramUserId, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
			&p.Description, &p.Height, &p.Weight, &p.IsFrozen, &p.IsBlocked, &p.IsPremium,
			&p.IsShowDistance, &p.IsInvisible, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline, &p.Age, &p.Distance)
		if err != nil {
			errorMessage := r.getErrorMessage("SelectListByTelegramUserId", "Scan")
			r.logger.Info(errorMessage, zap.Error(ErrNotRowsFound))
			continue
		}
		profileContent = append(profileContent, &p)
	}
	totalEntities, err := r.getTotalEntities(ctx, telegramUserId, searchGender, ageFrom, ageTo)
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
	ctx context.Context, telegramUserId, searchGender string, ageFrom, ageTo uint64) (uint64, error) {
	query := "SELECT COUNT(*)" +
		" FROM dating.profiles p" +
		" JOIN dating.profile_statuses ps ON p.telegram_user_id = ps.telegram_user_id" +
		" WHERE ps.is_frozen=false AND ps.is_blocked=false AND" +
		" (EXTRACT(YEAR FROM AGE(NOW(), p.birthday)) BETWEEN $3 AND $4) AND" +
		" ($2 = 'all' OR p.gender = $2) AND p.telegram_user_id <> $1"
	var totalEntities uint64
	err := r.db.QueryRowContext(ctx, query, telegramUserId, searchGender, ageFrom, ageTo).Scan(&totalEntities)
	if err != nil {
		errorMessage := r.getErrorMessage("getTotalEntities", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return 0, err
	}
	return totalEntities, nil
}

func (r *ProfileRepository) UpdateLastOnline(
	ctx context.Context, p *request.ProfileUpdateLastOnlineRequestRepositoryDto) error {
	query := "UPDATE dating.profiles SET last_online = $1 WHERE telegram_user_id = $2"
	_, err := r.db.ExecContext(ctx, query, &p.LastOnline, &p.TelegramUserId)
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
