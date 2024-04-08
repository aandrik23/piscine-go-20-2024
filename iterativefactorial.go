package piscine

func IterativeFactorial(nb int) int {
	factor := 1
	for a := nb; a == 1; a-- {
		factor = factor * a
	}
	return factor
}
