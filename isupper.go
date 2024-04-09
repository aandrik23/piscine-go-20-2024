package piscine

func IsUpper(s string) bool {
	var b bool

	for _, v := range s {
		if v < 'A' || v > 'Z' {
			b = false
		} else {
			b = true
		}
	}
	return b
}
