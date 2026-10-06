package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(0)
	defer timer.Stop()
	<-timer.C
	fmt.Println("таймер")
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()
	for range 2 {
		<-ticker.C
		fmt.Println("тик")
	}
}
