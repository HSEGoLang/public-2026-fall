package main

import "fmt"

func describe(value any) string {
	switch v := value.(type) {
	case nil:
		return "nothing"
	case string:
		return fmt.Sprintf("text: %s", v)
	case int, uint:
		return fmt.Sprintf("number: %v", v)
	default:
		return fmt.Sprintf("other: %T", v)
	}
}

func main() {
	for _, value := range []any{nil, "Go", 7, true} {
		fmt.Println(describe(value))
	}
}
