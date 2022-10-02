package piscine_go

func BasicJoin(elems []string) string {
	var res string
	for _, r := range elems {
		res += r
	}
	return res
}
