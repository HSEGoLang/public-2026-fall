package main

import (
	"context"
	"fmt"
)

// START VALUE_HELPERS OMIT
type myKey struct{}

func WithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, myKey{}, user)
}

func ContextUser(ctx context.Context) (string, bool) {
	v := ctx.Value(myKey{})
	s, ok := v.(string)

	return s, ok
}

// END VALUE_HELPERS OMIT

// START VALUE_MAIN OMIT
func main() {
	ctx := context.Background()

	user, ok := ContextUser(ctx)
	fmt.Printf("user=%q found=%t\n", user, ok)

	ctx = WithUser(ctx, "petya")
	user, ok = ContextUser(ctx)
	fmt.Printf("user=%q found=%t\n", user, ok)
}

// END VALUE_MAIN OMIT
