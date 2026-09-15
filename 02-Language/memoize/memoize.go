//go:build !solution

package memoize

// Memoize returns a lazy function that caches the result of source.
func Memoize(source Func) Func {
	return source
}
