package app

import (
	"github.com/Hidayathamir/gocheck/internal/config"
	transportgrpc "github.com/Hidayathamir/gocheck/internal/transport/grpc"
	"github.com/Hidayathamir/gocheck/internal/usecase"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
	"google.golang.org/grpc"
)

func registerGRPCServer(cfg *config.Config, grpcServer *grpc.Server, usecaseDigitalWallet *usecase.DigitalWallet) {
	tDigitalWallet := transportgrpc.NewDigitalWallet(cfg, usecaseDigitalWallet)

	gocheckgrpc.RegisterDigitalWalletServer(grpcServer, tDigitalWallet)
}
