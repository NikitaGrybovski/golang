package piscine_go

func RecursivePower(nb int, power int) int {
	var result int
	if power > 0 {
		result = nb * RecursivePower(nb, power-1)
		return result
	} else if nb == 0 && power == 0 || power == 0 {
		return 1
	} else {
		return 0
	}
}
