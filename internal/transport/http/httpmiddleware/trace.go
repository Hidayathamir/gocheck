package httpmiddleware

import (
	"github.com/Hidayathamir/gocheck/pkg/ctxutil"
	"github.com/Hidayathamir/gocheck/pkg/h"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TraceID set trace id to gin context. Will check if header contain trace id,
// if yes then will use trace id from header,
// if no then will generate new trace id.
func TraceID(c *gin.Context) {
	traceID := c.GetHeader(h.XTraceID)
	if traceID == "" {
		traceID = uuid.NewString()
	}
	ctxutil.SetTraceIDToGinCtx(c, traceID)
}
