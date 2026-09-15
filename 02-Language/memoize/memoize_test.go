package memoize_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"hse-go/02-Language/memoize"
)

func TestMemoizeIsLazyAndCallsSourceOnce(t *testing.T) {
	calls := 0
	cached := memoize.Memoize(func() (int, error) {
		calls++

		return 42, nil
	})

	require.Zero(t, calls)

	for range 5 {
		value, err := cached()
		require.NoError(t, err)
		require.Equal(t, 42, value)
	}

	require.Equal(t, 1, calls)
}

func TestMemoizeCachesError(t *testing.T) {
	wantErr := errors.New("temporary failure")
	calls := 0
	cached := memoize.Memoize(func() (int, error) {
		calls++

		return 17, wantErr
	})

	for range 3 {
		value, err := cached()
		require.Equal(t, 17, value)
		require.ErrorIs(t, err, wantErr)
	}

	require.Equal(t, 1, calls)
}

func TestMemoizeNilPanics(t *testing.T) {
	require.Panics(t, func() {
		memoize.Memoize(nil)
	})
}
