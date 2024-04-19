package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	for i := 0; i < len(str); i += 5 {
		end := i + 5
		if end > len(str) {
			end = len(str)
		}
		word := str[i:end]
		if end < len(str) && str[end-1] != ' ' {
			word += "\n"
		}
		result += word
	}

	return result
}
