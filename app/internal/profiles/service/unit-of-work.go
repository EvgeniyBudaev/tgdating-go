package service

import (
	"context"
	"database/sql"
)

type UnitOfWork struct {
	tx                  *sql.Tx
	blockRepository     BlockRepository
	complaintRepository ComplaintRepository
	filterRepository    FilterRepository
	imageRepository     ImageRepository
	likeRepository      LikeRepository
	navigatorRepository NavigatorRepository
	profileRepository   ProfileRepository
	telegramRepository  TelegramRepository
}

func NewUnitOfWork(
	tx *sql.Tx,
	br BlockRepository,
	cr ComplaintRepository,
	fr FilterRepository,
	ir ImageRepository,
	lr LikeRepository,
	nr NavigatorRepository,
	pr ProfileRepository,
	tr TelegramRepository) *UnitOfWork {
	return &UnitOfWork{
		tx:                  tx,
		blockRepository:     br,
		complaintRepository: cr,
		filterRepository:    fr,
		imageRepository:     ir,
		likeRepository:      lr,
		navigatorRepository: nr,
		profileRepository:   pr,
		telegramRepository:  tr,
	}
}

func (unit *UnitOfWork) BlockRepository() BlockRepository {
	return unit.blockRepository
}

func (unit *UnitOfWork) ComplaintRepository() ComplaintRepository {
	return unit.complaintRepository
}

func (unit *UnitOfWork) FilterRepository() FilterRepository {
	return unit.filterRepository
}

func (unit *UnitOfWork) ImageRepository() ImageRepository {
	return unit.imageRepository
}

func (unit *UnitOfWork) LikeRepository() LikeRepository {
	return unit.likeRepository
}

func (unit *UnitOfWork) NavigatorRepository() NavigatorRepository {
	return unit.navigatorRepository
}

func (unit *UnitOfWork) ProfileRepository() ProfileRepository {
	return unit.profileRepository
}

func (unit *UnitOfWork) TelegramRepository() TelegramRepository {
	return unit.telegramRepository
}

func (unit *UnitOfWork) Commit(ctx context.Context) error {
	return unit.tx.Commit()
}

func (unit *UnitOfWork) Rollback(ctx context.Context) error {
	return unit.tx.Rollback()
}
