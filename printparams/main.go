package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, parametr := range os.Args[1:] {
		for _, word := range parametr {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}

}
