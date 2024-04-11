package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args
	for i := 0; i < len(name)-1; {
		if name[i] > name[i+1] {
			name[i], name[i+1] = name[i+1], name[i]
			i = 0
		} else {
			i++
		}
	}
	for _, word := range name {
		for _, letter := range word {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
