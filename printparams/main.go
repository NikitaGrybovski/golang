package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args

	for i, r := range arg {
		if i != 0 {
			for _, a := range r {
				z01.PrintRune(a)
			}
			z01.PrintRune('\n')
		}
	}
}
