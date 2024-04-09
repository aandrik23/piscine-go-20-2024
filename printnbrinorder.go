package piscine

import (
	"github.com/01-edu/z01"
)

func OrderNumber(n int) []int {
	var digits []int
	if n == 0 {
		digits = append(digits, 0) // Handle the case where n is 0
	} else {
		// Extract digits and store them in a slice
		for n > 0 {
			digits = append(digits, n%10)
			n /= 10
		}
	}
	// Sort the digits in the slice using a simple sorting algorithm (e.g., bubble sort)
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits)-1-i; j++ {
			if digits[j] > digits[j+1] {
				// Swap the elements
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}
	return digits
}

func PrintNbrInOrder(n int) {
	sortedDigits := OrderNumber(n)
	for _, digit := range sortedDigits {
		z01.PrintRune(rune(digit) + '0')
	}
}
