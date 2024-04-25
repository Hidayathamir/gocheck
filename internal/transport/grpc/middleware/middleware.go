// Package middleware -.
package middleware

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Hidayathamir/gocheck/pkg/gocheck"
	"github.com/Hidayathamir/gocheck/pkg/gocheckerror"
	"github.com/Hidayathamir/gocheck/pkg/m"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// GetAuthFromCtx -.
func GetAuthFromCtx(ctx context.Context) (gocheck.Authorization, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err := status.Errorf(codes.Unauthenticated, "missing metadata")
		err = fmt.Errorf("%w: %w", gocheckerror.ErrUnauthenticated, err)
		return gocheck.Authorization{}, trace.Wrap(err)
	}

	mdAuth := md.Get(m.Authorization)
	if len(mdAuth) == 0 {
		err := status.Errorf(codes.Unauthenticated, "missing token")
		err = fmt.Errorf("%w: %w", gocheckerror.ErrUnauthenticated, err)
		return gocheck.Authorization{}, trace.Wrap(err)
	}

	auth := gocheck.Authorization{}
	err := json.Unmarshal([]byte(mdAuth[0]), &auth)
	if err != nil {
		err = fmt.Errorf("%w: %w", gocheckerror.ErrUnauthenticated, err)
		return gocheck.Authorization{}, trace.Wrap(err)
	}

	return auth, nil
}
