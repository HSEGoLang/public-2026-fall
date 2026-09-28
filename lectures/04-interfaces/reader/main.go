package main

import (
	"fmt"
	"io"
)

type finalReader struct{ done bool }

func (r *finalReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if r.done {
		return 0, io.EOF
	}
	r.done = true
	p[0] = 'G'
	return 1, io.EOF // полезный байт и конец потока одновременно
}

func main() {
	r := &finalReader{}
	buf := make([]byte, 8)
	for {
		n, err := r.Read(buf)
		fmt.Printf("data=%q\n", buf[:n])
		if err != nil {
			fmt.Println("EOF:", err == io.EOF)
			break
		}
	}
}
