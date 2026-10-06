package main

import "fmt"

func main() {
	answer := 0
	ready := make(chan struct{})
	go func() {
		answer = 2 * 2
		ready <- struct{}{}
	}()
	<-ready
	fmt.Println(answer)
}
