package runereverse

import (
	"testing"
	"unicode/utf8"
)

func TestRegression(t *testing.T) {
	if got := Reverse("éa"); got != "aé" {
		t.Errorf("Reverse = %q, want aé", got)
	}
}

func FuzzReverse(f *testing.F) {
	for _, s := range []string{"hello", "", "été"} {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		if !utf8.ValidString(s) {
			t.Skip()
		}

		rev := Reverse(s)
		if !utf8.ValidString(rev) || Reverse(rev) != s {
			t.Fatalf("reverse invariant failed for %q", s)
		}
	})
}
