package gocheckhttpmiddleware

import (
	"encoding/json"
	"fmt"

	"github.com/Hidayathamir/gocheck/pkg/gocheck"
	"github.com/Hidayathamir/gocheck/pkg/gocheckerror"
	"github.com/Hidayathamir/gocheck/pkg/h"
	"github.com/Hidayathamir/gocheck/pkg/trace"
	"github.com/gin-gonic/gin"
)

// GetAuthFromGinCtxHeader -.
func GetAuthFromGinCtxHeader(c *gin.Context) (gocheck.Authorization, error) {
	auth := gocheck.Authorization{}
	err := json.Unmarshal([]byte(c.GetHeader(h.Authorization)), &auth)
	if err != nil {
		err = fmt.Errorf("%w: %w", gocheckerror.ErrUnauthenticated, err)
		return gocheck.Authorization{}, trace.Wrap(err)
	}

	return auth, nil
}
