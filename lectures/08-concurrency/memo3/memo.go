// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package memo demonstrates a cache which allows duplicate computations.
// Construct it with New and do not copy it. The loader must be non-nil,
// concurrency-safe, non-panicking, and must not recursively request the same key.
// Concurrent misses may compute the same key more than once; different keys can
// also be computed concurrently. Each caller returns its own computed result;
// the last stored result is kept for future hits. Values and errors are cached.
// Published results must remain immutable.
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
// Get may duplicate computations on concurrent misses.
func (memo *Memo) Get(key string) (any, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()

	if !ok {
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}

	return res.value, res.err
}

// END GET OMIT
