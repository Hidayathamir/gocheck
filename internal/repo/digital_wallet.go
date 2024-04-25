package repo

import (
	"context"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/repo/cache"
	"github.com/Hidayathamir/gocheck/internal/repo/db"
	"github.com/Hidayathamir/gocheck/internal/table"
	"github.com/Hidayathamir/gocheck/pkg/gocheck"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/sirupsen/logrus"
)

// IDigitalWallet -.
type IDigitalWallet interface {
	GetUserByID(ctx context.Context, id uint) (table.User, error)
	CreateTransaction(ctx context.Context, transaction table.Transaction) (uint, error)
	UpdateUserBalance(ctx context.Context, userID uint, balance int) error
}

// DigitalWallet implements IDigitalWallet.
type DigitalWallet struct {
	cfg                config.Config
	pg                 *db.Postgres
	cacheDigitalWallet cache.IDigitalWallet
}

var _ IDigitalWallet = &DigitalWallet{}

// NewDigitalWallet -.
func NewDigitalWallet(cfg config.Config, pg *db.Postgres, cacheDigitalWallet cache.IDigitalWallet) *DigitalWallet {
	return &DigitalWallet{
		cfg:                cfg,
		pg:                 pg,
		cacheDigitalWallet: cacheDigitalWallet,
	}
}

// GetUserByID implements IDigitalWallet.
func (d *DigitalWallet) GetUserByID(ctx context.Context, id uint) (table.User, error) {
	user, err := d.cacheDigitalWallet.GetUserByID(ctx, id)
	if err == nil {
		return user, nil
	}

	db := d.pg.DB

	if tx, ok := d.pg.TxManager.GetTx(ctx); ok {
		db = tx
	}

	user = table.User{}
	err = db.Last(&user, id).Error
	if err != nil {
		return table.User{}, trace.Wrap(err)
	}

	err = d.cacheDigitalWallet.SetUserByID(ctx, user, gocheck.DefaultCacheExpire)
	if err != nil {
		logrus.Warn(trace.Wrap(err))
	}

	return user, nil
}

// CreateTransaction implements IDigitalWallet.
func (d *DigitalWallet) CreateTransaction(ctx context.Context, transaction table.Transaction) (uint, error) {
	db := d.pg.DB

	if tx, ok := d.pg.TxManager.GetTx(ctx); ok {
		db = tx
	}

	err := db.Create(&transaction).Error
	if err != nil {
		return 0, trace.Wrap(err)
	}

	return transaction.ID, err
}

// UpdateUserBalance implements IDigitalWallet.
func (d *DigitalWallet) UpdateUserBalance(ctx context.Context, userID uint, balance int) error {
	db := d.pg.DB

	if tx, ok := d.pg.TxManager.GetTx(ctx); ok {
		db = tx
	}

	err := db.Table("users").Where("id = ?", userID).Update("balance", balance).Error
	if err != nil {
		return trace.Wrap(err)
	}

	return nil
}