package tp

import (
	"math"
)

func Ft_min_window(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	start, minStart, minLen := 0, 0, math.MaxInt32
	count := 0
	windowFreq := make(map[byte]int)

	for end := 0; end < len(s); end++ {
		char := s[end]

		if tFreq[char] > 0 {
			windowFreq[char]++
			if windowFreq[char] <= tFreq[char] {
				count++
			}
		}

		for count == len(t) {
			if end-start+1 < minLen {
				minLen = end - start + 1
				minStart = start
			}

			startChar := s[start]
			if tFreq[startChar] > 0 {
				windowFreq[startChar]--
				if windowFreq[startChar] < tFreq[startChar] {
					count--
				}
			}
			start++
		}
	}

	if minLen != math.MaxInt32 {
		return s[minStart : minStart+minLen]
	}

	return ""
}
