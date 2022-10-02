package piscine_go

func UltimateDivMod(a *int, b *int) {
	resA := *a / *b
	resB := *a % *b
	*a = resA
	*b = resB
}
