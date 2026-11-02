package main

import (
	"fmt"
	"strings"
	"sync"

	memo4 "hse-go/lectures/08-concurrency/memo4"
)

// START CONCURRENT OMIT
// START CONCURRENTSETUP OMIT
func main() {
	var mu sync.Mutex

	calls := 0

	memo := memo4.New(func(key string) (any, error) {
		mu.Lock()
		calls++
		mu.Unlock()

		return strings.ToUpper(key), nil
	})

	var values [2]any

	var wg sync.WaitGroup
	// END CONCURRENTSETUP OMIT
	// START CONCURRENTRUN OMIT
	for i := range values {
		wg.Add(1)
		go func() {
			defer wg.Done()

			values[i], _ = memo.Get("tea")
		}()
	}

	wg.Wait()
	fmt.Println("один ключ:", values, "вычислений:", calls)
}

// END CONCURRENTRUN OMIT
// END CONCURRENT OMIT
