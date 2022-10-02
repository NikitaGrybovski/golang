package piscine_go

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	mas := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		mas[i] = i + min
	}
	return mas
}
