package cache

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/table"
)

// IDigitalWallet -.
type IDigitalWallet interface {
	GetUserByID(ctx context.Context, id uint) (table.User, error)
	SetUserByID(ctx context.Context, user table.User, expire time.Duration) error
}

// DigitalWallet implements IDigitalWallet.
type DigitalWallet struct {
	cfg   config.Config
	redis *Redis
}

var _ IDigitalWallet = &DigitalWallet{}

// NewDigitalWallet -.
func NewDigitalWallet(cfg config.Config, redis *Redis) *DigitalWallet {
	return &DigitalWallet{
		cfg:   cfg,
		redis: redis,
	}
}

///////////////////////////////// redis cache key /////////////////////////////////

const digitalWalletRedisKeyPrefix = "digital_wallet" //nolint:unused

func (d *DigitalWallet) keyUserByID(id uint) string { //nolint:unused
	keyList := []string{d.cfg.GetAppName(), digitalWalletRedisKeyPrefix, "UserByID", strconv.FormatUint(uint64(id), 10)}
	return strings.Join(keyList, ":")
}

///////////////////////////////// redis cache key /////////////////////////////////

// GetUserByID implements IDigitalWallet.
func (d *DigitalWallet) GetUserByID(context.Context, uint) (table.User, error) {
	panic("unimplemented")
}

// SetUserByID implements IDigitalWallet.
func (d *DigitalWallet) SetUserByID(context.Context, table.User, time.Duration) error {
	panic("unimplemented")
}
