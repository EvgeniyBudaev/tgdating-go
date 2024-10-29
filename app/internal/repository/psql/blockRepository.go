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
	errorFilePathBlock = "internal/repository/psql/blockRepository.go"
)

type BlockRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewBlockRepository(l logger.Logger, db *sql.DB) *BlockRepository {
	return &BlockRepository{
		logger: l,
		db:     db,
	}
}

func (r *BlockRepository) Add(
	ctx context.Context, p *request.BlockAddRequestRepositoryDto) (*entity.BlockEntity, error) {
	query := "INSERT INTO profile_blocks (session_id, blocked_user_session_id, is_blocked, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.BlockedUserSessionId, &p.IsBlocked, &p.CreatedAt,
		&p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorMessage := r.getErrorMessage("Add", "sql.ErrNoRows")
			r.logger.Debug(errorMessage, zap.Error(err))
			return nil, err
		}
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindById(ctx, id)
}

func (r *BlockRepository) Find(ctx context.Context, sessionId, blockedUserSessionId string) (*entity.BlockEntity, error) {
	p := &entity.BlockEntity{}
	query := "SELECT id, session_id, blocked_user_session_id, is_blocked, created_at, updated_at " +
		" FROM profile_blocks" +
		" WHERE session_id=$1 AND blocked_user_session_id=$2"
	row := r.db.QueryRowContext(ctx, query, sessionId, blockedUserSessionId)
	err := row.Scan(&p.Id, &p.SessionId, &p.BlockedUserSessionId, &p.IsBlocked, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		errorMessage := r.getErrorMessage("Find", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *BlockRepository) FindById(ctx context.Context, id uint64) (*entity.BlockEntity, error) {
	p := &entity.BlockEntity{}
	query := "SELECT id, session_id, blocked_user_session_id, is_blocked, created_at, updated_at " +
		" FROM profile_blocks" +
		" WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&p.Id, &p.SessionId, &p.BlockedUserSessionId, &p.IsBlocked, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *BlockRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathBlock)
}
