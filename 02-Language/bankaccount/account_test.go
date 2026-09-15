package bankaccount_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"hse-go/02-Language/bankaccount"
)

func newAccount(t *testing.T, initial bankaccount.Money) *bankaccount.Account {
	t.Helper()

	account, err := bankaccount.NewAccount(initial)
	require.NoError(t, err)

	return account
}

func balance(t *testing.T, account *bankaccount.Account) bankaccount.Money {
	t.Helper()

	value, err := bankaccount.Balance(account)
	require.NoError(t, err)

	return value
}

func TestNewAccount(t *testing.T) {
	account := newAccount(t, 125)
	require.Equal(t, bankaccount.Money(125), balance(t, account))

	_, err := bankaccount.NewAccount(-1)
	require.ErrorIs(t, err, bankaccount.ErrInvalidAmount)
}

func TestDepositAndWithdraw(t *testing.T) {
	account := newAccount(t, 100)

	require.NoError(t, bankaccount.Deposit(account, 25))
	require.Equal(t, bankaccount.Money(125), balance(t, account))

	require.NoError(t, bankaccount.Withdraw(account, 40))
	require.Equal(t, bankaccount.Money(85), balance(t, account))
}

func TestFailedWithdrawDoesNotChangeBalance(t *testing.T) {
	account := newAccount(t, 10)

	err := bankaccount.Withdraw(account, 11)
	require.ErrorIs(t, err, bankaccount.ErrInsufficientFunds)
	require.Equal(t, bankaccount.Money(10), balance(t, account))
}

func TestTransfer(t *testing.T) {
	from := newAccount(t, 100)
	to := newAccount(t, 20)

	require.NoError(t, bankaccount.Transfer(from, to, 35))
	require.Equal(t, bankaccount.Money(65), balance(t, from))
	require.Equal(t, bankaccount.Money(55), balance(t, to))
}

func TestFailedTransferChangesNeitherAccount(t *testing.T) {
	from := newAccount(t, 10)
	to := newAccount(t, 20)

	err := bankaccount.Transfer(from, to, 11)
	require.ErrorIs(t, err, bankaccount.ErrInsufficientFunds)
	require.Equal(t, bankaccount.Money(10), balance(t, from))
	require.Equal(t, bankaccount.Money(20), balance(t, to))
}
