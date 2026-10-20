// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Test the echo command with: go test ./06-testing/echo

// !+
package main

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestEchoBuffer(t *testing.T) {
	saved := out

	t.Cleanup(func() { out = saved })

	var buf bytes.Buffer

	out = &buf

	if err := echo(true, " ", []string{"Go", "tests"}); err != nil {
		t.Fatal(err)
	}

	if got := buf.String(); got != "Go tests\n" {
		t.Errorf("echo = %q, want %q", got, "Go tests\n")
	}
}

func TestEcho(t *testing.T) {
	saved := out

	t.Cleanup(func() { out = saved })

	tests := []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("echo(%v, %q, %q)", test.newline, test.sep, test.args)
		out = new(bytes.Buffer) // captured output

		if err := echo(test.newline, test.sep, test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

type failingWriter struct {
	calls  int
	failAt int
	err    error
}

func (w *failingWriter) Write(p []byte) (int, error) {
	w.calls++
	if w.calls == w.failAt {
		return 0, w.err
	}

	return len(p), nil
}

func TestEchoWriteError(t *testing.T) {
	saved := out

	t.Cleanup(func() { out = saved })

	for _, failAt := range []int{1, 2} {
		want := errors.New("disk full")

		out = &failingWriter{failAt: failAt, err: want}
		if err := echo(true, " ", []string{"hello"}); !errors.Is(err, want) {
			t.Fatalf("write %d: got %v, want %v", failAt, err, want)
		}
	}
}
