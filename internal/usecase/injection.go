package usecase

import (
	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/repo"
	"github.com/Hidayathamir/gocheck/internal/repo/cache"
	"github.com/Hidayathamir/gocheck/internal/repo/db"
)

// InitUsecaseDigitalWallet initializes the DigitalWallet usecase.
func InitUsecaseDigitalWallet(cfg *config.Config, pg *db.Postgres, redis *cache.Redis) *DigitalWallet {
	cacheDigitalWallet := cache.NewDigitalWallet(cfg, redis)
	repoDigitalWallet := repo.NewDigitalWallet(cfg, pg, cacheDigitalWallet)

	usecaseDigitalWallet := NewDigitalWallet(cfg, pg.TxManager, repoDigitalWallet)
	return usecaseDigitalWallet
}
