package piscine

func SplitWhiteSpaces(s string) []string {
	var answer []string
	word := []byte(s)
	var wordBuild string
	for i := 0; i < len(s); i++ {
		if word[i] != ' ' && word[i] != '\n' && word[i] != '\t' {
			if i < len(s)-1 {
				wordBuild = wordBuild + (string(word[i]))
			} else {
				answer = append(answer, wordBuild)
			}
		} else {
			answer = append(answer, wordBuild)
			wordBuild = ""
		}
	}
	return answer
}
