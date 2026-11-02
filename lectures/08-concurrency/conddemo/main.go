package main

import "fmt"

// START MAILBOX_MAIN OMIT
func main() {
	m := NewMailbox()
	done := make(chan struct{})

	go func() {
		m.Put(7)
		close(done)
	}()

	fmt.Println(m.Get())
	<-done
}

// END MAILBOX_MAIN OMIT
