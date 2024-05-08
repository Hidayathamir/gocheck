package httpmiddleware

import (
	"encoding/json"
	"fmt"

	"github.com/Hidayathamir/gocheck/pkg/gocheckerror"
	"github.com/Hidayathamir/gocheck/pkg/gocheckhttp"
	"github.com/Hidayathamir/gocheck/pkg/h"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/gin-gonic/gin"
)

// GetAuthFromGinCtxHeader extracts authorization information from the Gin context header.
func GetAuthFromGinCtxHeader(c *gin.Context) (gocheckhttp.Authorization, error) {
	auth := gocheckhttp.Authorization{}
	err := json.Unmarshal([]byte(c.GetHeader(h.Authorization)), &auth)
	if err != nil {
		err = fmt.Errorf("%w: %w", gocheckerror.ErrUnauthenticated, err)
		return gocheckhttp.Authorization{}, trace.Wrap(err)
	}

	return auth, nil
}
