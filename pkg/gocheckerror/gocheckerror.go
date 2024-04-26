// Package gocheckerror -.
package gocheckerror

import "errors"

var (
	// ErrUnauthenticated -.
	ErrUnauthenticated = errors.New("unauthenticated")

	// ErrInvalidRequest -.
	ErrInvalidRequest = errors.New("invalid request")
)
