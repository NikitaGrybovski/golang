package piscine_go

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 25 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if nb == 1 {
		return 1
	}
	for i := 1; i <= nb; i++ {
		result = result * i
	}
	return result
}
