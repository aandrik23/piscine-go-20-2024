package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	letter := os.Args[1:]

	cletter := []byte(letter)

	for i := 0; i <= len(letter)-1; {
		if letter[0] == "--upper" {
			i++
			if cletter >= "1" && cletter <= "26" {
				z01.PrintRune(cletter + 64)
			} else {
				z01.PrintRune(' ')
			}
		} else {
			if cletter >= "1" && cletter <= "26" {
				z01.PrintRune(cletter + 96)
			} else {
				z01.PrintRune(' ')
			}
		}
	}
}
