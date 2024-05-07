package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/entity"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/sirupsen/logrus"
)

// IDigitalWallet defines the interface for the DigitalWallet repository.
type IDigitalWallet interface {
	// GetUserByID get user by id from cache.
	GetUserByID(ctx context.Context, id uint) (entity.User, error)
	// SetUserByID set user by id to cache.
	SetUserByID(ctx context.Context, user entity.User, expire time.Duration) error
	// DelUserByID delete user by id in cache.
	DelUserByID(ctx context.Context, id uint) error
}

// DigitalWallet represents the implementation of the DigitalWallet repository.
type DigitalWallet struct {
	cfg   *config.Config
	redis *Redis
}

var _ IDigitalWallet = &DigitalWallet{}

// NewDigitalWallet creates a new instance of the ErajolBike repository.
func NewDigitalWallet(cfg *config.Config, redis *Redis) *DigitalWallet {
	return &DigitalWallet{
		cfg:   cfg,
		redis: redis,
	}
}

///////////////////////////////// redis cache key /////////////////////////////////

const digitalWalletRedisKeyPrefix = "digital_wallet"

func (d *DigitalWallet) keyUserByID(id uint) string {
	keyList := []string{d.cfg.GetAppName(), digitalWalletRedisKeyPrefix, "UserByID", strconv.FormatUint(uint64(id), 10)}
	return strings.Join(keyList, ":")
}

///////////////////////////////// redis cache key /////////////////////////////////

// GetUserByID implements IDigitalWallet.
func (d *DigitalWallet) GetUserByID(ctx context.Context, id uint) (entity.User, error) {
	val, err := d.redis.client.Get(ctx, d.keyUserByID(id)).Result()
	if err != nil {
		return entity.User{}, trace.Wrap(err)
	}

	user := entity.User{}
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		err := fmt.Errorf("able to get value from redis but error when json unmarshal, will try to delete redis cache key: %w", err)

		errDel := d.DelUserByID(ctx, id)
		if errDel != nil {
			logrus.Warn(trace.Wrap(errDel))
		}

		return entity.User{}, trace.Wrap(err)
	}

	return user, nil
}

// SetUserByID implements IDigitalWallet.
func (d *DigitalWallet) SetUserByID(ctx context.Context, user entity.User, expire time.Duration) error {
	jsonByte, err := json.Marshal(user)
	if err != nil {
		return trace.Wrap(err)
	}

	err = d.redis.client.Set(ctx, d.keyUserByID(user.ID), string(jsonByte), expire).Err()
	if err != nil {
		return trace.Wrap(err)
	}

	return nil
}

// DelUserByID implements IDigitalWallet.
func (d *DigitalWallet) DelUserByID(ctx context.Context, id uint) error {
	err := d.redis.client.Del(ctx, d.keyUserByID(id)).Err()
	if err != nil {
		err := fmt.Errorf("error delete redis cache key '%s': %w", d.keyUserByID(id), err)
		return trace.Wrap(err)
	}
	return nil
}
