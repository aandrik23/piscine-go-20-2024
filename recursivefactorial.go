package piscine

func RecursiveFactorial(nb int) int {
	// Error case: negative numbers or overflows
	if nb < 0 || nb > 20 {
		return 0
	}

	// Base case: factorial of 0 is 1
	if nb == 0 {
		return 1
	}

	// Recursive case: nb! = nb * (nb-1)!
	return nb * RecursiveFactorial(nb-1)
}
