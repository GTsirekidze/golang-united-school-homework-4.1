package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	r := []rune(input)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return output
}
