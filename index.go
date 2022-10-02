package piscine_go

func Index(s string, toFind string) int {
	length_s := 0
	length_toFind := 0
	for i := range s {
		length_s = i + 1
	}
	for i := range toFind {
		length_toFind = i + 1
	}

	t := 0

	for i := 0; i < length_s; i++ {
		l := 0
		t = i
		for l < length_toFind {
			if t < length_s && s[t] == toFind[l] {
				l++
				t++
			} else {
				break
			}
		}
		if l == length_toFind {
			return i
		}
	}
	return -1
}
