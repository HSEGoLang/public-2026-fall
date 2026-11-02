package main

import "fmt"

// START TELLER OMIT
func teller(
	deposits <-chan int, balances chan<- int,
	stop <-chan struct{}, done chan<- struct{},
) {
	defer close(done)

	var balance int

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case <-stop:
			return
		}
	}
}

// END TELLER OMIT

// START ACTOR_MAIN OMIT
func main() {
	deposits := make(chan int)
	balances := make(chan int)

	stop, done := make(chan struct{}), make(chan struct{})
	go teller(deposits, balances, stop, done)

	deposits <- 200

	deposits <- 100

	fmt.Println(<-balances)
	close(stop)
	<-done
}

// END ACTOR_MAIN OMIT
