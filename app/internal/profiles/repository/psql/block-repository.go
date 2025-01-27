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
	"time"
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

func (r *BlockRepository) Update(
	ctx context.Context, telegramUserId, blockedTelegramUserId string) (*response.ResponseDto, error) {
	updatedAt := time.Now().UTC()
	query := "UPDATE dating.profile_blocks SET is_blocked = $3, updated_at = $4" +
		" WHERE telegram_user_id = $1 AND blocked_telegram_user_id = $2"
	_, err := r.db.ExecContext(ctx, query, telegramUserId, blockedTelegramUserId, true, updatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("Update", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	blockResponse := &response.ResponseDto{
		Success: true,
	}
	return blockResponse, nil
}

func (r *BlockRepository) GetBlockedList(ctx context.Context,
	telegramUserId string) (*response.BlockedListResponseDto, error) {
	query := "WITH filtered_profiles AS (" +
		"SELECT pb.id, pb.telegram_user_id, pb.blocked_telegram_user_id, pb.is_blocked, pb.created_at, pb.updated_at," +
		" (SELECT url FROM dating.profile_images pi" +
		" JOIN dating.profile_image_statuses pis ON pi.id = pis.id" +
		" WHERE pi.telegram_user_id = pb.blocked_telegram_user_id AND" +
		" pis.is_blocked = false AND pis.is_private = false" +
		" ORDER BY pi.created_at DESC LIMIT 1) AS url" +
		" FROM dating.profile_blocks pb" +
		" WHERE pb.telegram_user_id = $1 AND pb.is_blocked = true" +
		" )" +
		" SELECT blocked_telegram_user_id, url" +
		" FROM filtered_profiles" +
		" ORDER BY updated_at DESC"
	rows, err := r.db.QueryContext(ctx, query, telegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("GetBlockedList", "QueryContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	content := make([]*response.BlockedListItemResponseDto, 0)
	for rows.Next() {
		p := &response.BlockedListItemResponseDto{}
		err := rows.Scan(&p.BlockedTelegramUserId, &p.Url)
		if err != nil {
			errorMessage := r.getErrorMessage("GetBlockedList", "Scan")
			r.logger.Debug(errorMessage, zap.Error(err))
			continue
		}
		content = append(content, p)
	}
	blockedList := &response.BlockedListResponseDto{
		Content: content,
	}
	return blockedList, nil
}

func (r *BlockRepository) FindBlock(ctx context.Context, telegramUserId, blockedTelegramUserId string) (*entity.BlockEntity, error) {
	p := &entity.BlockEntity{}
	query := "SELECT id, telegram_user_id, blocked_telegram_user_id, is_blocked, created_at, updated_at " +
		" FROM dating.profile_blocks" +
		" WHERE telegram_user_id = $1 AND blocked_telegram_user_id = $2"
	row := r.db.QueryRowContext(ctx, query, telegramUserId, blockedTelegramUserId)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.BlockedTelegramUserId, &p.IsBlocked, &p.CreatedAt, &p.UpdatedAt)
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

func (r *BlockRepository) Unblock(ctx context.Context, p *request.UnblockRequestDto) (*response.ResponseDto, error) {
	query := "UPDATE dating.profile_blocks SET is_blocked = $1, updated_at = $2" +
		" WHERE telegram_user_id = $3 AND blocked_telegram_user_id = $4"
	updatedAt := time.Now().UTC()
	_, err := r.db.ExecContext(ctx, query, false, updatedAt, p.TelegramUserId, p.BlockedTelegramUserId)
	if err != nil {
		errorMessage := r.getErrorMessage("Unblock", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	blockResponse := &response.ResponseDto{
		Success: true,
	}
	return blockResponse, nil
}

func (r *BlockRepository) DeleteRelatedProfiles(
	ctx context.Context, id string) (*response.ResponseDto, error) {
	query := "DELETE FROM dating.profile_blocks WHERE blocked_telegram_user_id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteRelatedProfiles", "ExecContext")
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
