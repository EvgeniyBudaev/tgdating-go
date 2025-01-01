package psql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/request"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/dto/response"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"go.uber.org/zap"
)

const (
	errorFilePathComplaint = "internal/repository/psql/complaint-repository.go"
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
	ctx context.Context, p *request.ComplaintAddRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "INSERT INTO dating.profile_complaints (telegram_user_id, criminal_telegram_user_id, reason, created_at," +
		" updated_at)" +
		" VALUES ($1, $2, $3, $4, $5)" +
		" RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.CriminalTelegramUserId, &p.Reason, p.CreatedAt,
		&p.UpdatedAt)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	complaintResponse := &response.ResponseDto{
		Success: true,
	}
	return complaintResponse, nil
}

func (r *ComplaintRepository) DeleteRelatedProfiles(
	ctx context.Context, id string) (*response.ResponseDto, error) {
	query := "DELETE FROM dating.profile_complaints WHERE criminal_telegram_user_id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		errorMessage := r.getErrorMessage("DeleteRelatedProfiles", "ExecContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	complaintResponse := &response.ResponseDto{
		Success: true,
	}
	return complaintResponse, nil
}

func (r *ComplaintRepository) GetCountUserComplaintsByToday(
	ctx context.Context, telegramUserId string) (uint64, error) {
	query := "SELECT COUNT(*)" +
		" FROM dating.profiles p" +
		" JOIN dating.profile_complaints pc ON p.telegram_user_id = pc.telegram_user_id" +
		" WHERE pc.telegram_user_id = $1" +
		" AND DATE(pc.created_at) = CURRENT_DATE"
	var countUserComplaints uint64
	err := r.db.QueryRowContext(ctx, query, telegramUserId).Scan(&countUserComplaints)
	if err != nil {
		errorMessage := r.getErrorMessage("getTotalEntities", "QueryRowContext")
		r.logger.Debug(errorMessage, zap.Error(err))
		return 0, err
	}
	return countUserComplaints, nil
}

func (r *ComplaintRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathComplaint)
}
