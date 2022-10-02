package piscine_go

func Split(s, sep string) []string {
	split := 0
	arr := []string{}

	for index := range s {
		if index < len(s)-len(sep) {
			if s[index:index+len(sep)] == sep {
				arr = append(arr, s[split:index])
				split = index + len(sep)
			}
		}
		if index == len(s)-1 {
			arr = append(arr, s[split:])
		}
	}
	return arr
}
