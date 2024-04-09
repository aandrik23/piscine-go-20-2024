package piscine

func LastRune(s string) rune {
	aa := []rune(s)
	ab := len(s) - 1
	return aa[ab]
}
