package main

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	memo1 "hse-go/lectures/08-concurrency/memo1"
	memo2 "hse-go/lectures/08-concurrency/memo2"
	memo3 "hse-go/lectures/08-concurrency/memo3"
	memo4 "hse-go/lectures/08-concurrency/memo4"
	"hse-go/lectures/08-concurrency/synconce"
)

type cache interface {
	Get(string) (any, error)
}

func newMemo(version string, compute func(string) (any, error)) cache {
	switch version {
	case "1":
		return memo1.New(compute)
	case "2":
		return memo2.New(compute)
	case "3":
		return memo3.New(compute)
	case "4":
		return memo4.New(compute)
	default:
		panic("choose cache version 1, 2, 3, or 4")
	}
}

type outcome struct {
	value any
	err   error
}

func makeCache(version string, compute func(string) (any, error)) cache {
	if version == "once" {
		return synconce.New(compute)
	}

	return newMemo(version, compute)
}

func getAsync(memo cache, key string) <-chan outcome {
	finished := make(chan outcome, 1)

	go func() {
		value, err := memo.Get(key)
		finished <- outcome{value: value, err: err}
	}()

	return finished
}

func takeResult(t *testing.T, finished <-chan outcome) outcome {
	t.Helper()

	select {
	case result := <-finished:
		return result
	case <-time.After(time.Second):
		t.Fatal("Get did not return before the test deadline")

		return outcome{}
	}
}

func takeStarted(t *testing.T, started <-chan string) string {
	t.Helper()

	select {
	case key := <-started:
		return key
	case <-time.After(time.Second):
		t.Fatal("loader did not start before the test deadline")

		return ""
	}
}

func TestCachesMemoizeValuesAndErrors(t *testing.T) {
	for _, version := range []string{"1", "2", "3", "4", "once"} {
		t.Run(version, func(t *testing.T) {
			calls := make(map[string]int)
			missing := errors.New("missing card")
			memo := makeCache(version, func(key string) (any, error) {
				calls[key]++
				if key == "missing" {
					return nil, missing
				}

				return "card:" + key, nil
			})

			for range 2 {
				value, err := memo.Get("tea")
				require.NoError(t, err)
				require.Equal(t, "card:tea", value)

				value, err = memo.Get("missing")
				require.ErrorIs(t, err, missing)
				require.Nil(t, value)
			}

			require.Equal(t, map[string]int{"tea": 1, "missing": 1}, calls)

			other := makeCache(version, func(string) (any, error) {
				return "other card", nil
			})
			value, err := other.Get("tea")
			require.NoError(t, err)
			require.Equal(t, "other card", value, "instances must have independent caches")
		})
	}
}

func TestMemo2SerializesDifferentKeys(t *testing.T) {
	started := make(chan string, 2)
	release := make(chan struct{})

	releaseAll := sync.OnceFunc(func() { close(release) })
	defer releaseAll()

	var active atomic.Int32

	memo := newMemo("2", func(key string) (any, error) {
		if active.Add(1) != 1 {
			t.Error("memo2 ran two loaders concurrently")
		}

		defer active.Add(-1)

		started <- key

		<-release

		return key, nil
	})
	first := getAsync(memo, "tea")

	require.Equal(t, "tea", takeStarted(t, started))

	second := getAsync(memo, "coffee")
	// This bounded observation gives the second call an opportunity to run;
	// it does not assert that a goroutine has reached an internal mutex.
	select {
	case key := <-started:
		t.Fatalf("loader for %q started before the first loader was released", key)
	case <-time.After(20 * time.Millisecond):
	}

	releaseAll()
	require.Equal(t, "tea", takeResult(t, first).value)
	require.Equal(t, "coffee", takeResult(t, second).value)
}

func TestMemo3CanDuplicateSameKey(t *testing.T) {
	started := make(chan string, 2)
	release := make(chan struct{})

	releaseAll := sync.OnceFunc(func() { close(release) })
	defer releaseAll()

	var calls atomic.Int32

	memo := newMemo("3", func(key string) (any, error) {
		calls.Add(1)

		started <- key

		<-release

		return key, nil
	})
	first := getAsync(memo, "tea")

	require.Equal(t, "tea", takeStarted(t, started))

	second := getAsync(memo, "tea")

	require.Equal(t, "tea", takeStarted(t, started))

	releaseAll()
	require.Equal(t, "tea", takeResult(t, first).value)
	require.Equal(t, "tea", takeResult(t, second).value)
	require.EqualValues(t, 2, calls.Load())
}

func TestSharedCachesComputeSameKeyOnce(t *testing.T) {
	for _, version := range []string{"4", "once"} {
		t.Run(version, func(t *testing.T) {
			started := make(chan string, 2)
			release := make(chan struct{})

			releaseAll := sync.OnceFunc(func() { close(release) })
			defer releaseAll()

			var calls atomic.Int32

			memo := makeCache(version, func(key string) (any, error) {
				calls.Add(1)

				started <- key

				<-release

				return "card:" + key, nil
			})
			first := getAsync(memo, "tea")

			require.Equal(t, "tea", takeStarted(t, started))

			second := getAsync(memo, "tea")
			// Give the second caller an opportunity to reach Get. The timeout is
			// an observation window, not a guarantee about goroutine scheduling.
			select {
			case <-second:
				t.Fatal("Get returned before the shared computation finished")
			case <-started:
				t.Fatal("same key started a second computation")
			case <-time.After(20 * time.Millisecond):
			}

			releaseAll()

			for _, finished := range []<-chan outcome{first, second} {
				result := takeResult(t, finished)
				require.NoError(t, result.err)
				require.Equal(t, "card:tea", result.value)
			}

			require.EqualValues(t, 1, calls.Load())
		})
	}
}

func TestDifferentKeysComputeConcurrently(t *testing.T) {
	for _, version := range []string{"3", "4", "once"} {
		t.Run(version, func(t *testing.T) {
			started := make(chan string, 2)
			release := make(chan struct{})

			releaseAll := sync.OnceFunc(func() { close(release) })
			defer releaseAll()

			memo := makeCache(version, func(key string) (any, error) {
				started <- key

				<-release

				return key, nil
			})
			first := getAsync(memo, "tea")

			require.Equal(t, "tea", takeStarted(t, started))

			second := getAsync(memo, "coffee")

			require.Equal(t, "coffee", takeStarted(t, started))

			releaseAll()
			require.Equal(t, "tea", takeResult(t, first).value)
			require.Equal(t, "coffee", takeResult(t, second).value)
		})
	}
}
