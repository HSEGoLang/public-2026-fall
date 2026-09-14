package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.Int("port", 8080, "server port")
	flag.Parse()

	fmt.Println(*port)
}
