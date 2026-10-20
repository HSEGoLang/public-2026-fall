package benchcases

import (
	"strings"
	"testing"
)

func repeatConcat(s string, n int) string {
	result := ""
	for range n {
		result += s
	}

	return result
}

func repeatBuiltin(s string, n int) string {
	return strings.Repeat(s, n)
}

func TestRepeat(t *testing.T) {
	for _, tc := range []struct {
		input string
		n     int
		want  string
	}{
		{input: "go", n: 3, want: "gogogo"},
		{input: "go", n: 0, want: ""},
		{input: "", n: 3, want: ""},
	} {
		if got := repeatConcat(tc.input, tc.n); got != tc.want {
			t.Errorf("repeatConcat(%q, %d) = %q, want %q", tc.input, tc.n, got, tc.want)
		}

		if got := repeatBuiltin(tc.input, tc.n); got != tc.want {
			t.Errorf("repeatBuiltin(%q, %d) = %q, want %q", tc.input, tc.n, got, tc.want)
		}
	}
}

func BenchmarkRepeatConcat(b *testing.B) {
	for b.Loop() {
		repeatConcat("go", 100)
	}
}

func BenchmarkRepeatBuiltin(b *testing.B) {
	for b.Loop() {
		repeatBuiltin("go", 100)
	}
}
