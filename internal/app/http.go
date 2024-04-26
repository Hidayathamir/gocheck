package app

import (
	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/repo"
	"github.com/Hidayathamir/gocheck/internal/repo/cache"
	"github.com/Hidayathamir/gocheck/internal/repo/db"
	transporthttp "github.com/Hidayathamir/gocheck/internal/transport/http"
	"github.com/Hidayathamir/gocheck/internal/usecase"
	"github.com/gin-gonic/gin"
)

func registerHTTPRouter(cfg config.Config, ginEngine *gin.Engine, pg *db.Postgres, redis *cache.Redis) {
	tDigitalWallet := injectionDigitalWalletHTTP(cfg, pg, redis)

	apiGroup := ginEngine.Group("api")
	{
		v1Group := apiGroup.Group("v1")
		{
			digitalWalletGroup := v1Group.Group("digital-wallet")
			{
				digitalWalletGroup.POST("transfer", tDigitalWallet.Transfer)
			}
		}
	}
}

func injectionDigitalWalletHTTP(cfg config.Config, pg *db.Postgres, redis *cache.Redis) *transporthttp.DigitalWallet {
	cacheDigitalWallet := cache.NewDigitalWallet(cfg, redis)
	repoDigitalWallet := repo.NewDigitalWallet(cfg, pg, cacheDigitalWallet)

	usecaseDigitalWallet := usecase.NewDigitalWallet(cfg, pg.TxManager, repoDigitalWallet)

	transportDigitalWallet := transporthttp.NewDigitalWallet(cfg, usecaseDigitalWallet)

	return transportDigitalWallet
}
