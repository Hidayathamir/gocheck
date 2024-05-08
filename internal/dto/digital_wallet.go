package dto

import (
	"context"

	"github.com/Hidayathamir/gocheck/internal/transport/grpc/grpcmiddleware"
	"github.com/Hidayathamir/gocheck/internal/transport/http/httpmiddleware"
	"github.com/Hidayathamir/gocheck/pkg/gocheckhttp"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	pbgocheck "github.com/Hidayathamir/protobuf/gocheck"
	"github.com/gin-gonic/gin"
)

// ReqDigitalWalletTransfer represents the request data structure for transfer.
type ReqDigitalWalletTransfer struct {
	SenderID    uint `validate:"required,nefield=RecipientID"`
	RecipientID uint `validate:"required"`
	Amount      int  `validate:"required"`
}

// LoadFromReqGRPC laods data from request grpc.
func (r *ReqDigitalWalletTransfer) LoadFromReqGRPC(ctx context.Context, req *pbgocheck.ReqDigitalWalletTransfer) error {
	auth, err := grpcmiddleware.GetAuthFromCtx(ctx)
	if err != nil {
		return trace.Wrap(err)
	}

	r.SenderID = auth.UserID
	r.RecipientID = uint(req.GetRecipientId())
	r.Amount = int(req.GetAmount())

	return nil
}

// LoadFromReqHTTP laods data from request http.
func (r *ReqDigitalWalletTransfer) LoadFromReqHTTP(c *gin.Context) error {
	auth, err := httpmiddleware.GetAuthFromGinCtxHeader(c)
	if err != nil {
		return trace.Wrap(err)
	}

	req := gocheckhttp.ReqDigitalWalletTransfer{}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		return trace.Wrap(err)
	}

	r.SenderID = auth.UserID
	r.RecipientID = req.RecipientID
	r.Amount = req.Amount

	return nil
}

// ResDigitalWalletTransfer represents the response data structure for transfer.
type ResDigitalWalletTransfer struct {
	ID uint
}

// ToResGRPC converts response to gRPC format.
func (r *ResDigitalWalletTransfer) ToResGRPC() *pbgocheck.ResDigitalWalletTransfer {
	return &pbgocheck.ResDigitalWalletTransfer{
		Id: uint64(r.ID),
	}
}

// ToResHTTP converts response to HTTP format.
func (r *ResDigitalWalletTransfer) ToResHTTP() gocheckhttp.ResDigitalWalletTransfer {
	return gocheckhttp.ResDigitalWalletTransfer{
		Data: gocheckhttp.ResDataDigitalWalletTransfer{
			ID: r.ID,
		},
	}
}
