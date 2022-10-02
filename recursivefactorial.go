package piscine_go

func RecursiveFactorial(nb int) int {
	var result int
	if nb < 0 || nb > 25 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if nb == 1 {
		return 1
	}
	result = nb * IterativeFactorial(nb-1)
	return result
}
