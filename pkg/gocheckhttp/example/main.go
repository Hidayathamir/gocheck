// Package main -.
package main

import (
	"context"

	"github.com/Hidayathamir/gocheck/pkg/gocheck"
	"github.com/Hidayathamir/gocheck/pkg/gocheckhttp"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/sirupsen/logrus"
)

//nolint:gomnd
func main() {
	base := "http://localhost:10010"

	client := gocheckhttp.NewDigitalWalletClient(base)

	ctx := context.Background()
	auth := gocheck.Authorization{UserID: 1}

	req := gocheckhttp.ReqDigitalWalletTransfer{
		RecipientID: 2,
		Amount:      10000,
	}

	res, err := client.Transfer(ctx, auth, req)
	fatalIfErr(err)

	logrus.Info(res)
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(trace.Wrap(err, trace.WithSkip(1)))
	}
}
