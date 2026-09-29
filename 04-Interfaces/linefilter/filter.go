//go:build !solution

package linefilter

import "io"

func Filter(dst io.Writer, src io.Reader, keep func(string) bool) (int, error) {
	panic("implement me")
}
