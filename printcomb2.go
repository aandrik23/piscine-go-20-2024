package piscine

import "github.com/01-edu/z01"

// PrintComb2 prints all possible combinations of two different two-digit numbers.
func PrintComb2() {
	var i, j int
	for i = 0; i <= 98; i++ {
		for j = i + 1; j <= 99; j++ {
			if i < 10 {
				z01.PrintRune('0')
			}
			z01.PrintRune(rune('0' + i))
			z01.PrintRune(' ')
			if j < 10 {
				z01.PrintRune('0')
			}
			z01.PrintRune(rune('0' + j))
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
