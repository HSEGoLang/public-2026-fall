// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package memo demonstrates an unsafe cache: use each Memo from one goroutine only.
// Construct it with New; its zero value is not ready for use. The loader must be
// non-nil, must not panic, and must not recursively request the same key.
// Both values and errors are cached. Published results must remain immutable.
package memo

// START VALUE OMIT
// Func computes a value for a key.
type Func func(key string) (any, error)

type result struct {
	value any
	err   error
}

// END VALUE OMIT

// START TYPES OMIT
// Memo caches calls to a Func. Construct it with New.
type Memo struct {
	f     Func
	cache map[string]result
}

// END TYPES OMIT

// START NEW OMIT
// New creates an empty cache for f.
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// END NEW OMIT

// START GET OMIT
// Get returns a cached result. It is not concurrency-safe.
func (memo *Memo) Get(key string) (any, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}

	return res.value, res.err
}

// END GET OMIT
