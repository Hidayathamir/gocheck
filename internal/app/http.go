package app

import (
	"github.com/Hidayathamir/gocheck/internal/config"
	transporthttp "github.com/Hidayathamir/gocheck/internal/transport/http"
	"github.com/Hidayathamir/gocheck/internal/usecase"
	"github.com/gin-gonic/gin"
)

func registerHTTPRouter(cfg *config.Config, ginEngine *gin.Engine, usecaseDigitalWallet *usecase.DigitalWallet) {
	tDigitalWallet := transporthttp.NewDigitalWallet(cfg, usecaseDigitalWallet)

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
