package grpc

import (
	"context"

	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/dto"
	"github.com/Hidayathamir/gocheck/internal/usecase"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
)

// DigitalWallet represents the gRPC server for the DigitalWallet service.
type DigitalWallet struct {
	gocheckgrpc.UnimplementedDigitalWalletServer

	cfg                  *config.Config
	usecaseDigitalWallet usecase.IDigitalWallet
}

var _ gocheckgrpc.DigitalWalletServer = &DigitalWallet{}

// NewDigitalWallet creates a new instance of DigitalWallet gRPC server.
func NewDigitalWallet(cfg *config.Config, usecaseDigitalWallet usecase.IDigitalWallet) *DigitalWallet {
	return &DigitalWallet{
		cfg:                  cfg,
		usecaseDigitalWallet: usecaseDigitalWallet,
	}
}

// Transfer implements gocheckgrpc.DigitalWalletServer.
func (d *DigitalWallet) Transfer(ctx context.Context, req *gocheckgrpc.ReqDigitalWalletTransfer) (*gocheckgrpc.ResDigitalWalletTransfer, error) {
	reqTransfer := dto.ReqDigitalWalletTransfer{}
	err := reqTransfer.LoadFromReqGRPC(ctx, req)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	resTransfer, err := d.usecaseDigitalWallet.Transfer(ctx, reqTransfer)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	res := resTransfer.ToResGRPC()

	return res, nil
}
