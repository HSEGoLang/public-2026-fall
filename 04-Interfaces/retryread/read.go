//go:build !solution

package retryread

func (e *AttemptError) Error() string { return "implement me" }
func (e *AttemptError) Unwrap() error { return nil }

func Read(client Client, key string, attempts int) (string, error) {
	panic("implement me")
}
