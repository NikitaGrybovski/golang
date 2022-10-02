package piscine_go

func ConcatParams(args []string) string {
	var word string
	for i, w := range args {
		word += w
		if i != len(args)-1 {
			word += "\n"
		}
	}
	return word
}
