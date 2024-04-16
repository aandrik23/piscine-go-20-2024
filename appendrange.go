package piscine

func AppendRange(min, max int) []int {
	i := 0
	var answer []int
	if min > max {
		return answer
	} else {
		for i = min - 1; i < max-1; i++ {
			answer = append(answer, i+1)
		}
	}
	return answer
}
