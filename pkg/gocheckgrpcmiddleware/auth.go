package gocheckgrpcmiddleware

import (
	"context"
	"encoding/json"

	"github.com/Hidayathamir/gocheck/pkg/m"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"google.golang.org/grpc/metadata"
)

// Authorization -.
type Authorization struct {
	UserID uint `json:"user_id"`
}

// SetAuthToCtx -.
func SetAuthToCtx(ctx context.Context, auth Authorization) (context.Context, error) {
	jsonByte, err := json.Marshal(auth)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(m.Authorization, string(jsonByte)))

	return ctx, nil
}
