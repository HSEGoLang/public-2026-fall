package linefilter

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestFiltering(t *testing.T) {
	var (
		out  bytes.Buffer
		seen []string
	)

	n, err := Filter(&out, strings.NewReader("ok\n\nerror: диск"), func(s string) bool {
		seen = append(seen, s)
		return strings.HasPrefix(s, "error:")
	})
	if err != nil || n != 1 || out.String() != "error: диск\n" {
		t.Fatalf("%d %v %q", n, err, out.String())
	}

	if !reflect.DeepEqual(seen, []string{"ok", "", "error: диск"}) {
		t.Fatal(seen)
	}
}

func TestLineBoundaries(t *testing.T) {
	for _, tc := range []struct {
		in, out string
		n       int
	}{
		{"", "", 0},
		{"\n", "\n", 1},
		{"a\n", "a\n", 1},
		{"a", "a\n", 1},
		{"a\n\n", "a\n\n", 2},
		{"a\r\n", "a\r\n", 1},
	} {
		var out bytes.Buffer

		n, err := Filter(&out, strings.NewReader(tc.in), func(string) bool { return true })
		if err != nil || n != tc.n || out.String() != tc.out {
			t.Fatalf("input %q: %d %v %q", tc.in, n, err, out.String())
		}
	}
}
