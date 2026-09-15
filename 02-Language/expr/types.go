//go:build !change

package expr

import "errors"

// Operation identifies the kind of an expression node.
type Operation uint8

const (
	OperationValue Operation = iota
	OperationAdd
	OperationSubtract
	OperationMultiply
	OperationDivide
)

// Expression is a node of an arithmetic expression tree.
type Expression struct {
	Operation Operation
	Value     int64
	Left      *Expression
	Right     *Expression
}

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrDivisionByZero    = errors.New("division by zero")
)
