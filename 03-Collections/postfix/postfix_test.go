package postfix_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"hse-go/03-Collections/postfix"
)

func TestEval(t *testing.T) {
	testCases := []struct {
		name       string
		expression string
		want       int
	}{
		{name: "number", expression: "42", want: 42},
		{name: "add", expression: "2 3 +", want: 5},
		{name: "subtract order", expression: "3 8 -", want: -5},
		{name: "integer division", expression: "8 3 /", want: 2},
		{name: "negative division rounds toward zero", expression: "-7 3 /", want: -2},
		{name: "nested", expression: "10 2 3 * -", want: 4},
		{name: "whitespace and negative number", expression: "\t-2\n 5 * ", want: -10},
		{name: "Unicode whitespace", expression: "2\u00a03\u2003+", want: 5},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := postfix.Eval(testCase.expression)
			require.NoError(t, err)
			require.Equal(t, testCase.want, got)
		})
	}
}

func TestEvalErrors(t *testing.T) {
	testCases := []struct {
		name       string
		expression string
		want       error
	}{
		{name: "empty", expression: "", want: postfix.ErrInvalidExpression},
		{name: "not enough operands", expression: "1 +", want: postfix.ErrInvalidExpression},
		{name: "extra operands", expression: "1 2", want: postfix.ErrInvalidExpression},
		{name: "unknown token", expression: "1 nope", want: postfix.ErrUnknownToken},
		{name: "division by zero", expression: "4 0 /", want: postfix.ErrDivisionByZero},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := postfix.Eval(testCase.expression)
			require.ErrorIs(t, err, testCase.want)
		})
	}
}
