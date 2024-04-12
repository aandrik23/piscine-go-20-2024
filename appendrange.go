package piscine

func AppendRanintge(min, max int) []int {

	i := 0
	var answer []int
	if min > max {
		answer = answer
	} else {

		for i = min; i < max; i++ {
			answer = append(answer, i+1)
		}
	}
	return answer

}
