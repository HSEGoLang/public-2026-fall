package main

import "fmt"

func execute(action func()) (err error) {
	defer func() {
		if value := recover(); value != nil {
			err = fmt.Errorf("action failed: %v", value)
		}
	}()
	action()
	return nil
}

func main() {
	err := execute(func() { panic("printer broken") })
	fmt.Println(err)
	fmt.Println("caller continues")
}
