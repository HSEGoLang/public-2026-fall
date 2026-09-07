//go:build !solution
// +build !solution

package sum

import "hse-go/sum/pkg"

func Sum(a, b int64) int64 {
	pkg.F()
	return a + b
}
