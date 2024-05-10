package http

import (
	"net/http"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/dto"
	"github.com/Hidayathamir/gocheck/internal/usecase"
	"github.com/Hidayathamir/gocheck/pkg/errutil"
	"github.com/Hidayathamir/gocheck/pkg/gocheckhttp"
	"github.com/gin-gonic/gin"
)

// DigitalWallet represents the HTTP server for the DigitalWallet service.
type DigitalWallet struct {
	cfg                  *config.Config
	usecaseDigitalWallet usecase.IDigitalWallet
}

// NewDigitalWallet creates a new instance of DigitalWallet HTTP server.
func NewDigitalWallet(cfg *config.Config, usecaseDigitalWallet usecase.IDigitalWallet) *DigitalWallet {
	return &DigitalWallet{
		cfg:                  cfg,
		usecaseDigitalWallet: usecaseDigitalWallet,
	}
}

// Transfer is the handler function for the Transfer endpoint.
func (d *DigitalWallet) Transfer(c *gin.Context) {
	reqTransfer := dto.ReqDigitalWalletTransfer{}
	err := reqTransfer.LoadFromReqHTTP(c)
	if err != nil {
		err := errutil.Wrap(err)
		res := gocheckhttp.ResDigitalWalletTransfer{Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	resTransfer, err := d.usecaseDigitalWallet.Transfer(c, reqTransfer)
	if err != nil {
		err := errutil.Wrap(err)
		res := gocheckhttp.ResDigitalWalletTransfer{Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := resTransfer.ToResHTTP()

	c.JSON(http.StatusOK, res)
}
