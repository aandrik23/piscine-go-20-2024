package piscine

func ToUpper(s string) string {
	var upper string
	for _, v := range s {
		if v >= 97 && v <= 122 {
			upper += string(v - 32)
		} else {
			upper += string(v)
		}
	}
	return upper
}
