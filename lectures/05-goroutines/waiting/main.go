package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("работа")
		}()
	}
	wg.Wait()
	fmt.Println("всё готово")
}
