//go:build !change

package memoize

// Func computes a value together with an optional error.
type Func func() (int, error)
