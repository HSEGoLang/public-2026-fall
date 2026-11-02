// Package synconce combines sync.Map and sync.Once to cache one result per key.
// Construct Cache with New and do not copy it. The loader must be non-nil,
// concurrency-safe, non-panicking, and must not recursively call this Cache.
// Different keys may be computed concurrently. Both values and errors are
// cached without eviction, and published results must remain immutable.
package synconce

import "sync"

// Func computes a value for a key.
type Func func(key string) (any, error)

// START TYPES OMIT
// Cache owns entries for one loader. Construct it with New.
type Cache struct {
	compute Func
	entries sync.Map // string -> *entry
}

type entry struct {
	once  sync.Once
	value any
	err   error
}

// END TYPES OMIT

// START NEW OMIT
// New creates an empty cache for compute.
func New(compute Func) *Cache {
	return &Cache{compute: compute}
}

// END NEW OMIT

// START GET OMIT
// Get returns or computes the cached result for key.
func (c *Cache) Get(key string) (any, error) {
	candidate := &entry{}
	stored, _ := c.entries.LoadOrStore(key, candidate)
	e := stored.(*entry)
	e.once.Do(func() {
		e.value, e.err = c.compute(key)
	})

	return e.value, e.err
}

// END GET OMIT
