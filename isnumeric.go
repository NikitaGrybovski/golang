package piscine_go

func IsNumeric(s string) bool {
	var count int
	var length int
	for i := range s {
		length = i + 1
	}
	for _, str := range s {
		for i := '0'; i <= '9'; i++ {
			if str == i {
				count++
			}
		}
	}
	if count == length {
		return true
	}
	return false
}
