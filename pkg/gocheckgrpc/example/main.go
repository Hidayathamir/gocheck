// Package main -.
package main

import (
	"context"
	"net"

	"github.com/Hidayathamir/gocheck/pkg/gocheck"
	"github.com/Hidayathamir/gocheck/pkg/gocheckgrpc"
	"github.com/Hidayathamir/gocheck/pkg/gocheckgrpcmiddleware"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//nolint:gomnd
func main() {
	conn, err := grpc.Dial(net.JoinHostPort("localhost", "11010"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	fatalIfErr(err)
	defer func() {
		err := conn.Close()
		warnIfErr(err)
	}()

	client := gocheckgrpc.NewDigitalWalletClient(conn)

	ctx := context.Background()
	auth := gocheck.Authorization{UserID: 1}
	ctx, err = gocheckgrpcmiddleware.SetAuthToCtx(ctx, auth)
	fatalIfErr(err)

	req := &gocheckgrpc.ReqDigitalWalletTransfer{
		RecipientId: 2,
		Amount:      10000,
	}

	res, err := client.Transfer(ctx, req)
	fatalIfErr(err)

	logrus.Info(res)
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(trace.Wrap(err, trace.WithSkip(1)))
	}
}

func warnIfErr(err error) {
	if err != nil {
		logrus.Warn(trace.Wrap(err, trace.WithSkip(1)))
	}
}
