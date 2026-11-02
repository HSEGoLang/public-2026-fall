package main

import (
	"fmt"
	"sync"
)

// START BANK OMIT
// START BANK_WRITE OMIT
type Account struct {
	mu      sync.Mutex
	balance int
}

func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.balance += amount
}

// END BANK_WRITE OMIT

func (a *Account) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.balance
}

// END BANK OMIT

// START BANK_MAIN OMIT
func main() {
	var (
		a  Account
		wg sync.WaitGroup
	)

	wg.Add(2)
	go func() { defer wg.Done(); a.Deposit(200) }()
	go func() { defer wg.Done(); a.Deposit(100) }()

	wg.Wait()
	fmt.Println(a.Balance())
}

// END BANK_MAIN OMIT
