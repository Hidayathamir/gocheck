package http

import (
	"net/http"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/dto"
	"github.com/Hidayathamir/gocheck/internal/usecase"
	"github.com/Hidayathamir/gocheck/pkg/gocheckhttp"
	"github.com/Hidayathamir/gocheck/pkg/gocheckhttpmiddleware"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/gin-gonic/gin"
)

// DigitalWallet -.
type DigitalWallet struct {
	cfg                  config.Config
	usecaseDigitalWallet usecase.IDigitalWallet
}

// NewDigitalWallet -.
func NewDigitalWallet(cfg config.Config, usecaseDigitalWallet usecase.IDigitalWallet) *DigitalWallet {
	return &DigitalWallet{
		cfg:                  cfg,
		usecaseDigitalWallet: usecaseDigitalWallet,
	}
}

// Transfer -.
func (d *DigitalWallet) Transfer(c *gin.Context) {
	auth, err := gocheckhttpmiddleware.GetAuthFromGinCtxHeader(c)
	if err != nil {
		err := trace.Wrap(err)
		res := gocheckhttp.ResDigitalWalletTransfer{Error: err.Error()}
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	req := gocheckhttp.ReqDigitalWalletTransfer{}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		err := trace.Wrap(err)
		res := gocheckhttp.ResDigitalWalletTransfer{Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	reqTransfer := dto.ReqTransfer{
		SenderID:    auth.UserID,
		RecipientID: req.RecipientID,
		Amount:      req.Amount,
	}

	resTransfer, err := d.usecaseDigitalWallet.Transfer(c, reqTransfer)
	if err != nil {
		err := trace.Wrap(err)
		res := gocheckhttp.ResDigitalWalletTransfer{Error: err.Error()}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := gocheckhttp.ResDigitalWalletTransfer{
		Data: gocheckhttp.ResDataDigitalWalletTransfer{
			ID: resTransfer.ID,
		},
	}

	c.JSON(http.StatusOK, res)
}
