//go:build !change

package bankaccount

import "errors"

// Money is an amount in the bank's smallest currency unit.
type Money int64

// Account is a bank account. Its balance can only be changed through package functions.
type Account struct {
	balance Money
}

var (
	ErrNilAccount        = errors.New("nil account")
	ErrInvalidAmount     = errors.New("invalid amount")
	ErrInsufficientFunds = errors.New("insufficient funds")
)
