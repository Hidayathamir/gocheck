// Package gocheckhttp -.
package gocheckhttp

// HTTPResponse -.
type HTTPResponse[Body any] struct {
	Body       Body
	StatusCode int
}
