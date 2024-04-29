package app

import (
	"github.com/Hidayathamir/gocheck/internal/config"
	"github.com/Hidayathamir/gocheck/internal/repo"
	"github.com/Hidayathamir/gocheck/internal/repo/cache"
	"github.com/Hidayathamir/gocheck/internal/repo/db"
	transportgrpc "github.com/Hidayathamir/gocheck/internal/transport/grpc"
	"github.com/Hidayathamir/gocheck/internal/usecase"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
	"google.golang.org/grpc"
)

func registerGRPCServer(cfg *config.Config, grpcServer *grpc.Server, pg *db.Postgres, redis *cache.Redis) {
	tDigitalWallet := injectionDigitalWalletGRPC(cfg, pg, redis)

	gocheckgrpc.RegisterDigitalWalletServer(grpcServer, tDigitalWallet)
}

func injectionDigitalWalletGRPC(cfg *config.Config, pg *db.Postgres, redis *cache.Redis) *transportgrpc.DigitalWallet {
	cacheDigitalWallet := cache.NewDigitalWallet(cfg, redis)
	repoDigitalWallet := repo.NewDigitalWallet(cfg, pg, cacheDigitalWallet)

	usecaseDigitalWallet := usecase.NewDigitalWallet(cfg, pg.TxManager, repoDigitalWallet)

	transportDigitalWallet := transportgrpc.NewDigitalWallet(cfg, usecaseDigitalWallet)

	return transportDigitalWallet
}
