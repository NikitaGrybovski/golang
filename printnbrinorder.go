package piscine_go

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n > 0 {
		var mas []int
		index := 0

		mas_count := 0

		var minIndex int

		for n != 0 {
			index = n % 10
			n /= 10
			mas = append(mas, index)
		}

		for count := range mas {
			mas_count = count + 1
		}
		for i := 0; i < mas_count-1; i++ {
			for j := 0; j < mas_count-i-1; j++ {
				if mas[j] > mas[j+1] {
					minIndex = mas[j]
					mas[j] = mas[j+1]
					mas[j+1] = minIndex
				}
			}
		}
		for i := 0; i < mas_count; i++ {
			z01.PrintRune(rune(mas[i] + 48))
		}
	}
}
