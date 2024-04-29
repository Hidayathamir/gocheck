package gocheckhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Hidayathamir/gocheck/pkg/gocheckhttpmiddleware"
	"github.com/Hidayathamir/gocheck/pkg/h"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/sirupsen/logrus"
)

// DigitalWalletClient represents a client for interacting with the DigitalWallet API.
type DigitalWalletClient struct {
	Base        string
	URLTransfer string
}

// NewDigitalWalletClient creates a new instance of DigitalWalletClient.
func NewDigitalWalletClient(base string) *DigitalWalletClient {
	return &DigitalWalletClient{
		Base:        base,
		URLTransfer: "/api/v1/digital-wallet/transfer",
	}
}

////////////////////////////////////////

func (d *DigitalWalletClient) getURLTransfer() string {
	return d.Base + d.URLTransfer
}

////////////////////////////////////////

// Transfer sends http request to create transfer.
func (d *DigitalWalletClient) Transfer(ctx context.Context, auth gocheckhttpmiddleware.Authorization, req ReqDigitalWalletTransfer) (ResDataDigitalWalletTransfer, error) {
	fail := func(err error) (ResDataDigitalWalletTransfer, error) {
		return ResDataDigitalWalletTransfer{}, trace.Wrap(err, trace.WithSkip(1))
	}

	jsonByte, err := json.Marshal(req)
	if err != nil {
		return fail(err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, d.getURLTransfer(), bytes.NewBuffer(jsonByte))
	if err != nil {
		return fail(err)
	}

	httpReq.Header.Add(h.ContentType, h.APPJSON)

	jsonByte, err = json.Marshal(auth)
	if err != nil {
		return fail(err)
	}

	httpReq.Header.Add(h.Authorization, string(jsonByte))

	httpClient := &http.Client{}
	httpRes, err := httpClient.Do(httpReq)
	if err != nil {
		return fail(err)
	}
	defer func() {
		err := httpRes.Body.Close()
		if err != nil {
			logrus.Warn(trace.Wrap(err))
		}
	}()

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return fail(err)
	}

	resBody := ResDigitalWalletTransfer{}
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		return fail(err)
	}

	isStatusCode2xx := string(httpRes.Status[0]) == "2"
	if !isStatusCode2xx || resBody.Error != "" {
		err := errors.New(resBody.Error)
		return fail(err)
	}

	return resBody.Data, nil
}
