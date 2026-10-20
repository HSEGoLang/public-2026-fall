package moreints

import "github.com/google/go-cmp/cmp"

func CheckReverse(input, want []int) string {
	return cmp.Diff(want, Reverse(input))
}
