package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	for i := 0; i < len(str); {
		end := i + 5
		if end > len(str) {
			end = len(str)
		}
		if end < len(str) && str[end-1] != ' ' {
			for end > i && str[end-1] != ' ' {
				end--
			}
		}
		result += str[i:end]
		if end < len(str) && str[end-1] != ' ' {
			result += "\n"
		}
		i = end
	}

	return result
}
