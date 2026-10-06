package main

import "fmt"

func ownVars() []func() int {
	var calls []func() int
	for i := 0; i < 3; i++ {
		calls = append(calls, func() int { return i })
	}
	return calls
}

func sharedVar() []func() int {
	var calls []func() int
	var i int
	for i = 0; i < 3; i++ {
		calls = append(calls, func() int { return i })
	}
	return calls
}

func printCalls(label string, calls []func() int) {
	values := make([]int, len(calls))
	for index, call := range calls {
		values[index] = call()
	}
	fmt.Println(label, values)
}

func main() {
	printCalls(":=", ownVars())
	printCalls("=", sharedVar())
}
