package main

import "fmt"

func main() {
	var callbacks []func()
	for _, value := range []string{"A", "B", "C"} {
		callbacks = append(callbacks, func() { fmt.Print(value) })
	}
	for _, callback := range callbacks {
		callback()
	}
	fmt.Println()
}
