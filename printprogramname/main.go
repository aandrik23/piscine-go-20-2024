package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	for _, v := range name {
		z01.PrintRune(rune(v))
	}
	z01.PrintRune('\n')

}
