package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	var caps bool
	num := 0
	for _, ar := range arg {
		if ar == "--upper" {
			arg = arg[1:]
			caps = true
		}
	}
	for _, r := range arg {
		num = 0
		for _, rn := range r {
			num = num*10 + int(rn-'0')
		}

		if num >= 1 && num <= 26 {
			if !caps {
				z01.PrintRune(rune(num) + 96)
			} else {
				z01.PrintRune(rune(num) + 64)
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	if num != 0 {
		z01.PrintRune('\n')
	}
}
