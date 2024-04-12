package piscine

func MakeRange(min, max int) []int {
	size := 0
	i := 0
	answer := make([]int, size)
	if min > max {
		return answer
	} else {
		for i = min - 1; i < max-1; i++ {
			answer[i] = i + 1
		}
	}
	return answer
}
