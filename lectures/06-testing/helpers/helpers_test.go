package helpers

import (
	"fmt"
	"testing"
)

type testingT interface {
	Helper()
	Errorf(string, ...any)
}

func checkPositive(t testingT, n int) {
	t.Helper()

	if n <= 0 {
		t.Errorf("got %d, want positive", n)
	}
}

type fakeT struct {
	message string
}

func (*fakeT) Helper() {}

func (f *fakeT) Errorf(format string, args ...any) {
	f.message = fmt.Sprintf(format, args...)
}

func TestHelperFailure(t *testing.T) {
	var fake fakeT
	checkPositive(&fake, -1)

	if fake.message != "got -1, want positive" {
		t.Errorf("message = %q", fake.message)
	}
}
