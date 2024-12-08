package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"go.uber.org/zap"
)

const (
	errorFilePathBlock = "internal/repository/psql/block-repository.go"
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
	ctx context.Context, p *request.BlockAddRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "INSERT INTO dating.profile_blocks (telegram_user_id, blocked_telegram_user_id, is_blocked, created_at," +
		" updated_at)" +
		" VALUES ($1, $2, $3, $4, $5)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.BlockedTelegramUserId, &p.IsBlocked, &p.CreatedAt,
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
	blockResponse := &response.ResponseDto{
		Success: true,
	}
	return blockResponse, nil
}

func (r *BlockRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathBlock)
}
