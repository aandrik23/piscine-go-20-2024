package main

import "github.com/01-edu/z01"

func main() {
	// Iterate over each letter of the alphabet
	for c := 'z'; c >= 'a'; c-- {
		// Print the current letter
		z01.PrintRune(c)
	}
	// Print a newline character after printing all the letters
	z01.PrintRune('\n')
}
