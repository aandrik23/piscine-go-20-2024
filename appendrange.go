package piscine

func AppendRange(min, max int) []int {

	i := 0
	var answer []int
	if min > max {
		answer = append(answer, 0)
	} else {
		for i = min; i < max; i++ {
			answer = append(answer, i+1)
		}
	}
	return answer
}
