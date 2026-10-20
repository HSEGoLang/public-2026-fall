package tparallel

import (
	"sync"
	"testing"
)

func TestEmptyRun(t *testing.T) {
	_ = t

	Run(nil)
	Run([]func(*T){})
}

func TestUniqueInstancesAndIndependentRuns(t *testing.T) {
	var (
		mu   sync.Mutex
		seen = make(map[*T]struct{})
		wg   sync.WaitGroup
	)

	record := func(test *T) {
		mu.Lock()
		defer mu.Unlock()

		if _, exists := seen[test]; exists {
			t.Error("T reused between test invocations")
		}

		seen[test] = struct{}{}
	}

	for range 3 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			Run([]func(*T){
				func(test *T) {
					record(test)
					test.Parallel()
					test.Run(record)
				},
				func(test *T) {
					record(test)
					test.Parallel()
					test.Run(record)
				},
			})
		}()
	}

	wg.Wait()

	if len(seen) != 12 {
		t.Errorf("completed %d test invocations; want 12", len(seen))
	}
}
