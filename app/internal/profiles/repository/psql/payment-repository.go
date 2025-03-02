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
	errorFilePathPayment = "internal/repository/psql/payment-repository.go"
)

type PaymentRepository struct {
	logger logger.Logger
	db     *sql.DB
}

func NewPaymentRepository(l logger.Logger, db *sql.DB) *PaymentRepository {
	return &PaymentRepository{
		logger: l,
		db:     db,
	}
}

func (r *PaymentRepository) Add(
	ctx context.Context, p *request.PaymentAddRequestRepositoryDto) (*response.ResponseDto, error) {
	query := "INSERT INTO dating.profile_payments (telegram_user_id, price, currency, tariff, created_at," +
		" available_until)" +
		" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, &p.TelegramUserId, &p.Price, &p.Currency, &p.Tariff, &p.CreatedAt,
		&p.AvailableUntil)
	id := uint64(0)
	err := row.Scan(&id)
	if err != nil {
		errorMessage := r.getErrorMessage("Add", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, err
	}
	paymentResponse := &response.ResponseDto{
		Success: true,
	}
	return paymentResponse, nil
}

func (r *PaymentRepository) FindLastByTelegramUserId(
	ctx context.Context, telegramUserId string) (*entity.PaymentEntity, error) {
	p := &entity.PaymentEntity{}
	query := "SELECT id, telegram_user_id, price, currency, tariff, created_at, available_until" +
		" FROM dating.profile_payments" +
		" WHERE telegram_user_id = $1" +
		" ORDER BY created_at DESC" +
		" LIMIT 1"
	row := r.db.QueryRowContext(ctx, query, telegramUserId)
	err := row.Scan(&p.Id, &p.TelegramUserId, &p.Price, &p.Currency, &p.Tariff, &p.CreatedAt, &p.AvailableUntil)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		errorMessage := r.getErrorMessage("FindLastByTelegramUserId", "Scan")
		r.logger.Debug(errorMessage, zap.Error(err))
		return nil, nil
	}
	return p, nil
}

func (r *PaymentRepository) getErrorMessage(repositoryMethodName string, callMethodName string) string {
	return fmt.Sprintf("error func %s, method %s by path %s", repositoryMethodName, callMethodName,
		errorFilePathPayment)
}
