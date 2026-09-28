package main

import (
	"fmt"
	"os"
)

func processAll(names []string) error {
	for _, name := range names {
		file, err := os.Open(name)
		if err != nil {
			return err
		}
		defer file.Close()
		if err := process(file); err != nil {
			return err
		}
	}
	return nil
}

func processAllFixed(names []string) error {
	for _, name := range names {
		if err := processOne(name); err != nil {
			return err
		}
	}
	return nil
}

func processOne(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	return process(file)
}

// The processing step uses the file and leaves ownership to its caller.
func process(file *os.File) error {
	_, err := file.Stat()
	return err
}

// Run the compiled demo with a low per-process file limit to show exhaustion:
// (ulimit -n 64; ./deferloop)
// A finite list bounds the demo even without that limit.
func main() {
	file, err := os.CreateTemp("", "deferloop-*")
	if err != nil {
		fmt.Println("setup:", err)
		return
	}
	name := file.Name()
	defer os.Remove(name)
	if err := file.Close(); err != nil {
		fmt.Println("setup close:", err)
		return
	}
	names := make([]string, 128)
	for i := range names {
		names[i] = name
	}
	fmt.Println("loop:", processAll(names))
	fmt.Println("helper:", processAllFixed(names))
}
