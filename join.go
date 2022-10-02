package piscine_go

func Join(strs []string, sep string) string {
	var res string

	for i, r := range strs {
		res += r
		if i < len(strs)-1 {
			res += sep
		}
	}
	return res
}
