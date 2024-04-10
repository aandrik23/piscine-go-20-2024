package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the full path of the program
	fullPath := os.Args[0]
	// Find the last '/' character to get the executable name
	executableName := fullPath
	for i := len(fullPath) - 1; i >= 0; i-- {
		if fullPath[i] == '/' {
			executableName = fullPath[i+1:]
			break
		}
	}
	// Iterate over each character in the executable name
	for _, char := range executableName {
		z01.PrintRune(char)
	}
	// Print a newline character at the end
	z01.PrintRune('\n')
}
