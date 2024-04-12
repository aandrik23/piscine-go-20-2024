package piscine

func Capitalize(s string) string {
	first_letter := true
	result := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' && first_letter {
			result += string(v - 32)
			first_letter = false
		} else if v >= 'A' && v <= 'Z' && !first_letter {
			result += string(v + 32)
		} else if ((v >= 'a' && v <= 'z') || (v >= '0' && v <= '9')) && !first_letter {
			result += string(v)
		} else if ((v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9')) && !first_letter {
			result += string(v)
		} else if ((v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9')) && first_letter {
			result += string(v)
			first_letter = false
		} else {
			result += string(v)
			first_letter = true
		}
	}
	return result
}
