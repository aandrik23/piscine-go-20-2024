package piscine

func ConcatParams(args []string) string {
	var answer string
	for i := 0; i < len(args); i++ {
		if i < len(args)-1 {
			answer += args[i] + "\n"
		} else {
			answer += args[i]
		}
	}
	return answer
}
