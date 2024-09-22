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

var (
	ErrNotRowsFoundBlock = errors.New("no rows found")
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

func (r *BlockRepository) AddBlock(
	ctx context.Context, p *request.BlockAddRequestRepositoryDto) (*entity.BlockEntity, error) {
	query := "INSERT INTO profile_blocks (session_id, blocked_user_session_id, is_blocked, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.BlockedUserSessionId, &p.IsBlocked, &p.CreatedAt,
		&p.UpdatedAt)
	if row == nil {
		errorMessage := r.getErrorMessage("AddBlock", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFoundBlock))
		return nil, ErrNotRowsFoundBlock
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

func (r *BlockRepository) FindBlockById(ctx context.Context, id uint64) (*entity.BlockEntity, error) {
	p := &entity.BlockEntity{}
	query := "SELECT id, session_id, blocked_user_session_id, is_blocked, created_at, updated_at " +
		" FROM profile_blocks" +
		" WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		errorMessage := r.getErrorMessage("FindBlockById", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(ErrNotRowsFoundBlock))
		return nil, ErrNotRowsFoundBlock
	}
	err := row.Scan(&p.Id, &p.SessionId, &p.BlockedUserSessionId, &p.IsBlocked, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindBlockById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *BlockRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathBlock)
}
