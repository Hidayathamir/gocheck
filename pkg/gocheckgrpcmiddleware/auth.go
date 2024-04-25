package gocheckgrpcmiddleware

import (
	"context"
	"encoding/json"

	"github.com/Hidayathamir/gocheck/pkg/gocheck"
	"github.com/Hidayathamir/gocheck/pkg/m"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"google.golang.org/grpc/metadata"
)

// SetAuthToCtx -.
func SetAuthToCtx(ctx context.Context, auth gocheck.Authorization) (context.Context, error) {
	jsonByte, err := json.Marshal(auth)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs(m.Authorization, string(jsonByte)))

	return ctx, nil
}
