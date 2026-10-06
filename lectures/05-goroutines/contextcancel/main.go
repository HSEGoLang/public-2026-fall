package main

import (
	"context"
	"fmt"
)

func generate(ctx context.Context, out chan<- int) {
	defer close(out)
	for value := 0; ; value++ {
		select {
		case out <- value:
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	background := context.Background()
	ctx, cancel := context.WithCancel(background)
	defer cancel()
	out := make(chan int)
	go generate(ctx, out)
	fmt.Println(<-out)
	cancel()
	for value := range out {
		fmt.Println(value)
	}
	fmt.Println("остановлено")
}
