package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	r := []rune(input)

	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
	}

	return output
}
