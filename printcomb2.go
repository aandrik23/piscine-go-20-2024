package piscine

import "github.com/01-edu/z01"

// PrintComb2 prints all possible combinations of two different two-digit numbers.
func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			z01.PrintRune(rune('0' + i/10)) // Print tens digit of i
			z01.PrintRune(rune('0' + i%10)) // Print units digit of i
			z01.PrintRune(' ')              // Print a space between numbers
			z01.PrintRune(rune('0' + j/10)) // Print tens digit of j
			z01.PrintRune(rune('0' + j%10)) // Print units digit of j

			// If not the last combination, print a comma and a space
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n') // Print a newline at the end
}
