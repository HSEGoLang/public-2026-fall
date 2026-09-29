//go:build !change

package retryread

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrDenied   = errors.New("access denied")
	ErrAttempts = errors.New("attempts must be positive")
)

type Client interface {
	Get(key string) (string, error)
}
type Temporary interface{ IsTemporary() bool }

// AttemptError describes a failed call sequence; Err is the last cause.
type AttemptError struct {
	Key      string
	Attempts int
	Err      error
}
