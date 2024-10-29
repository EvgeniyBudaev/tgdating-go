package psql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/entity"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/logger"
	"go.uber.org/zap"
)

const (
	errorFilePathComplaint = "internal/repository/psql/complaintRepository.go"
)

type ComplaintRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewComplaintRepository(l logger.Logger, db *sql.DB) *ComplaintRepository {
	return &ComplaintRepository{
		logger: l,
		db:     db,
	}
}

func (r *ComplaintRepository) Add(
	ctx context.Context, p *request.ComplaintAddRequestRepositoryDto) (*entity.ComplaintEntity, error) {
	query := "INSERT INTO profile_complaints (session_id, criminal_session_id, reason, is_deleted, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.SessionId, &p.CriminalSessionId, &p.Reason, &p.IsDeleted, &p.IsDeleted,
		&p.CreatedAt, &p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return r.FindById(ctx, id)
}

func (r *ComplaintRepository) FindById(ctx context.Context, id uint64) (*entity.ComplaintEntity, error) {
	p := &entity.ComplaintEntity{}
	query := "SELECT id, session_id, criminal_session_id, reason, is_deleted, created_at, updated_at " +
		" FROM profile_complaints" +
		" WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&p.Id, &p.SessionId, &p.CriminalSessionId, &p.Reason, &p.IsDeleted, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		errorMessage := r.getErrorMessage("FindById", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	return p, nil
}

func (r *ComplaintRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathComplaint)
}
