package main

import "sync"

// START ENTRY_SKETCH OMIT
// Entry is a sketch of a shared result slot, not a ready cache implementation.
type Entry struct {
	Once  sync.Once
	Value any
	Err   error
}

// END ENTRY_SKETCH OMIT
