//go:build !solution

package bankaccount

// NewAccount creates an account with the supplied initial balance.
func NewAccount(initial Money) (*Account, error) {
	return &Account{balance: initial}, nil
}

// Balance returns the current account balance.
func Balance(account *Account) (Money, error) {
	_ = account

	return 0, nil
}

// Deposit adds money to the account.
func Deposit(account *Account, amount Money) error {
	_, _ = account, amount

	return nil
}

// Withdraw removes money from the account.
func Withdraw(account *Account, amount Money) error {
	_, _ = account, amount

	return nil
}

// Transfer moves money from one account to another.
func Transfer(from, to *Account, amount Money) error {
	_, _, _ = from, to, amount

	return nil
}
