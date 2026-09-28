package main

import "fmt"

type Printer interface{ Print(string) error }

type Console struct{}

func (c Console) Print(line string) error {
	_, err := fmt.Println(line)
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
	out := PrefixPrinter{Printer: Console{}, Prefix: "[course] "}
	err := out.Print("Hello, Ada!")
	if err != nil {
		fmt.Println("error:", err)
	}
}
