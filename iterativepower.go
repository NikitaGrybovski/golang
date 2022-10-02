package piscine_go

func IterativePower(nb int, power int) int {
	result := nb
	if power > 0 {
		for i := 1; i < power; i++ {
			result = result * nb
		}
		return result
	} else if nb == 0 && power == 0 || power == 0 {
		return 1
	} else {
		return 0
	}
}
