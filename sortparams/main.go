package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	mas := os.Args[1:]
	for i := 0; i < len(mas)-1; i++ {
		for j := 0; j < len(mas)-1; j++ {
			if mas[j] > mas[j+1] {
				mas[j], mas[j+1] = mas[j+1], mas[j]
			}
		}
	}

	for _, a := range mas {
		for _, r := range a {
			z01.PrintRune(rune(r))
		}
		z01.PrintRune('\n')
	}
}
