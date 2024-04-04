package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
    // Handle negative numbers
    if n < 0 {
        z01.PrintRune('-')
        printPositiveNumber(-n)
    } else {
        printPositiveNumber(n)
    }
}

func printPositiveNumber(n int) {
    // Handle the case when n is 0
    if n == 0 {
        z01.PrintRune('0')
        return
    }

    // Create a buffer to store digits as characters
    buffer := make([]rune, 0)

    // Extract digits of n and store them in the buffer
    for n > 0 {
        digit := n % 10
        buffer = append(buffer, rune('0'+digit))
        n /= 10
    }

    // Print digits in reverse order (since we extracted them in reverse order)
    for i := len(buffer) - 1; i >= 0; i-- {
        z01.PrintRune(buffer[i])
    }
}

}
