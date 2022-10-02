package piscine_go

func NRune(s string, n int) rune {
	word := []rune(s)
	if n <= 0 {
		return 0
	}
	if n > len(word) {
		return 0
	}
	return word[n-1]
}
