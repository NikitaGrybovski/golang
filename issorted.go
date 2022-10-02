package piscine_go

func IsSorted(f func(a, b int) int, a []int) bool {
	result := true
	result2 := true

	for i := 1; i < len(a); i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			result2 = false
		}
	}

	for i := 1; i < len(a); i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			result = false
		}
	}
	return result || result2
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
