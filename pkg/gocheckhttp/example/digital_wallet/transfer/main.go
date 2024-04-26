// Package main -.
package main

import (
	"context"

	"github.com/Hidayathamir/gocheck/pkg/gocheckhttp"
	"github.com/Hidayathamir/gocheck/pkg/gocheckhttpmiddleware"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/sirupsen/logrus"
)

//nolint:gomnd
func main() {
	base := "http://localhost:10010"

	// new digital wallet client http
	client := gocheckhttp.NewDigitalWalletClient(base)

	// prepare request
	ctx := context.Background()
	auth := gocheckhttpmiddleware.Authorization{UserID: 1}

	req := gocheckhttp.ReqDigitalWalletTransfer{
		RecipientID: 2,
		Amount:      10000,
	}

	// hit api digital wallet transfer via http
	res, err := client.Transfer(ctx, auth, req)
	fatalIfErr(err)

	// print response
	logrus.Info("transfer id = ", res.ID)
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(trace.Wrap(err, trace.WithSkip(1)))
	}
}
