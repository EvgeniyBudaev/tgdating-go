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
	ctx context.Context, p *request.ProfileAddRequestRepositoryDto) (*response.ResponseDto, error) {
	birthday := p.Birthday.Format("2006-01-02")
	query := "INSERT INTO dating.profiles (telegram_user_id, display_name, birthday, gender, location, description," +
		" height, weight, created_at, updated_at, last_online)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.DisplayName, &birthday, &p.Gender, &p.Location,
		&p.Description, &p.Height, &p.Weight, &p.CreatedAt, &p.UpdatedAt, &p.LastOnline)
	if row == nil {
		errorMessage := r.getErrorMessage("Add", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	profileResponse := &response.ResponseDto{
		Success: true,
	}
	return profileResponse, nil
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

func (r *ProfileRepository) FindByTelegramUserId(
	ctx context.Context, telegramUserId string) (*entity.ProfileEntity, error) {
	p := &entity.ProfileEntity{}
	query := "SELECT p.telegram_user_id, p.display_name, p.birthday, p.gender, p.location, p.description," +
		" p.height, p.weight, p.created_at, p.updated_at, p.last_online" +
		" FROM dating.profiles p" +
		//" JOIN dating.profile_images pi ON p.telegram_user_id = pi.telegram_user_id" +
		" WHERE p.telegram_user_id = $1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	if row == nil {
		errorMessage := r.getErrorMessage("FindByTelegramUserId", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	err := row.Scan(&p.TelegramUserId, &p.DisplayName, &p.Birthday, &p.Gender, &p.Location,
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

func (r *ProfileRepository) GetDetail(ctx context.Context,
	telegramUserId, viewedTelegramUserId string) (*response.ProfileDetailResponseRepositoryDto, error) {
	var distance *float64
	s := &response.StatusResponseDto{}
	p := &response.ProfileDetailResponseRepositoryDto{}
	query := "WITH user_details AS (" +
		" SELECT" +
		" pn1.location AS user1_location," +
		" p2.telegram_user_id AS telegram_user_id," +
		" p2.display_name AS display_name," +
		" p2.birthday AS birthday," +
		" p2.location AS location," +
		" p2.description AS description," +
		" p2.height AS height," +
		" p2.weight AS weight," +
		" ps2.is_blocked AS is_blocked," +
		" ps2.is_frozen AS is_frozen," +
		" ps2.is_invisible AS is_invisible," +
		" ps2.is_premium AS is_premium," +
		" ps2.is_show_distance AS is_show_distance," +
		" (CASE" +
		" WHEN p2.last_online >= NOW() AT TIME ZONE 'UTC' - INTERVAL '5 minutes' THEN true ELSE false" +
		" END) AS is_online," +
		" pn2.location AS user2_location" +
		" FROM dating.profiles p1" +
		" JOIN dating.profiles p2 ON p2.telegram_user_id = $2" +
		" JOIN dating.profile_statuses ps2 ON ps2.telegram_user_id = p2.telegram_user_id" +
		" LEFT JOIN dating.profile_navigators pn1 ON pn1.telegram_user_id = p1.telegram_user_id" +
		" LEFT JOIN dating.profile_navigators pn2 ON pn2.telegram_user_id = p2.telegram_user_id" +
		" WHERE p1.telegram_user_id = $1" +
		" )" +
		" SELECT " +
		" telegram_user_id, display_name, birthday, location, description, height, weight," +
		" is_blocked, is_frozen, is_invisible, is_online, is_premium, is_show_distance," +
		" ST_DistanceSphere(user1_location, user2_location) AS distance" +
		" FROM user_details"
	row := r.db.QueryRowContext(ctx, query, telegramUserId, viewedTelegramUserId)
	if row == nil {
		errorMessage := r.getErrorMessage("GetDetail", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	err := row.Scan(&p.TelegramUserId, &p.DisplayName, &p.Birthday, &p.Location, &p.Description, &p.Height, &p.Weight,
		&s.IsBlocked, &s.IsFrozen, &s.IsInvisible, &s.IsOnline, &s.IsPremium, &s.IsShowDistance, &distance)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		errorMessage := r.getErrorMessage("GetDetail", "Scan")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	if err != nil {
		errorMessage := r.getErrorMessage("GetDetail", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	var n *response.NavigatorDistanceResponseRepositoryDto
	n = &response.NavigatorDistanceResponseRepositoryDto{
		Distance: distance,
	}
	p = &response.ProfileDetailResponseRepositoryDto{
		TelegramUserId: p.TelegramUserId,
		DisplayName:    p.DisplayName,
		Birthday:       p.Birthday,
		Location:       p.Location,
		Description:    p.Description,
		Height:         p.Height,
		Weight:         p.Weight,
		Navigator:      n,
		Status:         s,
	}
	return p, nil
}

func (r *ProfileRepository) GetShortInfo(
	ctx context.Context, telegramUserId string) (*response.ProfileShortInfoResponseDto, error) {
	p := &response.ProfileShortInfoResponseDto{}
	query := "SELECT p.telegram_user_id, ps.is_blocked, ps.is_frozen," +
		" pf.search_gender, pf.looking_for, pf.age_from, pf.age_to, pf.distance, pf.page, pf.size" +
		" FROM dating.profiles p" +
		" JOIN dating.profile_statuses ps ON p.telegram_user_id = ps.telegram_user_id" +
		" JOIN dating.profile_filters pf ON p.telegram_user_id = pf.telegram_user_id" +
		" WHERE p.telegram_user_id = $1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	if row == nil {
		errorMessage := r.getErrorMessage("GetShortInfo", "QueryRowContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	err := row.Scan(&p.TelegramUserId, &p.IsBlocked, &p.IsFrozen, &p.SearchGender, &p.LookingFor,
		&p.AgeFrom, &p.AgeTo, &p.Distance, &p.Page, &p.Size)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		errorMessage := r.getErrorMessage("GetShortInfo", "Scan")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowFound))
		return nil, ErrNotRowFound
	}
	if err != nil {
		errorMessage := r.getErrorMessage("GetShortInfo", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ProfileRepository) SelectList(ctx context.Context,
	pr *request.ProfileGetListRequestRepositoryDto) (*response.ProfileListResponseRepositoryDto, error) {
	telegramUserId := pr.TelegramUserId
	searchGender := pr.SearchGender
	ageFrom := pr.AgeFrom
	ageTo := pr.AgeTo
	page := pr.Page
	size := pr.Size
	offset := (page - 1) * size
	distance := pr.Distance * 1000 // in meters
	query := "WITH filtered_profiles AS (" +
		" SELECT p.id, p.telegram_user_id, p.birthday, p.gender, ps.is_blocked, ps.is_frozen," +
		" p.created_at, p.updated_at, p.last_online," +
		" EXTRACT(YEAR FROM AGE(NOW(), p.birthday)) AS age," +
		" COALESCE(" +
		" ST_Distance(" +
		" (SELECT location FROM dating.profile_navigators WHERE telegram_user_id = p.telegram_user_id)::geography," +
		" ST_SetSRID(ST_Force2D(ST_MakePoint(" +
		" (SELECT ST_X(location) FROM dating.profile_navigators WHERE telegram_user_id = $1)," +
		" (SELECT ST_Y(location) FROM dating.profile_navigators WHERE telegram_user_id = $1)" +
		" )), 4326)::geography)," +
		" NULL::numeric" +
		" ) AS distance," +
		" (SELECT url FROM dating.profile_images pi" +
		" JOIN dating.profile_image_statuses pis ON pi.id = pis.id" +
		" WHERE pi.telegram_user_id = p.telegram_user_id AND pis.is_blocked = false AND pis.is_private = false" +
		" ORDER BY pi.created_at DESC LIMIT 1) AS url," +
		" (CASE " +
		" WHEN p.last_online >= NOW() AT TIME ZONE 'UTC' - INTERVAL '5 minutes' THEN true ELSE false" +
		" END) AS is_online" +
		" FROM dating.profiles p" +
		" JOIN dating.profile_statuses ps ON p.telegram_user_id = ps.telegram_user_id" +
		" LEFT JOIN dating.profile_navigators pn ON p.telegram_user_id = pn.telegram_user_id" +
		" WHERE ps.is_frozen = false AND ps.is_blocked = false AND" +
		" (EXTRACT(YEAR FROM AGE(NOW(), p.birthday)) BETWEEN $3 AND $4) AND" +
		" ($2 = 'all' OR p.gender = $2) AND p.telegram_user_id <> $1 AND" +
		" NOT EXISTS (SELECT 1 FROM dating.profile_blocks" +
		" WHERE telegram_user_id = $1 AND blocked_telegram_user_id = p.telegram_user_id)" +
		" )" +
		" SELECT telegram_user_id, last_online, distance, url, is_online" +
		" FROM filtered_profiles" +
		" WHERE distance IS NULL OR (distance < $5 AND distance IS NOT NULL)" +
		" ORDER BY CASE WHEN distance IS NULL THEN 1 ELSE 0 END, distance ASC, last_online DESC" +
		" LIMIT $6 OFFSET $7"
	rows, err := r.db.QueryContext(ctx, query, telegramUserId, searchGender, ageFrom, ageTo, distance, size, offset)
	if err != nil {
		errorMessage := r.getErrorMessage("SelectListByTelegramUserId",
			"QueryContext")
		r.logger.Info(errorMessage, zap.Error(ErrNotRowsFound))
		return nil, err
	}
	defer rows.Close()
	profileContent := make([]*response.ProfileListItemResponseDto, 0)
	for rows.Next() {
		p := response.ProfileListItemResponseDto{}
		err := rows.Scan(&p.TelegramUserId, &p.LastOnline, &p.Distance, &p.Url, &p.IsOnline)
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
		" WHERE ps.is_frozen = false AND ps.is_blocked = false AND" +
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
