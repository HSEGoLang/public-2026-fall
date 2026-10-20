package greeting_test

import (
	"testing"

	greeting "hse-go/lectures/06-testing/externaltest"
)

func TestNormalizeName(t *testing.T) {
	if got := greeting.NormalizeName("  Alice  "); got != "Alice" {
		t.Errorf("NormalizeName = %q, want Alice", got)
	}
}

func TestHello(t *testing.T) {
	if got := greeting.Hello("  Alice  "); got != "Hello, Alice!" {
		t.Errorf("Hello = %q, want Hello, Alice!", got)
	}
}
