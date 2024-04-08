package piscine

func Sqrt(nb int) int {
	// Iterate from 0 to the square root of nb
	for i := 0; i*i <= nb; i++ {
		// If the square of i equals nb, return i
		if i*i == nb {
			return i
		}
	}
	// If the loop completes without finding a whole number square root, return 0
	return 0
}
