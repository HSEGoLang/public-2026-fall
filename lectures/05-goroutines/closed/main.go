package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 7
	close(ch)
	for range 3 {
		value, ok := <-ch
		fmt.Println(value, ok)
	}
}
