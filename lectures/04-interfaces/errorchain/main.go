package main

import (
	"errors"
	"fmt"
)

type OfflineError struct{ Printer string }

func (e *OfflineError) Error() string {
	return "printer offline: " + e.Printer
}

func main() {
	cause := &OfflineError{Printer: "hall"}
	err := fmt.Errorf("print greeting: %w", cause)
	fmt.Println(err == cause)
	fmt.Println(errors.Is(err, cause))
	var offline *OfflineError
	if errors.As(err, &offline) {
		fmt.Println(offline.Printer)
	}
}
