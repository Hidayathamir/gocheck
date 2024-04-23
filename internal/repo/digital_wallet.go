package repo

import (
	"context"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/pkg/trace"
	"github.com/Hidayathamir/gocheck/internal/repo/cache"
	"github.com/Hidayathamir/gocheck/internal/repo/db"
	"github.com/Hidayathamir/gocheck/internal/table"
	"github.com/Hidayathamir/gocheck/pkg/gocheck"
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
	cfg   config.Config
	pg    *db.Postgres
	redis cache.IDigitalWallet
}

var _ IDigitalWallet = &DigitalWallet{}

// NewDigitalWallet -.
func NewDigitalWallet(cfg config.Config, pg *db.Postgres, redis cache.IDigitalWallet) *DigitalWallet {
	return &DigitalWallet{
		cfg:   cfg,
		pg:    pg,
		redis: redis,
	}
}

// GetUserByID implements IDigitalWallet.
func (d *DigitalWallet) GetUserByID(ctx context.Context, id uint) (table.User, error) {
	user, err := d.redis.GetUserByID(ctx, id)
	if err == nil {
		return user, nil
	}

	// TODO: QUERY DB

	err = d.redis.SetUserByID(ctx, user, gocheck.DefaultCacheExpire)
	if err != nil {
		logrus.Warn(trace.Wrap(err))
	}

	panic("unimplemented")
}

// CreateTransaction implements IDigitalWallet.
func (d *DigitalWallet) CreateTransaction(context.Context, table.Transaction) (uint, error) {
	panic("unimplemented")
}

// UpdateUserBalance implements IDigitalWallet.
func (d *DigitalWallet) UpdateUserBalance(context.Context, uint, int) error {
	panic("unimplemented")
}
