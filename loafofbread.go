package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	for i := 0; i < len(str); i += 6 {
		end := i + 5
		if end >= len(str) {
			end = len(str)
		}
		result += str[i:end]
		if end < len(str)-1 && str[end] != ' ' {
			result += "\n"
		}
	}

	return result
}
