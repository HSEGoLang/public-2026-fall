package main

import (
	"fmt"
	"strings"

	memo1 "hse-go/lectures/08-concurrency/memo1"
)

// START SEQUENTIAL OMIT
func main() {
	calls := 0

	memo := memo1.New(func(key string) (any, error) {
		calls++
		return strings.ToUpper(key), nil
	})

	for _, key := range []string{"tea", "tea", "coffee"} {
		value, _ := memo.Get(key) // this loader always returns a nil error
		fmt.Println(key, "->", value)
	}

	fmt.Println("вычислений:", calls)
}

// END SEQUENTIAL OMIT
