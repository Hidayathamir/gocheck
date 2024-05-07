package grpc

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Hidayathamir/gocheck/internal/usecase/injection"
	"github.com/Hidayathamir/gocheck/pkg/gocheckgrpcmiddleware"
	"github.com/Hidayathamir/gocheck/pkg/m"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func TestIntegrationDigitalWalletTransfer(t *testing.T) {
	t.Parallel()

	t.Run("transfer success", func(t *testing.T) {
		t.Parallel()

		cfg, _, pg, redis := setup(t)

		usecaseDigitalWallet := injection.InitUsecaseDigitalWallet(cfg, pg, redis)

		tDigitalWallet := NewDigitalWallet(cfg, usecaseDigitalWallet)

		ctx := context.Background()
		auth := gocheckgrpcmiddleware.Authorization{UserID: 1}
		jsonByte, err := json.Marshal(auth)
		require.NoError(t, err)
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(m.Authorization, string(jsonByte)))

		req := &gocheckgrpc.ReqDigitalWalletTransfer{
			RecipientId: 2,
			Amount:      10000,
		}

		res, err := tDigitalWallet.Transfer(ctx, req)

		require.NoError(t, err)
		assert.NotEmpty(t, res)
	})
	t.Run("unauthenticated request should return error", func(t *testing.T) {
		t.Parallel()

		cfg, _, pg, redis := setup(t)

		usecaseDigitalWallet := injection.InitUsecaseDigitalWallet(cfg, pg, redis)

		tDigitalWallet := NewDigitalWallet(cfg, usecaseDigitalWallet)

		ctx := context.Background()

		req := &gocheckgrpc.ReqDigitalWalletTransfer{
			RecipientId: 2,
			Amount:      10000,
		}

		res, err := tDigitalWallet.Transfer(ctx, req)

		require.Error(t, err)
		assert.Empty(t, res)
	})
	t.Run("sender id not found should return transfer error", func(t *testing.T) {
		t.Parallel()

		cfg, _, pg, redis := setup(t)

		usecaseDigitalWallet := injection.InitUsecaseDigitalWallet(cfg, pg, redis)

		tDigitalWallet := NewDigitalWallet(cfg, usecaseDigitalWallet)

		ctx := context.Background()
		auth := gocheckgrpcmiddleware.Authorization{UserID: 1000000}
		jsonByte, err := json.Marshal(auth)
		require.NoError(t, err)
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(m.Authorization, string(jsonByte)))

		req := &gocheckgrpc.ReqDigitalWalletTransfer{
			RecipientId: 2,
			Amount:      10000,
		}

		res, err := tDigitalWallet.Transfer(ctx, req)

		require.Error(t, err)
		assert.Empty(t, res)
	})
	t.Run("recipient id not found should return transfer error", func(t *testing.T) {
		t.Parallel()

		cfg, _, pg, redis := setup(t)

		usecaseDigitalWallet := injection.InitUsecaseDigitalWallet(cfg, pg, redis)

		tDigitalWallet := NewDigitalWallet(cfg, usecaseDigitalWallet)

		ctx := context.Background()
		auth := gocheckgrpcmiddleware.Authorization{UserID: 1}
		jsonByte, err := json.Marshal(auth)
		require.NoError(t, err)
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(m.Authorization, string(jsonByte)))

		req := &gocheckgrpc.ReqDigitalWalletTransfer{
			RecipientId: 123123123123,
			Amount:      10000,
		}

		res, err := tDigitalWallet.Transfer(ctx, req)

		require.Error(t, err)
		assert.Empty(t, res)
	})
	t.Run("transfer to himself should return transfer error", func(t *testing.T) {
		t.Parallel()

		cfg, _, pg, redis := setup(t)

		usecaseDigitalWallet := injection.InitUsecaseDigitalWallet(cfg, pg, redis)

		tDigitalWallet := NewDigitalWallet(cfg, usecaseDigitalWallet)

		ctx := context.Background()
		auth := gocheckgrpcmiddleware.Authorization{UserID: 1}
		jsonByte, err := json.Marshal(auth)
		require.NoError(t, err)
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(m.Authorization, string(jsonByte)))

		req := &gocheckgrpc.ReqDigitalWalletTransfer{
			RecipientId: 1,
			Amount:      10000,
		}

		res, err := tDigitalWallet.Transfer(ctx, req)

		require.Error(t, err)
		assert.Empty(t, res)
	})
	t.Run("transfer lower than minimum amount should return transfer error", func(t *testing.T) {
		t.Parallel()

		cfg, _, pg, redis := setup(t)

		usecaseDigitalWallet := injection.InitUsecaseDigitalWallet(cfg, pg, redis)

		tDigitalWallet := NewDigitalWallet(cfg, usecaseDigitalWallet)

		ctx := context.Background()
		auth := gocheckgrpcmiddleware.Authorization{UserID: 1}
		jsonByte, err := json.Marshal(auth)
		require.NoError(t, err)
		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(m.Authorization, string(jsonByte)))

		req := &gocheckgrpc.ReqDigitalWalletTransfer{
			RecipientId: 2,
			Amount:      100,
		}

		res, err := tDigitalWallet.Transfer(ctx, req)

		require.Error(t, err)
		assert.Empty(t, res)
	})
}
