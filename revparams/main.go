package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	param := os.Args
	cntarg := len(param)
	for i := cntarg; i == 1; i-- {
		for _, word := range param[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
