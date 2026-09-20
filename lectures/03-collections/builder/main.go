package main

import (
	"fmt"
	"strconv"
	"strings"
)

func formatInts(values []int) string {
	var builder strings.Builder
	builder.WriteByte('[')

	for index, value := range values {
		if index > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(strconv.Itoa(value))
	}

	builder.WriteByte(']')
	return builder.String()
}

func main() {
	fmt.Println(formatInts([]int{10, 20, 30}))
}
