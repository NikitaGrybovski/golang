package piscine_go

func Map(f func(int) bool, a []int) []bool {
	arr := make([]bool, len(a))
	for i := range a {
		arr[i] = f(a[i])
	}
	return arr
}
