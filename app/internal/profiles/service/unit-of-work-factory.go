package service

import (
	"database/sql"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/logger"
	"github.com/EvgeniyBudaev/tgdating-go/app/internal/profiles/repository/psql"
)

type UnitOfWorkFactory struct {
	logger logger.Logger
	db     *sql.DB
}

func NewUnitOfWorkFactory(log logger.Logger, db *sql.DB) *UnitOfWorkFactory {
	return &UnitOfWorkFactory{
		logger: log,
		db:     db,
	}
}

func (factory *UnitOfWorkFactory) CreateUnit() *UnitOfWork {
	tx, err := factory.db.Begin()
	if err != nil {
		panic(err)
	}

	return NewUnitOfWork(
		tx,
		psql.NewBlockRepository(factory.logger, factory.db),
		psql.NewComplaintRepository(factory.logger, factory.db),
		psql.NewFilterRepository(factory.logger, factory.db),
		psql.NewImageRepository(factory.logger, factory.db),
		psql.NewLikeRepository(factory.logger, factory.db),
		psql.NewNavigatorRepository(factory.logger, factory.db),
		psql.NewProfileRepository(factory.logger, factory.db),
		psql.NewTelegramRepository(factory.logger, factory.db),
		psql.NewStatusRepository(factory.logger, factory.db),
	)
}
