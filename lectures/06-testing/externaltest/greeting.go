package greeting

import "strings"

func normalizeName(name string) string {
	return strings.TrimSpace(name)
}

func Hello(name string) string {
	return "Hello, " + normalizeName(name) + "!"
}
