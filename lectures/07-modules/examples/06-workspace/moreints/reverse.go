package moreints

func Reverse(input []int) []int {
	result := make([]int, len(input))
	copy(result, input)

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
