package piscine

func SplitWhiteSpaces(s string) []string {
	var answer []string
	word := []byte(s)
	var wordBuild string
	for i := 0; i < len(s); i++ {
		wordBuild = wordBuild + (string(word[i]))
		if i == len(s)-1 {
			answer = append(answer, wordBuild)
		} else if word[i] == ' ' || word[i] == '\n' || word[i] == '\t' {
			answer = append(answer, wordBuild)
			wordBuild = ""
		}
	}
	return answer
}
