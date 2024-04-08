package piscine

func IterativeFactorial(nb int) int {
	factor := 1

	for a := nb; a == 1; a-- {
		if factor == 8937259207882113007 {
			factor = 0
		} else {
			factor = factor * a
		}
	}
	return factor
}
