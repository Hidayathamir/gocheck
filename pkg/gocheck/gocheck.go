// Package gocheck -.
package gocheck

import "time"

// DefaultCacheExpire -.
const DefaultCacheExpire = 24 * time.Hour

// MinimumTransferAmount -.
const MinimumTransferAmount = 10000

// Authorization -.
type Authorization struct {
	UserID uint `json:"user_id"`
}
