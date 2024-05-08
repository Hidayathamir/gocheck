package app

import (
	transportgrpc "github.com/Hidayathamir/gocheck/internal/transport/grpc"
	pbgocheck "github.com/Hidayathamir/protobuf/gocheck"
	"google.golang.org/grpc"
)

func registerGRPCServer(grpcServer *grpc.Server, transportgrpcDigitalWallet *transportgrpc.DigitalWallet) {
	pbgocheck.RegisterDigitalWalletServer(grpcServer, transportgrpcDigitalWallet)
}
