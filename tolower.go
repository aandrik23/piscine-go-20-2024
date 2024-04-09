package piscine

func ToLower(s string) string {
	var lower string

	for _, v := range s {
		if v >= 65 && v <= 90 {
			lower += string(v + 32)
		} else {
			lower += string(v)
		}
	}
	return lower
}
