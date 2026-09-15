package expr_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"hse-go/02-Language/expr"
)

func value(number int64) *expr.Expression {
	return &expr.Expression{Value: number}
}

func operation(kind expr.Operation, left, right *expr.Expression) *expr.Expression {
	return &expr.Expression{Operation: kind, Left: left, Right: right}
}

func TestEvalValuesAndOperations(t *testing.T) {
	testCases := []struct {
		name       string
		expression *expr.Expression
		want       int64
	}{
		{name: "zero value", expression: &expr.Expression{}, want: 0},
		{name: "negative value", expression: value(-17), want: -17},
		{name: "add", expression: operation(expr.OperationAdd, value(7), value(5)), want: 12},
		{
			name:       "subtract keeps operand order",
			expression: operation(expr.OperationSubtract, value(3), value(8)),
			want:       -5,
		},
		{
			name:       "integer division",
			expression: operation(expr.OperationDivide, value(8), value(3)),
			want:       2,
		},
		{
			name: "nested",
			expression: operation(
				expr.OperationMultiply,
				operation(expr.OperationAdd, value(7), value(5)),
				operation(expr.OperationSubtract, value(10), value(4)),
			),
			want: 72,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := expr.Eval(testCase.expression)
			require.NoError(t, err)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestEvalErrors(t *testing.T) {
	testCases := []struct {
		name       string
		expression *expr.Expression
		want       error
	}{
		{name: "nil root", expression: nil, want: expr.ErrInvalidExpression},
		{
			name:       "missing operand",
			expression: operation(expr.OperationAdd, value(1), nil),
			want:       expr.ErrInvalidExpression,
		},
		{
			name:       "unknown operation",
			expression: operation(expr.Operation(255), value(1), value(2)),
			want:       expr.ErrInvalidExpression,
		},
		{
			name:       "division by zero",
			expression: operation(expr.OperationDivide, value(10), value(0)),
			want:       expr.ErrDivisionByZero,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := expr.Eval(testCase.expression)
			require.ErrorIs(t, err, testCase.want)
		})
	}
}

func TestEvalChecksOperationBeforeChildren(t *testing.T) {
	expression := operation(
		expr.Operation(255),
		operation(expr.OperationDivide, value(1), value(0)),
		value(2),
	)

	_, err := expr.Eval(expression)
	require.ErrorIs(t, err, expr.ErrInvalidExpression)
	require.NotErrorIs(t, err, expr.ErrDivisionByZero)
}
