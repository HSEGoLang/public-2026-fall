package fixtures

import (
	"os"
	"testing"
)

func newFile(t *testing.T) string {
	t.Helper()

	f, err := os.CreateTemp(t.TempDir(), "example-")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := f.Close(); err != nil {
			t.Error(err)
		}
	})

	return f.Name()
}

func TestFile(t *testing.T) {
	name := newFile(t)
	if _, err := os.Stat(name); err != nil {
		t.Fatal(err)
	}
}

type Pair struct {
	Source, Destination string
}

func newPair(t *testing.T) Pair {
	t.Helper()

	return Pair{Source: newFile(t), Destination: newFile(t)}
}

func copyFile(source, destination string) error {
	data, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	return os.WriteFile(destination, data, 0o600)
}

func TestCopy(t *testing.T) {
	files := newPair(t)
	if err := os.WriteFile(files.Source, []byte("hello"), 0o600); err != nil {
		t.Fatal(err)
	}

	if err := copyFile(files.Source, files.Destination); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(files.Destination)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != "hello" {
		t.Fatalf("copied = %q, want hello", data)
	}
}
