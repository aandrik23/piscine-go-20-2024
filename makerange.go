package piscine

func MakeRange(min, max int) []int {
	var answer []int
	if min >= max {
		return answer
	}
	size := max - min
	answer = make([]int, size)
	for i := 0; i < size; i++ {
		answer[i] = min + i
	}
	return answer
}
