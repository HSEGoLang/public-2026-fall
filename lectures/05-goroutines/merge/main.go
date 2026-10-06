package main

import (
	"fmt"
	"sync"
)

func merge(left, right <-chan int, out chan<- int) {
	var wg sync.WaitGroup
	forward := func(in <-chan int) {
		defer wg.Done()
		for value := range in {
			out <- value
		}
	}
	wg.Add(2)
	go forward(left)
	go forward(right)
	wg.Wait()
	close(out)
}

func main() {
	left, right := make(chan int, 2), make(chan int, 2)
	left <- 1
	left <- 2
	right <- 10
	right <- 20
	close(left)
	close(right)
	out := make(chan int)
	go merge(left, right, out)
	for value := range out {
		fmt.Println(value)
	}
}
