package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	n := len(input)
	reverse := make([]rune, n)
	for _, l := range input {
		n--
		reverse[n] = l
	}
	output = string(reverse)
	return output
}
