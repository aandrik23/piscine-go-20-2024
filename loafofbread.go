package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var newStr string
	for _, char := range str {
		if char != ' ' {
			newStr += string(char)
		}
	}

	var result string
	for i := 0; i < len(newStr); i += 5 {
		end := i + 5
		if end > len(newStr) {
			end = len(newStr)
		}
		word := newStr[i:end]
		if i+5 < len(newStr) {
			word += " "
		}
		result += word
	}

	return result
}
