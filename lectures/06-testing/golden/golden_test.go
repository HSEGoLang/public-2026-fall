package golden

import (
	"bytes"
	"os"
	"testing"
)

func TestGolden(t *testing.T) {
	want, err := os.ReadFile("testdata/hello.golden")
	if err != nil {
		t.Fatal(err)
	}

	got := []byte("Hello, world!\n")
	if !bytes.Equal(got, want) {
		t.Errorf("output = %q, want %q", got, want)
	}
}
