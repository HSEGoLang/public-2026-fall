package fuzzreverse

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	for _, s := range []string{"hello", "", "123"} {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		if !utf8.ValidString(s) {
			t.Skip()
		}

		rev := Reverse(s)
		if !utf8.ValidString(rev) {
			t.Fatalf("invalid UTF-8: Reverse(%q) = %q", s, rev)
		}

		if got := Reverse(rev); got != s {
			t.Fatalf("double reverse = %q, want %q", got, s)
		}
	})
}
