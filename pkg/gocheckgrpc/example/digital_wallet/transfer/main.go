// Package main is the entry point of the application.
package main

import (
	"context"
	"encoding/json"
	"net"

	"github.com/Hidayathamir/gocheck/pkg/gocheckgrpc"
	"github.com/Hidayathamir/gocheck/pkg/m"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	pbgocheck "github.com/Hidayathamir/protobuf/gocheck"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

//nolint:gomnd
func main() {
	conn, err := grpc.Dial(net.JoinHostPort("localhost", "11010"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	fatalIfErr(err)
	defer func() {
		err := conn.Close()
		warnIfErr(err)
	}()

	client := pbgocheck.NewDigitalWalletClient(conn)

	ctx := context.Background()
	auth := gocheckgrpc.Authorization{UserID: 1}
	jsonByte, err := json.Marshal(auth)
	fatalIfErr(err)

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(m.Authorization, string(jsonByte), m.TraceID, uuid.NewString()))

	req := &pbgocheck.ReqDigitalWalletTransfer{
		RecipientId: 2,
		Amount:      10000,
	}

	res, err := client.Transfer(ctx, req)
	fatalIfErr(err)

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
