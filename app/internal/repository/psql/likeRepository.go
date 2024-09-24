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
	errorFilePathILike = "internal/repository/psql/likeRepository.go"
)

var (
	ErrNotRowsFoundLike = errors.New("no rows found")
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

func (r *LikeRepository) AddLike(
	ctx context.Context, p *request.LikeAddRequestRepositoryDto) (*entity.LikeEntity, error) {
	query := "INSERT INTO profile_likes (session_id, liked_session_id, is_liked, is_deleted, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.LikedSessionId, &p.IsLiked, &p.IsDeleted, &p.CreatedAt,
		&p.UpdatedAt)
	if row == nil {
		errorMessage := r.getErrorMessage("AddLike", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFoundLike))
		return nil, ErrNotRowsFoundLike
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

func (r *LikeRepository) FindLikeById(ctx context.Context, id uint64) (*entity.LikeEntity, error) {
	p := &entity.LikeEntity{}
	query := "SELECT id, session_id, liked_session_id, is_liked, is_deleted, created_at, updated_at " +
		" FROM profile_likes" +
		" WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindLikeById", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFoundLike))
		return nil, ErrNotRowsFoundLike
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.LikedSessionId, &p.IsLiked, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindLikeById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *LikeRepository) FindLikeBySessionId(ctx context.Context, sessionId string) (*entity.LikeEntity, error) {
	p := &entity.LikeEntity{}
	query := "SELECT id, session_id, liked_session_id, is_liked, is_deleted, created_at, updated_at " +
		" FROM profile_likes" +
		" WHERE session_id=$1"
	row := r.db.QueryRowContext(ctx, query, sessionId)
	err := row.Scan(&p.Id, &p.SessionId, &p.LikedSessionId, &p.IsLiked, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		errorMessage := r.getErrorMessage("FindLikeBySessionId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, nil
	}
	return p, nil
}

func (r *LikeRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathILike)
}
