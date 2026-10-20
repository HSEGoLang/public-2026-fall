package worker

import (
	"context"
	"sync"
	"testing"
)

func TestWorker(_ *testing.T) {
	var wg sync.WaitGroup
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
	}()
}
