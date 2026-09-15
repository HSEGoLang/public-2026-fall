//go:build !solution

package expr

// Eval evaluates an arithmetic expression tree.
func Eval(expression *Expression) (int64, error) {
	_ = expression

	return 0, nil
}
