package ratelimit

import (
	"context"
	"testing"
	"time"

	"hse-go/tools/testtool"
)

func TestInvalidParameters(t *testing.T) {
	for _, params := range []struct {
		count    int
		interval time.Duration
	}{{0, time.Second}, {-1, time.Second}, {1, -time.Second}} {
		func() {
			defer func() {
				if recover() == nil {
					t.Errorf("NewLimiter(%d, %s) did not panic", params.count, params.interval)
				}
			}()

			l := NewLimiter(params.count, params.interval)
			l.Stop()
		}()
	}
}

func TestRepeatedStop(t *testing.T) {
	l := NewLimiter(1, time.Second)
	l.Stop()
	l.Stop()

	if err := l.Acquire(context.Background()); err != ErrStopped {
		t.Fatalf("Acquire after Stop = %v, want ErrStopped", err)
	}
}

func TestNoSyncPackageImported(t *testing.T) {
	testtool.CheckForbiddenImport(t, "sync")
}
