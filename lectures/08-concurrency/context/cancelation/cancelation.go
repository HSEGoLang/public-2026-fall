package cancelation

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// START CANCEL OMIT
func SimpleCancelation() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := make(chan int)
	done := make(chan error)

	go func() {
		_, err := receive(ctx, in)
		done <- err
	}()

	cancel()

	err := <-done
	fmt.Println("canceled:", errors.Is(err, context.Canceled))
}

// END CANCEL OMIT

// START TIMEOUT OMIT
func SimpleTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err := receive(ctx, make(chan int))
	fmt.Println("deadline:", errors.Is(err, context.DeadlineExceeded))
}

// END TIMEOUT OMIT

// START RECEIVE OMIT
func receive(ctx context.Context, in <-chan int) (int, error) {
	select {
	case value := <-in:
		return value, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// END RECEIVE OMIT

// START CAUSE OMIT
func CancelCause() {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	cancel(errors.New("maintenance"))
	fmt.Println("cause:", context.Cause(ctx))
}

// END CAUSE OMIT
