package piscine_go

func IsUpper(s string) bool {
	count := 0
	length := 0
	for i := range s {
		length = i + 1
	}
	for _, str := range s {
		for i := 'A'; i <= 'Z'; i++ {
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
