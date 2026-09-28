package main

import (
	"fmt"
	"io"
	"os"
)

type Printer interface {
	Print(line string) error
}

func Greet(out Printer, name string) error {
	return out.Print("Hello, " + name + "!")
}

type Console struct{}

func (c Console) Print(line string) error {
	_, err := fmt.Println(line)
	return err
}

type Notebook struct{ Lines []string }

func (n *Notebook) Print(line string) error {
	n.Lines = append(n.Lines, line)
	return nil
}

type FilePrinter struct{ File *os.File }

func (f FilePrinter) Print(line string) error {
	_, err := fmt.Fprintln(f.File, line)
	return err
}

func GreetWriter(out io.Writer, name string) error {
	_, err := fmt.Fprintf(out, "Hello, %s!\n", name)
	return err
}

type PrefixPrinter struct {
	Printer
	Prefix string
}

func (p PrefixPrinter) Print(line string) error {
	return p.Printer.Print(p.Prefix + line)
}

func main() {
	notebook := &Notebook{}
	for _, out := range []Printer{Console{}, notebook} {
		if err := Greet(out, "Ada"); err != nil {
			fmt.Println("error:", err)
			return
		}
	}
	fmt.Println(notebook.Lines)
}
