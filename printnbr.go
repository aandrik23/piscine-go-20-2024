package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		// If n is the minimum possible integer value,
		// print it as a string
		str := "-9223372036854775808"
		for _, ch := range str {
			z01.PrintRune(ch)
		}
	} else if n < 0 {
		z01.PrintRune('-')
		printPositiveNumber(-n)
	} else {
		printPositiveNumber(n)
	}
}

func printPositiveNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	buffer := make([]rune, 0)

	for n > 0 {
		digit := n % 10
		buffer = append(buffer, rune('0'+digit))
		n /= 10
	}

	for i := len(buffer) - 1; i >= 0; i-- {
		z01.PrintRune(buffer[i])
	}
}
