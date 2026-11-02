// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package memo demonstrates a cache with one lock around lookup and computation.
// Construct it with New and do not copy it. The loader must be non-nil, must not
// panic, and must not call Get on this Memo, even for a different key.
// All calls to the loader are serialized, including calls for different keys.
// Both values and errors are cached. Published results must remain immutable.
package memo

import "sync"

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
	mu    sync.Mutex // protects cache
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
// Get holds one lock during lookup and computation.
func (memo *Memo) Get(key string) (any, error) {
	memo.mu.Lock()
	defer memo.mu.Unlock()

	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}

	return res.value, res.err
}

// END GET OMIT
