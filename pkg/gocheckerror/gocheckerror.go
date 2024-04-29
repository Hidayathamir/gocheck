// Package gocheckerror provides custom error definitions.
package gocheckerror

import "errors"

var (
	// ErrUnauthenticated represents an unauthenticated error.
	ErrUnauthenticated = errors.New("unauthenticated")

	// ErrInvalidRequest represents an invalid request error.
	ErrInvalidRequest = errors.New("invalid request")
)
