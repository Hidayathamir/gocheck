// Package main -.
package main

import (
	"context"

	"github.com/Hidayathamir/gocheck/pkg/gocheckhttp"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/sirupsen/logrus"
)

func main() {
	base := "http://localhost:10010"

	gocheckhttpDigitalWalletClient := gocheckhttp.NewDigitalWalletClient(base)
	req := gocheckhttp.ReqDigitalWalletTransfer{
		SenderID:    1,
		RecipientID: 2,      //nolint:gomnd
		Amount:      500000, //nolint:gomnd
	}
	res, err := gocheckhttpDigitalWalletClient.Transfer(context.Background(), req)
	fatalIfErr(err)

	logrus.Info(res.Body)
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(trace.Wrap(err, trace.WithSkip(1)))
	}
}
