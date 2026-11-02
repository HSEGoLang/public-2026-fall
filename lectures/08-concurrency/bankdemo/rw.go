package main

import "sync"

// START RW_BANK OMIT
type AccountRW struct {
	mu      sync.RWMutex
	balance int
}

func (a *AccountRW) Balance() int {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.balance
}

// END RW_BANK OMIT

func (a *AccountRW) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.balance += amount
}
