package piscine_go

func AppendRange(min, max int) []int {
	var mas []int

	for i := min; i < max; i++ {
		mas = append(mas, i)
	}
	return mas
}
