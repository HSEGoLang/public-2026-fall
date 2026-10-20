package morestrings

import "github.com/google/go-cmp/cmp"

func CheckReverse(input, want string) string {
	return cmp.Diff(want, ReverseRunes(input))
}
