package repo

import (
	"context"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/entity"
	"github.com/Hidayathamir/gocheck/internal/entity/table"
	"github.com/Hidayathamir/gocheck/internal/repo/cache"
	"github.com/Hidayathamir/gocheck/internal/repo/db"
	"github.com/Hidayathamir/gocheck/pkg/errutil"
	"github.com/Hidayathamir/gocheck/pkg/gocheck"
	"github.com/Hidayathamir/gocheck/pkg/q"
	"github.com/sirupsen/logrus"
)

// IDigitalWallet defines the interface for the DigitalWallet repository.
type IDigitalWallet interface {
	// GetUserByID get user by id from cache or database.
	GetUserByID(ctx context.Context, id uint) (entity.User, error)
	// CreateTransaction create new transaction record.
	CreateTransaction(ctx context.Context, transaction entity.Transaction) (uint, error)
	// UpdateUserBalance update user balance.
	UpdateUserBalance(ctx context.Context, userID uint, balance int) error
}

// DigitalWallet represents the implementation of the DigitalWallet repository.
type DigitalWallet struct {
	cfg                *config.Config
	pg                 *db.Postgres
	cacheDigitalWallet cache.IDigitalWallet
}

var _ IDigitalWallet = &DigitalWallet{}

// NewDigitalWallet creates a new instance of the DigitalWallet repository.
func NewDigitalWallet(cfg *config.Config, pg *db.Postgres, cacheDigitalWallet cache.IDigitalWallet) *DigitalWallet {
	return &DigitalWallet{
		cfg:                cfg,
		pg:                 pg,
		cacheDigitalWallet: cacheDigitalWallet,
	}
}

// GetUserByID implements IDigitalWallet.
func (d *DigitalWallet) GetUserByID(ctx context.Context, id uint) (entity.User, error) {
	user, err := d.cacheDigitalWallet.GetUserByID(ctx, id)
	if err == nil {
		return user, nil
	}

	db := d.pg.DB

	if tx, ok := d.pg.TxManager.GetTx(ctx); ok {
		db = tx
	}

	user = entity.User{}
	err = db.Last(&user, id).Error
	if err != nil {
		return entity.User{}, errutil.Wrap(err)
	}

	err = d.cacheDigitalWallet.SetUserByID(ctx, user, gocheck.DefaultCacheExpire)
	if err != nil {
		logrus.Warn(errutil.Wrap(err))
	}

	return user, nil
}

// CreateTransaction implements IDigitalWallet.
func (d *DigitalWallet) CreateTransaction(ctx context.Context, transaction entity.Transaction) (uint, error) {
	db := d.pg.DB

	if tx, ok := d.pg.TxManager.GetTx(ctx); ok {
		db = tx
	}

	err := db.Create(&transaction).Error
	if err != nil {
		return 0, errutil.Wrap(err)
	}

	return transaction.ID, err
}

// UpdateUserBalance implements IDigitalWallet.
func (d *DigitalWallet) UpdateUserBalance(ctx context.Context, userID uint, balance int) error {
	db := d.pg.DB

	if tx, ok := d.pg.TxManager.GetTx(ctx); ok {
		db = tx
	}

	err := db.
		Table(table.User.TableName()).
		Where(q.Equal(table.User.ID), userID).
		Update(table.User.Balance, balance).
		Error
	if err != nil {
		return errutil.Wrap(err)
	}

	return nil
}
