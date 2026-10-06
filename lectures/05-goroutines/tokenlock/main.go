package main

import (
	"fmt"
	"sync"
)

func increment(total *int, lock chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	<-lock
	defer func() { lock <- struct{}{} }()
	(*total)++
}

func main() {
	total := 0
	lock := make(chan struct{}, 1)
	lock <- struct{}{}
	var wg sync.WaitGroup
	wg.Add(2)
	go increment(&total, lock, &wg)
	go increment(&total, lock, &wg)
	wg.Wait()
	fmt.Println(total)
}
