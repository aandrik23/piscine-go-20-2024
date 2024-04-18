package piscine

func JumpOver(str string) string {
	var result string
	if len(str) <= 2 {
		result = "\n"
		return result
	} else {
		for i := 2; i < len(str); i = i + 3 {
			result += string(str[i])
		}
		result += "\n"
	}
	return result
}
