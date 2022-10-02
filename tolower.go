package piscine_go

func ToLower(s string) string {
	mas := []rune(s)

	for i := range mas {
		if mas[i] >= 'A' && mas[i] <= 'Z' {
			mas[i] = mas[i] + 32
		}
	}
	return string(mas)
}
