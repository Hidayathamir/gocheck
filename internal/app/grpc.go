package app

import (
	transportgrpc "github.com/Hidayathamir/gocheck/internal/transport/grpc"
	protobufgocheck "github.com/Hidayathamir/protobuf/gocheck"
	"google.golang.org/grpc"
)

func registerGRPCServer(grpcServer *grpc.Server, transportgrpcDigitalWallet *transportgrpc.DigitalWallet) {
	protobufgocheck.RegisterDigitalWalletServer(grpcServer, transportgrpcDigitalWallet)
}
