package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	index := 0
	arg := os.Args[0]

	for i, r := range arg {
		if r == '/' {
			index = i
		}
	}
	for _, r := range arg[index+1:] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
