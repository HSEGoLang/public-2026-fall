//go:build !change

package postfix

import "errors"

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrUnknownToken      = errors.New("unknown token")
	ErrDivisionByZero    = errors.New("division by zero")
)
