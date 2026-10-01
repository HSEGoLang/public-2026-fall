package retryread

import (
	"errors"
	"fmt"
	"testing"
)

type transient struct{}

func (transient) Error() string     { return "busy" }
func (transient) IsTemporary() bool { return true }

type response struct {
	value string
	err   error
}
type fakeClient struct {
	responses []response
	keys      []string
}

func (c *fakeClient) Get(key string) (string, error) {
	c.keys = append(c.keys, key)
	if len(c.responses) == 0 {
		panic("unexpected request")
	}

	r := c.responses[0]
	c.responses = c.responses[1:]

	return r.value, r.err
}

func TestRetryThenSuccess(t *testing.T) {
	c := &fakeClient{
		responses: []response{{"", fmt.Errorf("server: %w", transient{})}, {"value", nil}},
	}

	value, err := Read(c, "theme", 3)
	if err != nil || value != "value" || len(c.keys) != 2 {
		t.Fatal(value, err, c.keys)
	}

	for _, key := range c.keys {
		if key != "theme" {
			t.Fatal(key)
		}
	}
}

func TestPermanentAndExhausted(t *testing.T) {
	for _, cause := range []error{fmt.Errorf("get: %w", ErrNotFound), errors.New("broken"), transient{}} {
		c := &fakeClient{responses: []response{{"discard", cause}}}
		value, err := Read(c, "key", 1)

		var attempt *AttemptError
		if value != "" || !errors.As(err, &attempt) || attempt.Attempts != 1 ||
			attempt.Key != "key" ||
			attempt.Err != cause ||
			!errors.Is(err, cause) {
			t.Fatal(value, err)
		}
	}
}

func TestInvalidAttempts(t *testing.T) {
	c := &fakeClient{}
	for _, attempts := range []int{0, -1} {
		if v, err := Read(c, "key", attempts); v != "" || !errors.Is(err, ErrAttempts) {
			t.Fatal(v, err)
		}
	}

	if len(c.keys) != 0 {
		t.Fatal(c.keys)
	}
}
