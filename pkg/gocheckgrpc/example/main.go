// Package main -.
package main

import (
	"context"
	"net"

	"github.com/Hidayathamir/gocheck/pkg/gocheckgrpc"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(net.JoinHostPort("localhost", "11010"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	fatalIfErr(err)
	defer func() {
		err := conn.Close()
		warnIfErr(err)
	}()

	gocheckgrpcDigitalWalletClient := gocheckgrpc.NewDigitalWalletClient(conn)
	req := &gocheckgrpc.ReqDigitalWalletTransfer{
		SenderId:    1,
		RecipientId: 2,      //nolint:gomnd
		Amount:      500000, //nolint:gomnd
	}
	res, err := gocheckgrpcDigitalWalletClient.Transfer(context.Background(), req)
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
