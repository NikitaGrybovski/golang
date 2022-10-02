package piscine_go

func Sqrt(nb int) int {
	var result int
	result = nb / 2
	if nb == 1 {
		return 1
	}
	if nb > 0 {
		for i := 2; i < result; i++ {
			result = (nb/result + result) / 2
		}
		if result == 1 {
			return 0
		}
		if nb%result == 0 && (nb/result+result)%2 == 0 {
			return result
		} else {
			return 0
		}
	} else {
		return 0
	}
}
