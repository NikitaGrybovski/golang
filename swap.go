package piscine_go

func Swap(a *int, b *int) {
	*b, *a = *a, *b
}
