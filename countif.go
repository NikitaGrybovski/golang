package piscine_go

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, r := range tab {
		if f(r) {
			count++
		}
	}
	return count
}
