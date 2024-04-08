package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}

	factor := 1

	for a := 2; a <= nb; a++ {
		factor = factor * a
	}

	return factor
}
