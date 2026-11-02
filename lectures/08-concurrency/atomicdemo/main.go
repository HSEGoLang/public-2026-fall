package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// START STORE_BUFFERING OMIT
func bothZero() bool {
	// START SB_STATE OMIT
	var (
		x, y         atomic.Int32
		readX, readY int32
		wg           sync.WaitGroup
	)

	wg.Add(2)
	// END SB_STATE OMIT

	// START SB_WORKERS OMIT
	go func() {
		defer wg.Done()

		x.Store(1)

		readY = y.Load()
	}()
	go func() {
		defer wg.Done()

		y.Store(1)

		readX = x.Load()
	}()
	// END SB_WORKERS OMIT

	// START SB_JOIN OMIT
	wg.Wait()

	return readX == 0 && readY == 0
	// END SB_JOIN OMIT
}

// END STORE_BUFFERING OMIT

func main() {
	fmt.Println("both zero:", bothZero())
}
