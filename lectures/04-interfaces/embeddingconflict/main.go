package main

import "fmt"

// LOGGERS OMIT
type LoggerA struct{}

func (LoggerA) Log(s string) { fmt.Println("A:", s) }

type LoggerB struct{}

func (LoggerB) Log(s string) { fmt.Println("B:", s) }

type S struct {
	LoggerA
	LoggerB
}

// END LOGGERS OMIT

func main() {
	var s S
	// s.Log("hi")
	s.LoggerA.Log("hi")
	s.LoggerB.Log("hi")
}
