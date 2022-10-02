package piscine_go

func BasicAtoi(s string) int {
	result := 0
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		v := runes[i]
		result = result*10 + stringToInt(string(v))
	}
	return result
}

func stringToInt(s string) int {
	if s == "0" {
		return 0
	} else if s == "1" {
		return 1
	} else if s == "2" {
		return 2
	} else if s == "3" {
		return 3
	} else if s == "4" {
		return 4
	} else if s == "5" {
		return 5
	} else if s == "6" {
		return 6
	} else if s == "7" {
		return 7
	} else if s == "8" {
		return 8
	} else if s == "9" {
		return 9
	} else {
		return 0
	}
}
