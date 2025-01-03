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
	errorFilePathILike = "internal/repository/psql/like-repository.go"
)

type LikeRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewLikeRepository(l logger.Logger, db *sql.DB) *LikeRepository {
	return &LikeRepository{
		logger: l,
		db:     db,
	}
}

func (r *LikeRepository) Add(
	ctx context.Context, p *request.LikeAddRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "INSERT INTO dating.profile_likes (telegram_user_id, liked_telegram_user_id, is_liked, created_at," +
		" updated_at)" +
		" VALUES ($1, $2, $3, $4, $5)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.LikedTelegramUserId, &p.IsLiked, &p.CreatedAt,
		&p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	likeResponse := &response.ResponseDto{
		Success: true,
	}
	return likeResponse, nil
}

func (r *LikeRepository) Update(
	ctx context.Context, p *request.LikeUpdateRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "UPDATE dating.profile_likes SET is_liked = $1, updated_at = $2" +
		" WHERE id = $3"
	_, err := r.db.ExecContext(ctx, query, &p.IsLiked, &p.UpdatedAt, &p.Id)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	likeResponse := &response.ResponseDto{
		Success: true,
	}
	return likeResponse, nil
}

func (r *LikeRepository) DeleteRelatedProfiles(
	ctx context.Context, id string) (*response.ResponseDto, error) {
	query := "DELETE FROM dating.profile_likes WHERE liked_telegram_user_id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteRelatedProfiles", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	likeResponse := &response.ResponseDto{
		Success: true,
	}
	return likeResponse, nil
}

func (r *LikeRepository) FindById(ctx context.Context, id uint64) (*entity.LikeEntity, error) {
	p := &entity.LikeEntity{}
	query := "SELECT id, telegram_user_id, liked_telegram_user_id, is_liked, created_at, updated_at " +
		" FROM dating.profile_likes" +
		" WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.LikedTelegramUserId, &p.IsLiked, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, nil
	}
	return p, nil
}

func (r *LikeRepository) FindLastLike(ctx context.Context, telegramUserId string) (*entity.LikeEntity, error) {
	p := &entity.LikeEntity{}
	query := "SELECT id, telegram_user_id, liked_telegram_user_id, is_liked, created_at, updated_at " +
		" FROM dating.profile_likes" +
		" WHERE liked_telegram_user_id = $1 ORDER BY created_at DESC LIMIT 1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.LikedTelegramUserId, &p.IsLiked, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		errorMessage := r.getErrorMessage("FindLastLike", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, nil
	}
	return p, nil
}

func (r *LikeRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathILike)
}
