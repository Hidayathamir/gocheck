package injection

import (
	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/repo"
	"github.com/Hidayathamir/gocheck/internal/repo/cache"
	"github.com/Hidayathamir/gocheck/internal/repo/db"
	"github.com/Hidayathamir/gocheck/internal/repomw/repologgermw"
	"github.com/Hidayathamir/gocheck/internal/usecase"
	"github.com/Hidayathamir/gocheck/internal/usecasemw/usecaseloggermw"
)

// InitUsecaseDigitalWallet initializes the DigitalWallet usecase.
func InitUsecaseDigitalWallet(cfg *config.Config, pg *db.Postgres, redis *cache.Redis) usecase.IDigitalWallet {
	cacheDigitalWallet := cache.NewDigitalWallet(cfg, redis)

	var repoDigitalWallet repo.IDigitalWallet
	repoDigitalWallet = repo.NewDigitalWallet(cfg, pg, cacheDigitalWallet)
	repoDigitalWallet = repologgermw.NewDigitalWallet(cfg, repoDigitalWallet)

	var usecaseDigitalWallet usecase.IDigitalWallet
	usecaseDigitalWallet = usecase.NewDigitalWallet(cfg, pg.TxManager, repoDigitalWallet)
	usecaseDigitalWallet = usecaseloggermw.NewDigitalWallet(cfg, usecaseDigitalWallet)

	return usecaseDigitalWallet
}
