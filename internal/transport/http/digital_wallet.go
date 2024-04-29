package http

import (
	"net/http"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/dto"
	"github.com/Hidayathamir/gocheck/internal/transport/http/middleware"
	"github.com/Hidayathamir/gocheck/internal/usecase"
	"github.com/Hidayathamir/gocheck/pkg/gocheckhttp"
	"github.com/Hidayathamir/gocheck/pkg/trace"
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
	auth, err := middleware.GetAuthFromGinCtxHeader(c)
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
