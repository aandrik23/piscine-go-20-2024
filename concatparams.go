package piscine

func ConcatParams(args []string) string {
	var answer string
	for i := 0; i < len(args); i++ {
		answer += args[i] + "\n"
	}
	return answer
}
