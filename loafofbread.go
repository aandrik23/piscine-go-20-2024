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
		result += word
		// Insert a space if there are more characters in the new string
		if end < len(newStr) && end%5 != 0 {
			result += " "
		}
	}
	// Remove the last space, if any
	if result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result
}
