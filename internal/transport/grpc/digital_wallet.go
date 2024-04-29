package grpc

import (
	"context"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/dto"
	"github.com/Hidayathamir/gocheck/internal/transport/grpc/middleware"
	"github.com/Hidayathamir/gocheck/internal/usecase"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
)

// DigitalWallet -.
type DigitalWallet struct {
	gocheckgrpc.UnimplementedDigitalWalletServer

	cfg                  *config.Config
	usecaseDigitalWallet usecase.IDigitalWallet
}

var _ gocheckgrpc.DigitalWalletServer = &DigitalWallet{}

// NewDigitalWallet -.
func NewDigitalWallet(cfg *config.Config, usecaseDigitalWallet usecase.IDigitalWallet) *DigitalWallet {
	return &DigitalWallet{
		cfg:                  cfg,
		usecaseDigitalWallet: usecaseDigitalWallet,
	}
}

// Transfer implements gocheckgrpc.DigitalWalletServer.
func (d *DigitalWallet) Transfer(ctx context.Context, req *gocheckgrpc.ReqDigitalWalletTransfer) (*gocheckgrpc.ResDigitalWalletTransfer, error) {
	auth, err := middleware.GetAuthFromCtx(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	reqTransfer := dto.ReqTransfer{
		SenderID:    auth.UserID,
		RecipientID: uint(req.GetRecipientId()),
		Amount:      int(req.GetAmount()),
	}

	resTransfer, err := d.usecaseDigitalWallet.Transfer(ctx, reqTransfer)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	res := &gocheckgrpc.ResDigitalWalletTransfer{
		Id: uint64(resTransfer.ID),
	}

	return res, nil
}
