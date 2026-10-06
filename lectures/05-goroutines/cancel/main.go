package main

import "fmt"

func produce(out chan<- int, done <-chan struct{}, finished chan<- struct{}) {
	defer close(finished)
	for value := 0; ; value++ {
		select {
		case out <- value:
		case <-done:
			return
		}
	}
}

func main() {
	out := make(chan int)
	done, finished := make(chan struct{}), make(chan struct{})
	go produce(out, done, finished)
	fmt.Println(<-out)
	close(done)
	<-finished
	fmt.Println("остановлено")
}
