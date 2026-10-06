package main

import (
	"fmt"
	"sync"
)

func configure() {
	fmt.Println("начало настройки")
	fmt.Println("конец настройки")
}

func main() {
	var once sync.Once
	var wg sync.WaitGroup
	for range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(configure)
			fmt.Println("после Do")
		}()
	}
	wg.Wait()
}
