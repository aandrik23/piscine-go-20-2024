package piscine

func BasicAtoi(s string) int {
	r := 0
	for _, v := range s {
		if v < '0' || v > '9' {
			return 0
		} else {
			r = r*10 + int(v-'0')
		}
	}
	return r
}
