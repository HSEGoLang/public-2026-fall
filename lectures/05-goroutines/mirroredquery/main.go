package main

import (
	"fmt"
	"time"
)

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses
}

func request(hostname string) string {
	delays := map[string]time.Duration{
		"asia.gopl.io":     60 * time.Millisecond,
		"europe.gopl.io":   20 * time.Millisecond,
		"americas.gopl.io": 40 * time.Millisecond,
	}
	time.Sleep(delays[hostname]) // model waiting for a server; no HTTP
	return hostname
}

func main() {
	fmt.Println(mirroredQuery())
}
