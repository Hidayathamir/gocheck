// Package main is the entry point of the application.
package main

import (
	"context"
	"net"

	"github.com/Hidayathamir/gocheck/pkg/gocheckgrpcmiddleware"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	gocheckgrpc "github.com/Hidayathamir/protobuf/gocheck"
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

	// new digital wallet client grpc
	client := gocheckgrpc.NewDigitalWalletClient(conn)

	// prepare request
	ctx := context.Background()
	auth := gocheckgrpcmiddleware.Authorization{UserID: 1}
	ctx, err = gocheckgrpcmiddleware.SetAuthToCtx(ctx, auth)
	fatalIfErr(err)

	req := &gocheckgrpc.ReqDigitalWalletTransfer{
		RecipientId: 2,
		Amount:      10000,
	}

	// hit api digital wallet transfer via grpc
	res, err := client.Transfer(ctx, req)
	fatalIfErr(err)

	// print response
	logrus.Info("transfer id = ", res.GetId())
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
