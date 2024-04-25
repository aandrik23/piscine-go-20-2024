package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	var word string
	for _, ch := range s {
		if ch == ' ' || ch == '\n' || ch == '\t' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(ch)
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
