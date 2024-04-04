package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	// Handle negative numbers
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	// Handle the case when n is 0
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Create an empty slice to store digits
	digits := make([]int, 0)

	// Extract digits of n and store them in the slice
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}

	// Print digits in reverse order (since we extracted them in reverse order)
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune('0' + digits[i]))
	}
}
