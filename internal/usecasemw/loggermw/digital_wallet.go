package loggermw

import (
	"context"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/dto"
	"github.com/Hidayathamir/gocheck/internal/usecase"
	"github.com/Hidayathamir/gocheck/pkg/trace"
)

// DigitalWallet represents the implementation of the DigitalWallet logger middleware.
type DigitalWallet struct {
	cfg  *config.Config
	next usecase.IDigitalWallet
}

var _ usecase.IDigitalWallet = &DigitalWallet{}

// NewDigitalWallet creates a new instane of DigitalWallet logger middleware.
func NewDigitalWallet(cfg *config.Config, next usecase.IDigitalWallet) *DigitalWallet {
	return &DigitalWallet{
		cfg:  cfg,
		next: next,
	}
}

// Transfer implements usecase.IDigitalWallet.
func (d *DigitalWallet) Transfer(ctx context.Context, req dto.ReqDigitalWalletTransfer) (dto.ResDigitalWalletTransfer, error) {
	res, err := d.next.Transfer(ctx, req)

	log(ctx, trace.FuncName(), req, res, err)

	return res, err
}
