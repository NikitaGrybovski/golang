package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) > 0 {
		for i, r := range arg {
			arg, err := ioutil.ReadFile(arg[i])
			if err != nil {
				errFile := "ERROR: open " + string(r) + ": no such file or directory"
				printStr(errFile)
				z01.PrintRune('\n')
				os.Exit(1)
			}
			printStr(string(arg))
		}
	} else {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
	}
}

func printStr(s string) {
	sRune := []rune(s)
	for i := 0; i < len(sRune); i++ {
		z01.PrintRune(sRune[i])
	}
}
