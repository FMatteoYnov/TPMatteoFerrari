package tp

func Ft_max_substring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {
		char := s[i]

		if index, found := charIndex[char]; found && index >= start {
			start = index + 1
		}

		charIndex[char] = i

		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
