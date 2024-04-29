// Package cache -.
package cache

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// Redis -.
type Redis struct {
	client *redis.Client
}

// NewRedis -.
func NewRedis(cfg *config.Config) (*Redis, error) {
	addr := net.JoinHostPort(cfg.GetRedisHost(), cfg.GetRedisPort())

	var redisClient *redis.Client
	var errInitRedis error
	const maxAttemptInitRedis = 10
	for i := 0; i < maxAttemptInitRedis; i++ {
		redisClient = redis.NewClient(&redis.Options{Addr: addr})

		errInitRedis = redisClient.Ping(context.Background()).Err()
		if errInitRedis == nil {
			break
		}

		errInitRedis = fmt.Errorf("error ping redis: %w", errInitRedis)

		logrus.
			WithField("attempt left", maxAttemptInitRedis-i-1).
			Warn(trace.Wrap(errInitRedis))

		time.Sleep(time.Second)
	}
	if errInitRedis != nil {
		errInitRedis := fmt.Errorf("error ping redis %d times: %w", maxAttemptInitRedis, errInitRedis)
		return nil, trace.Wrap(errInitRedis)
	}

	redis := &Redis{client: redisClient}

	logrus.Info("success create redis connection 🟢")

	return redis, nil
}
