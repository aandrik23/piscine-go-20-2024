package piscine

func DescendAppendRange(max, min int) []int {
	if min >= max {
		return []int{}
	}
	var s []int
	for i := max; i > min; i-- {
		s = append(s, i)
	}
	return s
}
