package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}

	factor := 1

	for a := nb; a == 0; a-- {
		factor = factor * a
	}

	return factor
}
