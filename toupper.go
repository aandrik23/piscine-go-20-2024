package piscine

func ToUpper(s string) string {
	for _, v := range s {
		if v >= 97 && v <= 122 {
			v = v - 32
		}
	}
}
