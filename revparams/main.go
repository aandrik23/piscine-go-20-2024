package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	param := os.Args[1:]
	cntarg := len(param) - 1
	for i := cntarg; i >= 0; i-- {
		for _, word := range param[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
