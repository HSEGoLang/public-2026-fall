package main

import (
	"fmt"
	"sync/atomic"
)

// START CONFIG_TYPE OMIT
type Config struct {
	Greeting string
}

// END CONFIG_TYPE OMIT

// START CONFIG_SNAPSHOT OMIT
func main() {
	var current atomic.Value
	current.Store(Config{Greeting: "hello"})
	old := current.Load().(Config)

	current.Store(Config{Greeting: "hi"})
	newest := current.Load().(Config)

	fmt.Println("old:", old.Greeting)
	fmt.Println("new:", newest.Greeting)
}

// END CONFIG_SNAPSHOT OMIT
