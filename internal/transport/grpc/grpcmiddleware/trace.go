package grpcmiddleware

import (
	"context"

	"github.com/Hidayathamir/gocheck/pkg/ctxutil"
	"github.com/Hidayathamir/gocheck/pkg/m"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// TraceID set trace id to context. Will check if metadata contain trace id,
// if yes then will use trace id from metadata,
// if no then will generate new trace id.
func TraceID(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	traceID := getTraceID(ctx)
	ctx = ctxutil.SetTraceIDToCtx(ctx, traceID)

	res, err := handler(ctx, req)

	return res, err
}

func getTraceID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uuid.NewString()
	}

	mdTraceID := md.Get(m.TraceID)
	if len(mdTraceID) == 0 {
		return uuid.NewString()
	}

	return mdTraceID[0]
}
