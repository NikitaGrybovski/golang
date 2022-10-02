package piscine_go

func AlphaCount(s string) int {
	var count int
	for _, str := range s {
		for i := 'a'; i <= 'z'; i++ {
			if str == i {
				count++
			}
		}
		for i := 'A'; i <= 'Z'; i++ {
			if str == i {
				count++
			}
		}
	}
	return count
}
