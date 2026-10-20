package morestrings

func Reverse(input []rune) []rune {
	result := make([]rune, len(input))
	copy(result, input)

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
