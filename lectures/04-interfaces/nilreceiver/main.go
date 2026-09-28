package main

import "fmt"

type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	var empty *IntList
	fmt.Println(empty.Sum())
	list := &IntList{Value: 3, Tail: &IntList{Value: 4}}
	fmt.Println(list.Sum())
}
