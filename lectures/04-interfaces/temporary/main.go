package main

import (
	"errors"
	"fmt"
)

type Temporary interface {
	IsTemporary() bool
}

type BusyError struct{ Printer string }

func (e *BusyError) Error() string     { return "printer busy: " + e.Printer }
func (e *BusyError) IsTemporary() bool { return true }

func main() {
	err := fmt.Errorf("print greeting: %w", &BusyError{Printer: "hall"})
	_, direct := err.(Temporary)
	var temporary Temporary
	found := errors.As(err, &temporary)
	fmt.Println("assertion:", direct)
	fmt.Println("errors.As:", found)
	if found {
		fmt.Println("temporary:", temporary.IsTemporary())
	}
}
