package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	count := 0

	for s := range arg {
		count = s + 1
	}
	for i := count - 1; i >= 1; i-- {
		for _, r := range arg[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
