// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package memo demonstrates a cache with one shared result entry per key.
// Construct it with New and do not copy it. The loader must be non-nil,
// concurrency-safe, non-panicking, and must not recursively call this Memo.
// Requests for one key share a computation; different keys can be computed
// concurrently. Both values and errors are cached without eviction.
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
	cache map[string]*entry
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// END TYPES OMIT

// START NEW OMIT
// New creates an empty cache for f.
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// END NEW OMIT

// START GET OMIT
// Get shares a result with other callers for the same key.
func (memo *Memo) Get(key string) (any, error) {
	memo.mu.Lock()

	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}

	return e.res.value, e.res.err
}

// END GET OMIT
