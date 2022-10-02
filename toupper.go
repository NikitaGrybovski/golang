package piscine_go

func ToUpper(s string) string {
	mas := []rune(s)
	for i := range mas {
		if mas[i] >= 'a' && mas[i] <= 'z' {
			mas[i] = mas[i] - 32
		}
	}
	return string(mas)
}
