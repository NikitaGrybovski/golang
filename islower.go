package piscine_go

func IsLower(s string) bool {
	count := 0
	length := 0
	for i := range s {
		length = i + 1
	}
	for _, str := range s {
		for i := 'a'; i <= 'z'; i++ {
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
