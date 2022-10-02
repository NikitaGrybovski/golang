package piscine_go

func BasicAtoi2(s string) int {
	result := 0
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		v := runes[i]
		result = result*10 + int(v) - '0'
	}
	return result
}
