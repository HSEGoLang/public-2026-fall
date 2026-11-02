// This program intentionally races. Run with -race to see the diagnostic.
package main

import (
	"fmt"
	"sync"

	"hse-go/lectures/08-concurrency/pitfalls"
)

// START RACE_MAIN OMIT
func main() {
	var (
		a  pitfalls.Account
		wg sync.WaitGroup
	)

	wg.Add(2)
	go func() { defer wg.Done(); a.Deposit(200) }()
	go func() { defer wg.Done(); a.Deposit(100) }()

	wg.Wait()
	fmt.Println(a.Balance)
}

// END RACE_MAIN OMIT
