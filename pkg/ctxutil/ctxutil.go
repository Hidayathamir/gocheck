package ctxutil

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const ctxKeyTraceID = "github.com/Hidayathamir/goout#ctxKeyTraceID"

// SetTraceIDToCtx set trace id into context.
func SetTraceIDToCtx(ctx context.Context, traceID string) context.Context {
	ctx = context.WithValue(ctx, ctxKeyTraceID, traceID) //nolint:revive,staticcheck
	return ctx
}

// SetTraceIDToGinCtx set trace id into gin context.
func SetTraceIDToGinCtx(c *gin.Context, traceID string) {
	c.Set(string(ctxKeyTraceID), traceID)
}

// GetTraceIDFromCtx return trace id from context.
func GetTraceIDFromCtx(ctx context.Context) string {
	traceIDAny := ctx.Value(ctxKeyTraceID)
	if traceIDAny == nil {
		logrus.Warn("context does not have trace id")
		return "context does not have trace id"
	}

	traceID, ok := traceIDAny.(string)
	if !ok {
		logrus.Warn("trace id in context is not string")
		return "trace id in context is not string"
	}

	return traceID
}
