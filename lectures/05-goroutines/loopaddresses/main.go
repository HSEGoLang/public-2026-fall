package main

import "fmt"

func pointerKeys() int {
	seen := make(map[*int]bool)
	for i := 0; i < 100_000; i++ {
		seen[&i] = true
	}
	return len(seen)
}

func stringKeys() int {
	seen := make(map[string]bool)
	for i := 0; i < 100_000; i++ {
		seen[fmt.Sprintf("%p", &i)] = true
	}
	return len(seen)
}

func main() {
	fmt.Println("указатели:", pointerKeys())
	fmt.Println("строки:", stringKeys())
}
