package piscine

func NRune(s string, n int) rune {
	aa := []rune(s)
	var ac rune
	if n <= 0 || n > len(s) {
		return ac
	} else {
		return aa[n-1]
	}
}
