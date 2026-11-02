package main

import (
	"fmt"
	"strings"
	"sync"
)

// START MAP_CANDIDATES OMIT
func main() {
	var cache sync.Map

	calls := 0
	calculate := func(key string) string {
		calls++
		return strings.ToUpper(key)
	}
	first, _ := cache.LoadOrStore("tea", calculate("tea"))
	second, _ := cache.LoadOrStore("tea", calculate("tea"))
	fmt.Println(first, second, "calls:", calls)
}

// END MAP_CANDIDATES OMIT
