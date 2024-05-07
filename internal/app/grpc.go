package app

import (
	transportgrpc "github.com/Hidayathamir/gocheck/internal/transport/grpc"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
	"google.golang.org/grpc"
)

func registerGRPCServer(grpcServer *grpc.Server, transportgrpcDigitalWallet *transportgrpc.DigitalWallet) {
	gocheckgrpc.RegisterDigitalWalletServer(grpcServer, transportgrpcDigitalWallet)
}
