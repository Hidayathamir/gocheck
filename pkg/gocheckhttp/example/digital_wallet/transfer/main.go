// Package main is the entry point of the application.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Hidayathamir/gocheck/pkg/errutil"
	"github.com/Hidayathamir/gocheck/pkg/gocheckhttp"
	"github.com/Hidayathamir/gocheck/pkg/h"
	"github.com/sirupsen/logrus"
)

//nolint:gomnd
func main() {
	req := gocheckhttp.ReqDigitalWalletTransfer{
		RecipientID: 2,
		Amount:      10000,
	}

	jsonByte, err := json.Marshal(req)
	fatalIfErr(err)

	ctx := context.Background()
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:10010/api/v1/digital-wallet/transfer", bytes.NewBuffer(jsonByte))
	fatalIfErr(err)

	httpReq.Header.Add(h.ContentType, h.APPJSON)

	auth := gocheckhttp.Authorization{UserID: 1}
	jsonByte, err = json.Marshal(auth)
	fatalIfErr(err)

	httpReq.Header.Add(h.Authorization, string(jsonByte))

	httpClient := &http.Client{}
	httpRes, err := httpClient.Do(httpReq)
	fatalIfErr(err)
	defer func() {
		err := httpRes.Body.Close()
		warnIfErr(err)
	}()

	body, err := io.ReadAll(httpRes.Body)
	fatalIfErr(err)

	resBody := gocheckhttp.ResDigitalWalletTransfer{}
	err = json.Unmarshal(body, &resBody)
	fatalIfErr(err)

	isStatusCode2xx := string(httpRes.Status[0]) == "2"
	if !isStatusCode2xx || resBody.Error != "" {
		err := errors.New(resBody.Error)
		logrus.Fatal(errutil.Wrap(err))
	}

	logrus.Info("transfer id = ", resBody.Data.ID)
}

func fatalIfErr(err error) {
	if err != nil {
		logrus.Fatal(errutil.Wrap(err, errutil.WithSkip(1)))
	}
}

func warnIfErr(err error) {
	if err != nil {
		logrus.Warn(errutil.Wrap(err, errutil.WithSkip(1)))
	}
}
