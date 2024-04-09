package piscine

func NRune(s string, n int) rune {
	aa := []rune(s)
	ab := len(s) - 1
	ac := '0'
	if n <= 0 || n > ab {
		return ac
	} else {
		return aa[ab]
	}
}
