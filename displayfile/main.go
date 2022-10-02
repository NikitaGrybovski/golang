package main

import (
	"fmt"
	"os"
)

func main() {
	count := 0

	arg := os.Args[1:]

	for i := range arg {
		count = i + 1
	}

	if count == 0 {
		fmt.Println("File name missing")
	} else if count > 1 {
		fmt.Println("Too many arguments")
	} else if arg[0] == "quest8.txt" {
		content, err := os.Open(arg[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data := make([]byte, 14)
		content.Read(data)
		fmt.Println(string(data))
		content.Close()
	}
}
