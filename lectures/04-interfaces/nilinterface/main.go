package main

import "fmt"

type Printer interface{ Print(string) error }
type Notebook struct{ Lines []string }

func (n *Notebook) Print(line string) error {
	n.Lines = append(n.Lines, line)
	return nil
}

func main() {
	var notebook *Notebook
	var out Printer = notebook
	fmt.Println(notebook == nil, out == nil)
	fmt.Printf("%T\n", out)
	// out.Print("hello")
	var fixed Printer
	fmt.Println(fixed == nil)
}
