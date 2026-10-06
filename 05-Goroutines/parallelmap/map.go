//go:build !solution

package parallelmap

// Map applies f to every value using at most workers concurrent calls.
// It preserves the input order and waits for all its workers to finish.
func Map(values []int, workers int, f func(int) int) []int {
	_, _, _ = values, workers, f
	return nil // TODO: implement the bounded worker pool.
}
