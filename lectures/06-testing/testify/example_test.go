package testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	if got, want := Sum(1, 2), 4; got != want {
		t.Errorf("Sum(%d, %d) = %d, want %d", 1, 2, got, want)
	}
}

func TestSum0(t *testing.T) {
	assert.Equalf(t, 4, Sum(1, 2), "Sum(%d, %d)", 1, 2)
}

func TestSumWithRequire(t *testing.T) {
	require.Equal(t, 3, Sum(1, 2))
	t.Log("после успешной проверки")
}

func TestSumFromWorker(t *testing.T) {
	result := make(chan int, 1)
	go func() { result <- Sum(1, 2) }()

	got := <-result
	require.Equal(t, 3, got)
}

func assertGood(t *testing.T, i int) {
	t.Helper()

	if i != 0 {
		t.Errorf("i (%d) != 0", i)
	}
}

func TestA(t *testing.T) {
	assertGood(t, 0)
	assertGood(t, 1)
}
