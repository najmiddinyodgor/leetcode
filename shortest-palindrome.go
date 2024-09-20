package main

func ShortestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	reversed := string(runes)

	for i := 0; i < len(s); i++ {
		if s[:len(s)-i] == reversed[i:] {
			return reversed[:i] + s
		}
	}

	return ""
}
