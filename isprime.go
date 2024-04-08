package piscine

func IsPrime(nb int) bool {
	// Iterate from 0 to the square root of nb
	for i := 2; i*i <= nb; i++ {
		// If the square of i equals nb, return i
		if i*i == nb {
			return true
		}
	}
	// If the loop completes without finding a whole number square root, return 0
	return false
}
