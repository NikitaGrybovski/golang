package piscine_go

func SplitWhiteSpaces(s string) []string {
	count := 0
	word := ""
	index := 0
	for i, v := range s {
		if v == ' ' && s[i+1] != ' ' {
			count++
		}
	}
	result := make([]string, count+1)
	for _, v := range s {
		if v == ' ' || v == '\t' || v == '\n' {
			if word != "" {
				result[index] = word
				index++
				word = ""

			}
		} else {
			word += string(v)
		}
	}
	if word != "" {
		result[len(result)-1] = word
	}
	return result
}
