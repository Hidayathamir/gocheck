package app

import (
	transporthttp "github.com/Hidayathamir/gocheck/internal/transport/http"
	"github.com/gin-gonic/gin"
)

func registerHTTPRouter(
	ginEngine *gin.Engine,
	transporthttpDigitalWallet *transporthttp.DigitalWallet,
) {
	apiGroup := ginEngine.Group("api")
	{
		v1Group := apiGroup.Group("v1")
		{
			digitalWalletGroup := v1Group.Group("digital-wallet")
			{
				digitalWalletGroup.POST("transfer", transporthttpDigitalWallet.Transfer)
			}
		}
	}
}
