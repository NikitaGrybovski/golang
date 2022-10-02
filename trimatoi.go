package piscine_go

func TrimAtoi(s string) int {
	var mas []int
	index := 0
	minus := -1
	first_index := 0
	mas_count := 0
	result := 0

	for _, run := range s {
		if run == '-' {
			minus = index
		}
		for i := '0'; i <= '9'; i++ {
			if run == i {
				if first_index == 0 {
					first_index = index
				}
				mas = append(mas, int(run-'0'))
			}
			index++
		}
	}
	for count := range mas {
		mas_count = count + 1
	}

	for i := 0; i < mas_count; i++ {
		result = result*10 + mas[i]
	}

	if minus < first_index && minus != -1 {
		result = result * -1
	}
	return result
}
