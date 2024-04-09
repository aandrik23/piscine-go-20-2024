package piscine

func IsLower(s string) bool {
	var b bool

	for _, v := range s {
		if v < 'a' || v > 'z' {
			b = false
		} else {
			b = true
		}
	}
	return b
}
