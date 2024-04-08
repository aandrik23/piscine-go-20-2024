package piscine

func IsPrime(nb int) bool {
	i := 0

	if i == 2 {
		return true
	}

	for i := 2; i*i <= nb; i++ {
		// If the square of i equals nb, return true
		if i*i == nb {
			return true
		}
	}
	// If the loop completes without finding a whole number square root, return false
	return false
}
