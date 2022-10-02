package piscine_go

func checkChar(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}

func Capitalize(s string) string {
	mas := []rune(s)

	for i, r := range mas {
		if checkChar(r) {
			if i == 0 || checkChar(mas[i-1]) == false {
				if mas[i] >= 'a' && mas[i] <= 'z' {
					mas[i] = r - 32
				}
			} else {
				if mas[i] >= 'A' && mas[i] <= 'Z' {
					mas[i] = r + 32
				}
			}
		}
	}
	return string(mas)
}
