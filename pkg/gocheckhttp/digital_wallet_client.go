package gocheckhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Hidayathamir/gocheck/pkg/h"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/sirupsen/logrus"
)

// DigitalWalletClient -.
type DigitalWalletClient struct {
	Base        string
	URLTransfer string
}

// NewDigitalWalletClient -.
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

// Transfer -.
func (d *DigitalWalletClient) Transfer(ctx context.Context, req ReqDigitalWalletTransfer) (HTTPResponse[ResDigitalWalletTransfer], error) {
	fail := func(err error) (HTTPResponse[ResDigitalWalletTransfer], error) {
		return HTTPResponse[ResDigitalWalletTransfer]{}, trace.Wrap(err, trace.WithSkip(1))
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

	res := HTTPResponse[ResDigitalWalletTransfer]{
		Body:       resBody,
		StatusCode: httpRes.StatusCode,
	}

	return res, nil
}
